package hamsoyamodels

import (
	"MF/db"
	"MF/helperfunc"
	"github.com/sirupsen/logrus"
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
	Error             error           `json:"error"`
	Page              int64           `json:"page"`
	TotalPage         int64           `json:"totalPage"`
	URL               string          `json:"url"`
	HamsoyaClientList []HamsoyaClient `json:"data"`
}

func GetHamsoyaClientById(id int64) (HamsoyaClient HamsoyaClient, err error) {
	postgresDb := db.GetHamsoyaPostgresDb()
	postgresDb.Where("id = ?", id).First(&HamsoyaClient)
	if err := postgresDb.Where("id = ?", id).First(&HamsoyaClient).Error; err != nil {
		return HamsoyaClient, err
	}
	return HamsoyaClient, nil
}

func GetHamsoyaClients(client HamsoyaClient, clientsSlice *ResponseHamsoyaClients, time helperfunc.TimeInterval, page int64) (clientsSliceOver *ResponseHamsoyaClients) {

	if err := db.GetHamsoyaPostgresDb().Where(&client).Where(`create_date > ? and create_date < ?`, time.From, time.To).Limit(100).Offset(page * 100).Order("create_date desc").Find(&clientsSlice.HamsoyaClientList).Error; err != nil {
		clientsSlice.Error = err
		logrus.Println(" ", err)
	}
	return
}

func GetHamsoyaClientsCount(client HamsoyaClient, time helperfunc.TimeInterval) (clientsSlice ResponseHamsoyaClients) {

	if err := db.GetHamsoyaPostgresDb().Table("clients").Where(&client).Where(`create_date > ? and create_date < ?`, time.From, time.To).Count(&clientsSlice.TotalPage).Error; err != nil {
		clientsSlice.Error = err
		logrus.Println(" ", err)
	}
	return
}
