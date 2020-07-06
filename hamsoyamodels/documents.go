package hamsoyamodels

import (
	"MF/db"
	"MF/helperfunc"
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
	Error          error        `json:"error"`
	Page           int64        `json:"page"`
	TotalPage      int64        `json:"totalPage"`
	URL            string       `json:"url"`
	Documents []HamsoyaDocument `json:"Documents"`
}

func GetHamsoyaDocuments(Document HamsoyaDocument, documentSlice *ResponseHamsoyaDocuments, time helperfunc.TimeInterval, page int64) (DocumentsSliceOver *ResponseHamsoyaDocuments) {
	postgresDb := db.GetHamsoyaPostgresDb()

	if err := postgresDb.Where(&Document).Where(`create_date > ? and create_date < ?`, time.From, time.To).Limit(100).Offset(page * 100).Order("create_date desc").Find(&documentSlice.Documents).Error; err != nil {
		documentSlice.Error = err
	}
	return
}

func GetHamsoyaDocumentsCount(Document HamsoyaDocument, time helperfunc.TimeInterval) (DocumentsSliceOver ResponseHamsoyaDocuments) {
	postgresDb := db.GetHamsoyaPostgresDb()

	if err := postgresDb.Table("documents").Where(&Document).Where(`create_date > ? and create_date < ?`, time.From, time.To).Count(&DocumentsSliceOver.TotalPage).Error; err != nil {
		DocumentsSliceOver.Error = err
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
