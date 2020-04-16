package models

import (
	"MF/db"
	"encoding/xml"
	"time"
)

// PostCheckReqRawXML ...
type PostCheckReqRawXML struct {
	XMLName      xml.Name  `xml:"request"`
	ID           int       `xml:"-" gorm:"column:id"`
	Command      string    `xml:"command" gorm:"column:command"`
	QueueID      int       `xml:"queueID" gorm:"column:queue_id"`
	PreSharedKey string    `xml:"preSharedKey" gorm:"column:pre_shared_key"`
	HashSum      string    `xml:"hashSum" gorm:"column:hash_sum"`
	CreatedAt    time.Time `xml:"-" gorm:"column:create_time;type:timestamp"`
	IPAddress    string    `xml:"-" gorm:"column:ip_address"`
}
//
//SaveModel saves PostCheckReqRawXML model in db
func (postCheckReq *PostCheckReqRawXML) SaveModel() {
	db := db.GetPostgresDb()
	db.Create(postCheckReq)
}

//TableName for changing struct name to db name
func (postCheckReq PostCheckReqRawXML) TableName() string {
	return "tb_request_log"
}

// PostCheckResXML ...
type PostCheckResXML struct {
	XMLName xml.Name `xml:"response"`
	Command string   `xml:"command"`
	QueueID int      `xml:"queueID"`
	Result  int      `xml:"result"`
	Reason  string   `xml:"reason"`
	HashSum string   `xml:"hashSum"`
}
