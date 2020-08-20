package models

import (
	"MF/db"
	"MF/helperfunc"
	"encoding/xml"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

// PaymentReqRawXML структура ...
type PaymentReqRawXML struct {
	XMLName           xml.Name  `xml:"request"`
	ID                int       `xml:"-" gorm:"column:id"`
	Command           string    `xml:"command" gorm:"column:command"`
	Vendor            int       `xml:"vendor,omitempty" gorm:"column:vendor"`
	Amount            float64   `xml:"amount" gorm:"column:amount"`
	AmountWithCommiss float64   `xml:"amountWithCommiss" gorm:"column:amount_with_commiss"`
	AccountPayer      string    `xml:"accountPayer" gorm:"column:account_payer"`
	CardHash          string    `xml:"cardHash" gorm:"card_hash"`
	AccountReceiver   string    `xml:"accountReceiver,omitempty" gorm:"column:account_receiver"`
	PreSharedKey      string    `xml:"preSharedKey" gorm:"column:pre_shared_key"`
	Qr                string    `xml:"qr,omitempty" gorm:"column:qr_code"`
	Type              string    `xml:"requestType,omitempty" gorm:"column:request_type"`
	PrecheckQueueID   int       `xml:"precheckQueueID" gorm:"column:precheck_queue_id"`
	HashSum           string    `xml:"hashSum" gorm:"column:hash_sum"`
	CreatedAt         time.Time `xml:"-" gorm:"column:create_time;type:timestamp"`
	IPAddress         string    `xml:"-" gorm:"column:ip_address"`
	GateWay           string    `xml:"gateWay" gorm:"column:gate_way; default:'MEGAFON'"`
}

//
////SaveModel saves PaymentReqRawXML model in db
func (payReq *PaymentReqRawXML) SaveModel() {
	db := db.GetPostgresDb()
	db.Create(payReq)
}

//TableName for changing struct name to db name
func (payReq PaymentReqRawXML) TableName() string {
	return "tb_request_log"
}

//
////FindByID finds transaction by id
func (payReq *PaymentReqRawXML) FindByID(ID uint) {
	db := db.GetPostgresDb()
	db.Find(payReq, ID)
}

// PaymentResXML структура ...
type PaymentResXML struct {
	XMLName         xml.Name `xml:"response"`
	Command         string   `xml:"command"`
	Vendor          int      `xml:"vendor,omitempty"`
	Amount          float64  `xml:"amount"`
	Result          int      `xml:"result"`
	Reason          string   `xml:"reason"`
	QueueID         int      `xml:"queueID"`
	AccountPayer    string   `xml:"accountPayer"`
	AccountReceiver string   `xml:"accountReceiver,omitempty"`
	Qr              string   `xml:"qr,omitempty" gorm:"column:qr_code"`
	HashSum         string   `xml:"hashSum"`
}

// TableTransaction model for saving information about transactions to db table
type TableTransaction struct {
	ID                  int              `gorm:"column:id"`
//	Request             PaymentReqRawXML `gorm:"foreignkey:RequestID"`
	RequestID           uint             `gorm:"column:request_id"`
	Vendor              int              `gorm:"column:vendor_id"`
	Qr                  string           `gorm:"column:qr_code"`
	AccountPayer        string           `gorm:"column:account_payer"`
	AccountReceiver     string           `gorm:"colimn:account_receiver"`
	StateID             string           `gorm:"column:state_id"`
	CreatedAt           time.Time        `gorm:"column:create_time"`
	PaymentID           int64            `gorm:"column:payment_id"`
	Route               int              `gorm:"column:route"`
	SendToCft           int              `gorm:"column:send_to_cft"`
	AggregatorStatus    string           `gorm:"column:aggregator_status"`
	AggregatorTransTime time.Time        `gorm:"column:agr_trans_time"`
}

// struct for view
type ViewTransaction struct {
	ID                int64     `xml:"id"`
	RequestId         int64     `xml:"request_id"`
	PaymentID         int64     `xml:"payment_id"`
	PreCheckQueueID   int64     `xml:"pre_check_queue_id"`
	Vendor            int       `xml:"vendor"`
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
	QrCode            string    `xml:"qr_code"`
	NameRus           string    `xml:"name_rus"`
	TimeDiff          time.Time `xml:"time_diff"`
}

//type ResponseViewTransactions struct {
//	Count               int64
//	ViewTransactionList []ViewTransaction
//}
type ResponseViewTransactions struct {
	Page                int64             `json:"page"`
	TotalPage           int64             `json:"totalPage"`
	URL                 string            `json:"url"`
	ViewTransactionList []ViewTransaction `json:"data"`
}

type ResponseTransactions struct {
	Page            int64              `json:"page"`
	TotalPage       int64              `json:"totalPage"`
	URL             string             `json:"url"`
	TransactionList []TableTransaction `json:"data"`
}

//
func GetTableTransactionsCount(transaction TableTransaction, time helperfunc.TimeInterval) (Transaction ResponseTransactions) {
	postgresDb := db.GetPostgresDb()
	err := postgresDb.Table(`tb_transaction`).Select("tb_transaction.*").
		Where(&transaction).Where(`create_time > ? and create_time < ?`, time.From, time.To).
		Count(&Transaction.TotalPage).Error
	if err != nil {
		//Transaction.Error = err
		fmt.Println("can't take from db")
	}
	return Transaction
}

func GetTableTransactions(transaction TableTransaction, transactionSlice *ResponseTransactions, time helperfunc.TimeInterval, page int64) (Transaction *ResponseTransactions) {
	postgresDb := db.GetPostgresDb()
	err := postgresDb.Table(`tb_transaction`).Select("tb_transaction.*").
		Where(&transaction).Where(`create_time > ? and create_time < ?`, time.From, time.To).
		Limit(100).Offset(page * 100).Order("create_time desc").Find(&transactionSlice.TransactionList).Error
	if err != nil {
		//		transactionSlice.Error = err
		fmt.Println("can't take from db")
	}
	return
}

func GetTableTransactionsByID(id int64) (Transaction TableTransaction) {
	postgresDb := db.GetPostgresDb()
	fmt.Println("i am id = ", id)
	err := postgresDb.Table(`tb_transaction`).
		Where(`id = ?`, id).First(&Transaction).Error
	if err != nil {
		//		transactionSlice.Error = err
		fmt.Println("can't take from db")
	}
	return
}

func ChangeNewStateID(Id int, StateID string) (error, TableTransaction){
	postgresDb := db.GetPostgresDb()
	Transaction := GetTableTransactionsByID(int64(Id))
	var newtrans TableTransaction
	newtrans.StateID = StateID
	err := postgresDb.Model(&Transaction).Update(newtrans).Error
	if err != nil {
		return err, newtrans
	}
	return nil, newtrans
}

//
func GetViewTransactionsByID(id int64) (Transaction ViewTransaction) {
	postgresDb := db.GetPostgresDb()
	fmt.Println("i am id = ", id)
	err := postgresDb.Table(`view_transaction`).
		Where(`id = ?`, id).First(&Transaction).Error
	if err != nil {
		//		transactionSlice.Error = err
		fmt.Println("can't take from db")
	}
	return
}

////Get From view_transaction
func GetViewTransactions(transaction ViewTransaction, transactionSlice *ResponseViewTransactions, time helperfunc.TimeInterval, page int64) (Transaction *ResponseViewTransactions) {
	postgresDb := db.GetPostgresDb()
	err := postgresDb.Table(`view_transaction`).Select("view_transaction.*").
		Where(&transaction).Where(`create_time > ? and create_time < ?`, time.From, time.To).
		Limit(100).Offset(page * 100).Order("create_time desc").Find(&transactionSlice.ViewTransactionList).Error
	if err != nil {
		//		transactionSlice.Error = err
		fmt.Println("can't take from db")
	}
	return
}

func GetViewTransactionsCount(transaction ViewTransaction, time helperfunc.TimeInterval) (Transaction ResponseViewTransactions) {
	postgresDb := db.GetPostgresDb()
	err := postgresDb.Table(`view_transaction`).Select("view_transaction.*").
		Where(&transaction).Where(`create_time > ? and create_time < ?`, time.From, time.To).
		Count(&Transaction.TotalPage).Error
	if err != nil {
		//Transaction.Error = err
		fmt.Println("can't take from db")
	}
	return Transaction
}

////SaveModel saves TableTransaction model in db
func (tableTransaction *TableTransaction) SaveModel() {
	db := db.GetPostgresDb()
	if err := db.Create(tableTransaction).Error; err != nil {
		logrus.Warn("TransactionSaveModel ", err)
	}
}

//
//// UpdateModel updates transaction model
func (tableTransaction *TableTransaction) UpdateModel() {
	db := db.GetPostgresDb()
	if err := db.Save(tableTransaction).Error; err != nil {
		logrus.Warnln("UpdateStatusTransaction ", err)
	}
}

//
//FindByID finds transaction by id
func (tableTransaction *TableTransaction) FindByID(queueID int) {
	db := db.GetPostgresDb()
	db.Find(tableTransaction, queueID)
}

//
//// IsPaymentAccepted to check whether payment was accepted or not
func (tableTransaction *TableTransaction) IsPaymentAccepted(precheckQueueID int) bool {
	db := db.GetPostgresDb()
	db.Joins("INNER JOIN tb_request_log as rl ON rl.id = tb_transaction.request_id AND rl.precheck_queue_id = ?", precheckQueueID).First(tableTransaction)
	if tableTransaction.ID > 0 {
		return true
	}
	return false
}

////
//// Delete transaction info from table
func (tableTransaction *TableTransaction) Delete() {
	db := db.GetPostgresDb()
	db.Delete(tableTransaction)
}

//
//// NotHandledTransactions returns all not handled transactions
func (tableTransaction *TableTransaction) NotHandledTransactions() []TableTransaction {
	var transactions []TableTransaction
	db := db.GetPostgresDb()
	db.Where("vendor_id <> ? AND (state_id = ? OR state_id = ?) AND payment_id <> 0", "0", "Accepted", "Sent").Find(&transactions)
	return transactions
}

//// NotSendedPayments gets all transactions where not sended to payment system
func (tableTransaction *TableTransaction) NotSendedPayments() []TableTransaction {
	var transactions []TableTransaction
	db := db.GetPostgresDb()
	db.Where("vendor_id <> ? AND state_id = ? AND payment_id = ?", "0", "Accepted", "0").Find(&transactions)
	return transactions
}

func (tableTransaction *TableTransaction) GetTransactionsForRefund() []TableTransaction {
	var transactions []TableTransaction
	db := db.GetPostgresDb()
	db.Joins("INNER JOIN tb_request_log as rl ON rl.id = tb_transaction.request_id").Where("state_id = ? AND rl.request_type IN (?,?)", "Failed", "card_online", "card_onsite").Find(&transactions)
	return transactions
}

//Suma transaction summ
type Suma struct {
	Sum float64
}

//
// GetMonthTransSum gets transaction sum by account payer (need for limits)
func GetMonthTransSum(acountPayer string) float64 {
	db := db.GetPostgresDb()
	now := time.Now()
	beginOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.Local)
	// var sum float64
	var suma Suma
	db.Raw(`SELECT sum(rl.amount_with_commiss) as sum
	FROM public.tb_transaction as tr
	INNER JOIN public.tb_request_log as rl on tr.request_id = rl.id
	where tr.account_payer = ? and (tr.state_id = 'Processed' or tr.state_id = 'Accepted' or tr.state_id = 'Sent') and (tr.create_time ::date between ? and ?)
	GROUP BY (tr.account_payer)`, acountPayer, beginOfMonth.Format("2006-01-02"), now.Format("2006-01-02")).Scan(&suma)
	return suma.Sum
}

//TableName for changing struct name to db name
func (tableTransaction TableTransaction) TableName() string {
	return "tb_transaction"
}
func (*ViewTransaction) TableName() string {
	return "view_transaction"
}

// Request For Cancel transaction
func CancelMegafonTransaction(link string) bool {
	resp, err := http.Get(link)
	if err != nil {
		return false
	}
	if resp.StatusCode == http.StatusOK {
		return true
	}
	return false
}