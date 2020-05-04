package models

import (
	"MF/db"
	"log"
	"time"
)
// Create New struct for report
type ViewReport struct {
	ID                int64     `xml:"id"`
	RequestId         int64     `xml:"request_id"`
	PaymentID         int64     `xml:"payment_id"`
	//PreCheckQueueID   int64     `xml:"pre_check_queue_id"`
	VendorID            int       `xml:"vendor_id"`
	VendorName        string    `xml:"vendor_name"`
	Route             int       `xml:"route"`
	RequestType       string    `xml:"request_type"` ///?
	SendToCft         int       `xml:"send_to_cft"`
	AccountPayer      string    `xml:"account_payer"`
	AccountReceiver   string    `xml:"account_receiver"`
	StateID           string    `xml:"state_id"`
	Aggregator        string    `xml:"aggregator"`
	CreateTime        time.Time `xml:"create_time"`
	Amount            float64     `xml:"amount"`
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

func GetViewReport(size, page int64) (Report []ViewReport, err error) {
	postgresDb := db.GetPostgresDb()
	err = postgresDb.Table(`view_report`).Select("view_report.*").Limit(size).Offset(page * size).Scan(&Report).Error
	if err != nil {
		log.Printf("can't get from db view report %e\n", err)
		return nil, err
	}
	return Report, nil
}