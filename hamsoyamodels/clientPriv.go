package hamsoyamodels

import "time"

type HamsoyaClientPriv struct {
	Id          int64     `xml:"id"`
	Name        string    `xml:"name"`
	Address     string    `xml:"address"`
	INN         string    `xml:"inn"`
	Sex         string    `xml:"sex"`
	DocNo       string    `xml:"doc_no"`
	IssuingAuth string    `xml:"issuing_auth"`
	BirthDate   time.Time `xml:"birth_date"`
	IssueDate   time.Time `xml:"issue_date"`
	ExpiryDate  time.Time `xml:"expiry_date"`
}

type ListHamsoyaClientPrivResponse struct {
	Error error
	Size int64
	ClientPrivs []HamsoyaClientPriv
}

func (*HamsoyaClientPriv) TableName() string {
	return "client_priv"
}
