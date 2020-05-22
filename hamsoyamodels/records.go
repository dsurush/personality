package hamsoyamodels

import (
	"MF/db"
	"time"
)

type HamsoyaRecord struct {
	Id         int64     `xml:"id"`
	AccountId  int64     `xml:"account_id"`
	DocumentId int64     `xml:"document_id"`
	Amount     float64   `xml:"amount"`
	StartSaldo float64   `xml:"start_saldo"`
	Type       string    `xml:"type"`
	CreateDate time.Time `xml:"create_date"`
}

func (*HamsoyaRecord) TableName() string {
	return "records"
}

type ResponseHamsoyaRecords struct {
	Error error
	Count int64
	HamsoyaRecordList []HamsoyaRecord
}

func GetHamsoyaRecordById(id int64) (HamsoyaRecord HamsoyaRecord, err error){
	postgresDb := db.GetHamsoyaPostgresDb()
	postgresDb.Where("id = ?", id).First(&HamsoyaRecord)
	if err := postgresDb.Where("id = ?", id).First(&HamsoyaRecord).Error; err != nil {
		return HamsoyaRecord, err
	}
	return HamsoyaRecord, nil
}

func GetHamsoyaRecords(hamsoyaRecord HamsoyaRecord, rows, pages int64) (HamsoyaRecord ResponseHamsoyaRecords){
	pages--
	if pages < 0 {
		pages = 0
	}
	postgresDb := db.GetHamsoyaPostgresDb()
	if err := postgresDb.Where(&hamsoyaRecord).Limit(rows).Offset(rows * pages).Find(&HamsoyaRecord.HamsoyaRecordList).Count(&HamsoyaRecord.Count).Error; err != nil{
		HamsoyaRecord.Error = err
	}
	return
}