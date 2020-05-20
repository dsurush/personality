package hamsoyamodels

import (
	"MF/db"
	"time"
)

type HamsoyaDocument struct {
	Id          int64     `xml:"id"`
	AccountDt   int64     `xml:"account_dt"`
	AccountCt   int64     `xml:"account_ct"`
	Amount      float64   `xml:"amount"`
	Description string    `xml:"description"`
	TransId     int64     `xml:"trans_id"`
	StatusId    int64     `xml:"status_id"`
	CreateDate  time.Time `xml:"create_date"`
	CancelDocId int64     `xml:"cancel_doc_id"`
}

func (*HamsoyaDocument) TableName() string {
	return "documents"
}

type ResponseHamsoyaDocuments struct {
	Error     error
	Count     int64
	Documents []HamsoyaDocument
}

func GetHamsoyaDocuments(Document HamsoyaDocument, rows, pages int64) (Documents ResponseHamsoyaDocuments) {
	postgresDb := db.GetHamsoyaPostgresDb()
	pages--
	if pages < 0 {
		pages = 0
	}
	if err := postgresDb.Where(&Document).Limit(rows).Offset(rows * pages).Find(&Documents.Documents).Count(&Documents.Count).Error; err != nil {
		Documents.Error = err
	}
	return
}

func GetHamsoyaDocument(id int64) (Document HamsoyaDocument, err error) {
	postgresDb := db.GetHamsoyaPostgresDb()

	if err := postgresDb.Where("id = ?", id).First(&Document).Error; err != nil {
		return Document, err
	}
	return Document, nil
}