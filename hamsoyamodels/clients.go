package hamsoyamodels

import (
	"MF/db"
	"time"
)

type HamsoyaClient struct {
	Id         int64     `xml:"id"`
	PhoneNum   string    `xml:"phone_num"`
	AgentId    int64     `xml:"agent_id"`
	TypeId     int64     `xml:"type_id"`
	IsActive   bool      `xml:"is_active"`
	CreateDate time.Time `xml:"create_date"`
	CloseDate  time.Time `xml:"close_date"`
	Identify   bool      `xml:"identify"`
	Name       string    `xml:"name"`
}

func (*HamsoyaClient) TableName() string {
	return "clients"
}

type ResponseHamsoyaClients struct {
	Error             error
	Count             int64
	HamsoyaClientList []HamsoyaClient
}

func GetHamsoyaClientById(id int64) (HamsoyaClient HamsoyaClient, err error) {
	postgresDb := db.GetHamsoyaPostgresDb()
	postgresDb.Where("id = ?", id).First(&HamsoyaClient)
	if err := postgresDb.Where("id = ?", id).First(&HamsoyaClient).Error; err != nil {
		return HamsoyaClient, err
	}
	return HamsoyaClient, nil
}

func GetHamsoyaClients(hamsoyaClient HamsoyaClient, rows, pages int64) (HamsoyaClient ResponseHamsoyaClients) {
	pages--
	if pages < 0 {
		pages = 0
	}
	postgresDb := db.GetHamsoyaPostgresDb()
	if err := postgresDb.Where(&hamsoyaClient).Limit(rows).Offset(rows * pages).Find(&HamsoyaClient.HamsoyaClientList).Count(&HamsoyaClient.Count).Error; err != nil {
		HamsoyaClient.Error = err
	}
	return
}
