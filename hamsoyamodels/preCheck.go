package hamsoyamodels

import (
	"MF/db"
	"time"
)

type HamsoyaPreCheck struct {
	Id              int64     `xml:"id"`
	AgentId         int64     `xml:"agent_id"`
	VendorId        int64     `xml:"vendor_id"`
	ClientPayerId   int64     `xml:"client_paye_id"`
	ClientPayer     string    `xml:"clien_payer"`
	AccountPayer    string    `xml:"account_payer"`
	ClientReceiver  string    `xml:"client_receiver"`
	AccountReceiver string    `xml:"account_receiver"`
	Qr              string    `xml:"qr"`
	RequestType     string    `xml:"request_type"`
	Amount          float64   `xml:"amount"`
	ExternalFee     float64   `xml:"external_fee"`
	InternalFee     float64   `xml:"internal_fee"`
	Description     string    `xml:"description"`
	ExtPrecheckId   string    `xml:"ext_precheck_id"`
	HashForPayment  string    `xml:"hash_for_payment"`
	CreateDate      time.Time `xml:"create_date"`
	TopupPreCheckId int64     `xml:"topup_pre_check_id"`
}

func (*HamsoyaPreCheck) TableName() string {
	return "pre_check"
}

type ResponseHamsoyaPreChecks struct {
	Error               error
	Count               int64
	HamsoyaPreCheckList []HamsoyaPreCheck `json:"hamsoya_pre_check_list"`
}

func GetHamsoyaPreCheckById(id int64) (HamsoyaPreCheck HamsoyaPreCheck, err error) {
	postgresDb := db.GetHamsoyaPostgresDb()
	postgresDb.Where("id = ?", id).First(&HamsoyaPreCheck)
	if err := postgresDb.Where("id = ?", id).First(&HamsoyaPreCheck).Error; err != nil {
		return HamsoyaPreCheck, err
	}
	return HamsoyaPreCheck, nil
}

func GetHamsoyaPreChecks(hamsoyaPreCheck HamsoyaPreCheck, rows, pages int64) (HamsoyaPreCheck ResponseHamsoyaPreChecks) {
	pages--
	if pages < 0 {
		pages = 0
	}
	postgresDb := db.GetHamsoyaPostgresDb()
	if err := postgresDb.Where(&hamsoyaPreCheck).Limit(rows).Offset(rows * pages).Find(&HamsoyaPreCheck.HamsoyaPreCheckList).Count(&HamsoyaPreCheck.Count).Error; err != nil {
		HamsoyaPreCheck.Error = err
	}
	return
}
