package hamsoyamodels

import (
	"MF/db"
	"MF/helperfunc"
	"github.com/sirupsen/logrus"
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
	Error             error           `json:"error"`
	Page              int64           `json:"page"`
	TotalPage         int64           `json:"totalPage"`
	URL               string          `json:"url"`
	HamsoyaRecordList []HamsoyaRecord `json:"data"`
}

func GetHamsoyaRecordById(id int64) (HamsoyaRecord HamsoyaRecord, err error) {
	postgresDb := db.GetHamsoyaPostgresDb()
	postgresDb.Where("id = ?", id).First(&HamsoyaRecord)
	if err := postgresDb.Where("id = ?", id).First(&HamsoyaRecord).Error; err != nil {
		return HamsoyaRecord, err
	}
	return HamsoyaRecord, nil
}


func GetHamsoyaRecords(record HamsoyaRecord, recordsSlice *ResponseHamsoyaRecords, time helperfunc.TimeInterval, page int64) (recordsSliceOver *ResponseHamsoyaRecords) {

	if err := db.GetHamsoyaPostgresDb().Where(&record).Where(`create_date > ? and create_date < ?`, time.From, time.To).Limit(100).Offset(page * 100).Order("create_date desc").Find(&recordsSlice.HamsoyaRecordList).Error; err != nil {
		recordsSlice.Error = err
		logrus.Println(" ", err)
	}
	return
}

func GetHamsoyaRecordsCount(record HamsoyaRecord, time helperfunc.TimeInterval) (recordsSlice ResponseHamsoyaRecords) {

	if err := db.GetHamsoyaPostgresDb().Table("clients").Where(&record).Where(`create_date > ? and create_date < ?`, time.From, time.To).Count(&recordsSlice.TotalPage).Error; err != nil {
		recordsSlice.Error = err
		logrus.Println(" ", err)
	}
	return
}
