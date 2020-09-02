package models

import (
	"MF/db"
	"MF/helperfunc"
	"fmt"
	"log"
	"time"
)

// Create New struct for report
type ViewReport struct {
	ID        int64 `xml:"id"`
	RequestId int64 `xml:"request_id"`
	PaymentID int64 `xml:"payment_id"`
	//PreCheckQueueID   int64     `xml:"pre_check_queue_id"`
	VendorID          int       `xml:"vendor_id"`
	VendorName        string    `xml:"vendor_name"`
	Route             int       `xml:"route"`
	RequestType       string    `xml:"request_type"` ///?
	SendToCft         int       `xml:"send_to_cft"`
	AccountPayer      string    `xml:"account_payer"`
	AccountReceiver   string    `xml:"account_receiver"`
	StateID           string    `xml:"state_id"`
	Aggregator        string    `xml:"aggregator"`
	CreateTime        time.Time `xml:"create_time"`
	Amount            float64   `xml:"amount"`
	AmountWithCommiss float64   `xml:"amount_with_commiss"`
	Commiss           float64   `xml:"commiss"`
	CardHash          string    `xml:"card_hash"`
	AgrTransTime      time.Time `xml:"agr_trans_time"`
	AggregatorStatus  string    `xml:"aggregator_status"`
	GateWay           string    `xml:"gate_way"`
	//	QrCode            string    `xml:"qr_code"`
	//	NameRus           string    `xml:"name_rus"`
	//	TimeDiff          time.Time `xml:"time_diff"`
}
type ResponseViewReports struct {
	Error          error        `json:"error"`
	Page           int64        `json:"page"`
	TotalPage      int64        `json:"totalPage"`
	URL            string       `json:"url"`
	ViewReportList []ViewReport `json:"data"`
}

func (viewReport ViewReport) TableName() string {
	return "view_report"
}

func GetViewReport(report ViewReport, viewReportSlice *ResponseViewReports, time helperfunc.TimeInterval, page int64) (reportSliceOver *ResponseViewReports) {
	postgresDb := db.GetPostgresDb()
	if err := postgresDb.Where(&report).Where(`create_time > ? and create_time < ?`, time.From, time.To).Limit(100).Offset(page * 100).Order("create_time desc").Find(&viewReportSlice.ViewReportList).Error; err != nil {
		log.Printf("can't get from db view report %e\n", err)
		viewReportSlice.Error = err
	}
	return
}

func GetViewReportForExcel(report ViewReport, viewReportSlice *ResponseViewReports, time helperfunc.TimeInterval) (reportSliceOver *ResponseViewReports) {
	postgresDb := db.GetPostgresDb()
	if err := postgresDb.Where(&report).Where(`create_time > ? and create_time < ?`, time.From, time.To).Limit(200).Order("create_time desc").Find(&viewReportSlice.ViewReportList).Error; err != nil {
		log.Printf("can't get from db view report %e\n", err)
		viewReportSlice.Error = err
	}
	return
}

func GetViewReportCount(report ViewReport, time helperfunc.TimeInterval) (ReportSlice ResponseViewReports) {
	postgresDb := db.GetPostgresDb()
	if err := postgresDb.Table("view_report").Where(&report).Where(`create_time > ? and create_time < ?`, time.From, time.To).Count(&ReportSlice.TotalPage).Error; err != nil {
		log.Printf("can't get from db view report %e\n", err)
		ReportSlice.Error = err
	}
	return
}

func GetViewReportById(id int64) (ViewReport ViewReport) {
	postgresDb := db.GetPostgresDb()
	fmt.Println("i am id = ", id)
	err := postgresDb.Table(`view_report`).
		Where(`id = ?`, id).First(&ViewReport).Error
	if err != nil {
		//		transactionSlice.Error = err
		fmt.Println("can't take from db")
	}
	return
}
