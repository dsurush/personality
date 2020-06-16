package models

import (
	"MF/db"
	"encoding/xml"
	"fmt"
	"log"
	"time"
)

// PreCheckReqRawXML Структура (для проверки доступности вендора)
type PreCheckReqRawXML struct {
	XMLName           xml.Name  `xml:"request"`
	ID                int       `xml:"-" gorm:"column:id"`
	Command           string    `xml:"command" gorm:"column:command"`
	Type              string    `xml:"requestType" gorm:"column:request_type"`
	Vendor            int       `xml:"vendor,omitempty" gorm:"column:vendor"`
	Qr                string    `xml:"qr,omitempty" gorm:"column:qr_code"`
	Amount            float64   `xml:"amount" gorm:"column:amount"`
	AmountWithCommiss float64   `xml:"amountWithCommiss" gorm:"amount_with_commiss"`
	AccountPayer      string    `xml:"accountPayer" gorm:"column:account_payer"`
	CardHash          string    `xml:"cardHash" gorm:"card_hash"`
	AccountReceiver   string    `xml:"accountReceiver,omitempty" gorm:"column:account_receiver"`
	PreSharedKey      string    `xml:"preSharedKey" gorm:"column:pre_shared_key"`
	HashSum           string    `xml:"hashSum" gorm:"column:hash_sum"`
	CreatedAt         time.Time `xml:"-" gorm:"column:create_time;type:timestamp"`
	IPAddress         string    `xml:"-" gorm:"column:ip_address"`
	GateWay           string    `xml:"gateWay" gorm:"column:gate_way; default:'MEGAFON'"`
}

//SaveModel saves PreCheckReqRawXML model in db
func (preCheckReq *PreCheckReqRawXML) SaveModel() {
	db := db.GetPostgresDb()
	if err := db.Create(preCheckReq).Error; err != nil {
		log.Println("PreCheckReqRawXMLSaveModel: ", err)
	}
}

//FindFirst finds first request by id
func (preCheckReq *PreCheckReqRawXML) FindFirst(id uint) {
	db := db.GetPostgresDb()
	db.First(preCheckReq, id)
}

//TableName for changing struct name to db name
func (preCheckReq PreCheckReqRawXML) TableName() string {
	return "tb_request_log"
}

// PreCheckResXML ...
type PreCheckResXML struct {
	XMLName         xml.Name `xml:"response"`
	Command         string   `xml:"command"`
	Vendor          int      `xml:"vendor,omitempty"`
	Type            string   `xml:"requestType,omitempty"`
	Amount          float64  `xml:"amount"`
	AccountPayer    string   `xml:"accountPayer"`
	AccountReceiver string   `xml:"accountReceiver,omitempty"`
	Qr              string   `xml:"qr,omitempty" gorm:"column:qr_code"`
	MerchantID      int64    `xml:"huMerchantID,omitempty"`
	MerchantName    string   `xml:"huMerchantName,omitempty"`
	Reason          string   `xml:"reason"`
	Result          int      `xml:"result"`
	PrecheckQueueID int      `xml:"precheckQueueID"`
	Comment         string   `xml:"info,omitempty"`
	HashSum         string   `xml:"hashSum"`
}

//TablePreeCheck ...
type TablePreeCheck struct {
	ID            int               `gorm:"column:id"`
	Request       PreCheckReqRawXML `gorm:"foreignkey:RequestID"`
	RequestID     uint              `gorm:"column:request_id"`
	Status        bool              `gorm:"column:status"`
	CreatedAt     time.Time         `gorm:"column:create_time"`
	Sector        int               `gorm:"column:sector"`  // need for barki tojik
	ClientName    string            `gorm:"column:name"`    // need for barki tojik
	ClientAddress string            `gorm:"column:address"` // need for barki tojik
}

//
//SaveModel saves TablePreeCheck model in db
func (tablePreeCheck *TablePreeCheck) SaveModel() {
	db := db.GetPostgresDb()
	db.Create(tablePreeCheck)
}

// UpdateModel updates transaction model
func (tablePreeCheck *TablePreeCheck) UpdateModel() {
	db := db.GetPostgresDb()
	db.Save(tablePreeCheck)
}

//FindFirst finds precheck by id and cretae_time last ten minutes
func (tablePreeCheck *TablePreeCheck) FindFirst(preCheckQueueID int) {
	db := db.GetPostgresDb()
	//now := fmt.Sprint(time.Now().Format("2006-01-02 15:04:05"))
	now := fmt.Sprint(time.Now().Add(time.Second).Format("2006-01-02 15:04:05"))
	nowMinusTenMin := fmt.Sprint(time.Now().Add(time.Minute * -20).Format("2006-01-02 15:04:05"))
	db.Where("id = ? AND create_time BETWEEN ? AND ?", preCheckQueueID, nowMinusTenMin, now).Last(tablePreeCheck)
}

//TableName for changing struct name to db name
func (tablePreeCheck TablePreeCheck) TableName() string {
	return "tb_precheck"
}
