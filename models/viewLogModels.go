package models

import (
	"MF/db"
	"log"
	"time"
)

type ViewLog struct {
	Id              int64
	Command         string  `xml:"command"`
	Response        string  `xml:"response"`
	PrecheckQueueID int64   `xml:"precheck_queue_id"`
	QueueID         int64   `xml:"queue_id"`
	Type            string  `xml:"request_type,omitempty"`
	VendorID        int64   `xml:"vendor_id,omitempty"`
	AccountPayer    string  `xml:"account_payer"`
	Amount          float64 `xml:"amount"`
	AmountWithCommiss  float64 `xml:"amount_with_commiss"`
	CreatedAt         time.Time `xml:"create_time"`
	GateWay           string    `xml:"gate_way"`
}
func GetViewLog(size, page int64) (Report []ViewLog, err error) {
	postgresDb := db.GetPostgresDb()
	err = postgresDb.Table(`view_log`).Select("view_log.*").Limit(size).Offset(page * size).Scan(&Report).Error
	if err != nil {
		log.Printf("can't get from db view log %e\n", err)
		return nil, err
	}
	return Report, nil
}