package models

import (
	db2 "MF/db"
	"encoding/xml"
	"time"
)

// PingHumoXML struct for test
type PingHumoXML struct {
	XMLName xml.Name `xml:"response"`
	CodeTag string   `xml:"status"`
}

// ResponseLog ...
type ResponseLog struct {
	ID           int       `gorm:"column:id"`
	RequestID    int       `gorm:"column:request_id"`
	Response     []byte    `gorm:"column:response"`
	ResponseTime time.Time `gorm:"column:create_time"`
}
//
////SaveModel saves ResponseLog model in db
func (responseLog *ResponseLog) SaveModel() {
	db := db2.GetPostgresDb()
	db.Create(responseLog)
}

//TableName for changing struct name to db name
func (responseLog ResponseLog) TableName() string {
	return "tb_response_log"
}

// ResponseXML структура для возврата ошибок
type ResponseXML struct {
	XMLName xml.Name `xml:"response"`
	Code    int      `xml:"code"`
	Message string   `xml:"message"`
}

//SaveModel saves ResponseXML model in db
func (errorResponse *ResponseXML) SaveModel() {
	db := db2.GetPostgresDb()
	db.Create(errorResponse)
}

//TableName for changing struct name to db name
func (errorResponse ResponseXML) TableName() string {
	return "tb_response_log"
}

// RawXML ...
type RawXML struct {
	ID           int    `xml:"-" gorm:"column:id"`
	Command      string `xml:"command" gorm:"column:command"`
	Type         string `xml:"requestType" gorm:"-"`
	PreSharedKey string `xml:"preSharedKey"`
	HashSum      string `xml:"hashSum"`
	GateWay      string `xml:"gateWay" gorm:"column:gate_way; default:'MEGAFON'"`
	IPAddress    string `xml:"-" gorm:"column:ip_address"`
}

//SaveModel saves RawXML model in db
func (rawXML *RawXML) SaveModel() {
	db := db2.GetPostgresDb()
	db.Create(rawXML)
}

//TableName for changing struct name to db name
func (rawXML RawXML) TableName() string {
	return "tb_request_log"
}

type RefundedCardTransactions struct {
	ID             uint64    `gorm:"column:id"`
	CardHash       string    `gorm:"column:card_hash"`
	RefundTime     time.Time `gorm:"column:refund_time"`
	RefundAmount   float64   `form:"column:refund_amount"`
	Description    string    `xml:"column:description"`
	RefundForTrans int       `xml:"column:refund_for_trans"`
}

func (refTrans *RefundedCardTransactions) SaveModel() {
	db := db2.GetPostgresDb()
	db.Create(refTrans)
}

func (refTrans RefundedCardTransactions) TableName() string {
	return "tb_refunded_transactions"
}
