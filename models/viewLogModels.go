package models

import (
	"MF/db"
	"log"
	"time"
)

type ViewLog struct {
	Id                int64
	Command           string    `xml:"command"`
	Response          string    `xml:"response"`
	PrecheckQueueID   int64     `xml:"precheck_queue_id"`
	QueueID           int64     `xml:"queue_id"`
	Type              string    `xml:"request_type,omitempty"`
	VendorID          int64     `xml:"vendor_id,omitempty"`
	AccountPayer      string    `xml:"account_payer"`
	Amount            float64   `xml:"amount"`
	AmountWithCommiss float64   `xml:"amount_with_commiss"`
	CreatedAt         time.Time `xml:"create_time"`
	GateWay           string    `xml:"gate_way"`
}

type ViewLogDTO struct {
	Id           int64     `xml:"id"`
	Command      string    `xml:"command"`
	Type         string    `xml:"request_type,omitempty"`
	VendorID     int64     `xml:"vendor_id,omitempty"`
	AccountPayer string    `xml:"account_payer"`
	CreatedAt    time.Time `xml:"create_time"`
}

func GetViewLogs(size, page int64) (viewLogs []ViewLog, err error) {
	postgresDb := db.GetPostgresDb()
	page--
	if page < 0 {
		page = 0
	}
	err = postgresDb.Table(`view_log`).Select("view_log.*").Limit(size).Offset(page * size).Scan(&viewLogs).Error
	if err != nil {
		log.Printf("can't get from db view log %e\n", err)
		return nil, err
	}
	return viewLogs, nil
}

func GetViewLogById(id int64) (viewLog ViewLog, err error) {
	postgresDb := db.GetPostgresDb()
	err = postgresDb.Table(`view_log`).Select("view_log.*").Where("id = ?", id).Scan(&viewLog).Error
	if err != nil {
		log.Printf("can't get from db one view log %e\n", err)
		return viewLog, err
	}
	return viewLog, err
}

func GetViewLogsDTO(size, page int64) (viewLogs []ViewLogDTO, err error) {
	postgresDb := db.GetPostgresDb()
	page--
	if page < 0 {
		page = 0
	}
	err = postgresDb.Table(`view_log`).Select("view_log. id, command, request_type, vendor_id, account_payer, create_time ").Limit(size).Offset(page * size).Scan(&viewLogs).Error
	if err != nil {
		log.Printf("can't get from db view log %e\n", err)
		return nil, err
	}
	return viewLogs, nil
}
