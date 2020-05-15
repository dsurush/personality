package hamsoyamodels

import "time"

type HamsoyaViewTrans struct {
	ID             int64     `xml:"id"`
	DateTime       string    `xml:"datetime"`
	RequestType    string    `xml:"request_type"` ///?
	VendorID       int64     `xml:"vendor_id"`
	PhoneNum       string    `xml:"phone_num"`
	ClientReceiver string    `xml:"client_receiver"`
	Amount         float64   `xml:"amount"`
	ExternalFee    float64   `xml:"external_fee"`
	TotalAmount    float64   `xml:"total_amount"`
	Code           string    `xml:"code"`
	ExtTransId     string    `xml:"ext_trans_id"`
	Description    string    `xml:"description"`
	PrDescription  string    `xml:"pr_description"`
	QrCode         string    `xml:"qr"`
	CreateDate     time.Time `xml:"create_date"`
}

func (*HamsoyaViewTrans) TableName() string {
	return "view_trans"
}

type HamsoyaViewTransaction struct {
	TransType         string    `xml:"trans_type"`
	TransId           int64     `xml:"trans_id"`
	TransDate         time.Time `xml:"trans_date"`
	TransStatus       string    `xml:"trans_status"`
	TransExtStatus    string    `xml:"trans_ext_status"`
	AgentId           int64     `xml:"agent_id"`
	VendorId          int64     `xml:"vendor_id"`
	ClientPayerId     int64     `xml:"client_paye_id"`
	ClientPayer       string    `xml:"clien_payer"`
	AccountPayerId    int64     `xml:"account_paye_id"`
	AccountPayer      string    `xml:"account_payer"`
	ClientReceiverId  int64     `xml:"client_receiver_id"`
	ClientReceiver    string    `xml:"client_receiver"`
	AccountReceiverId int64     `xml:"account_receiver_id"`
	AccountReceiver   string    `xml:"account_receiver"`
	Amount            float64   `xml:"amount"`
	ExternalFee       float64   `xml:"external_fee"`
	InternalFee       float64   `xml:"internal_fee"`
	TotalAmount       float64   `xml:"total_amount"`
	Description       string    `xml:"description"`
	IsTransAvailable  bool      `xml:"is_trans_available"`
	AccExtFeeId       int64     `xml:"acc_ext_fee_id"`
}

func (*HamsoyaViewTransaction) TableName() string{
	return "view_transaction"
}