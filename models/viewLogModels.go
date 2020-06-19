package models

import (
	"MF/db"
	"MF/helperfunc"
	"github.com/sirupsen/logrus"
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

type ResponseViewLogs struct {
	Error       error
	Count       int64
	ViewLogList []ViewLog
}

func GetViewLogs(size, page int64) (viewLogs ResponseViewLogs) {
	postgresDb := db.GetPostgresDb()
	page--
	if page < 0 {
		page = 0
	}
	err := postgresDb.Table(`view_log`).Select("view_log.*").Limit(size).Offset(page * size).Scan(&viewLogs.ViewLogList).Count(&viewLogs.Count).Error
	if err != nil {
		log.Printf("can't get from db view log %e\n", err)
		viewLogs.Error = err
		return
	}
	return
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

//func GetViewLogsDTO(size, page int64) (viewLogs []ViewLogDTO, err error) {
//	postgresDb := db.GetPostgresDb()
//	page--
//	if page < 0 {
//		page = 0
//	}
//	err = postgresDb.Table(`view_log`).Select("view_log. id, command, request_type, vendor_id, account_payer, create_time ").Limit(size).Offset(page * size).Scan(&viewLogs).Error
//	if err != nil {
//		log.Printf("can't get from db view log %e\n", err)
//		return nil, err
//	}
//	return viewLogs, nil
//}
type ResponseViewLogsList struct {
	Error          error        `json:"error"`
	Page           int64        `json:"page"`
	TotalPage      int64        `json:"totalPage"`
	URL            string       `json:"url"`
	ViewLogList []ViewLogDTO `json:"logs_list"`
}

func GetViewLogsDTO(ViewLogDTO ViewLogDTO, ViewLogDTOsSlice *ResponseViewLogsList, time helperfunc.TimeInterval, page int64) (ViewLogsDTOSliceOver *ResponseViewLogsList) {

	if err := db.GetPostgresDb().Where(&ViewLogDTO).Where(`create_time > ? and create_time < ?`, time.From, time.To).Limit(100).Offset(page * 100).Order("create_time desc").Find(&ViewLogDTOsSlice.ViewLogList).Error; err != nil {
		ViewLogDTOsSlice.Error = err
		logrus.Println(" ", err)
	}
	return
}

func GetViewLogsDTOCount(ViewLogDTO ViewLogDTO, time helperfunc.TimeInterval) (ViewLogsDTOSliceOver ResponseViewLogsList) {
	if err := db.GetPostgresDb().Table("view_log").Where(&ViewLogDTO).Where(`create_time > ? and create_time < ?`, time.From, time.To).Count(&ViewLogsDTOSliceOver.TotalPage).Error; err != nil {
		ViewLogsDTOSliceOver.Error = err
		logrus.Println(" ", err)
	}
	return
}

func (*ViewLogDTO) TableName() string {
	return "view_log"
}