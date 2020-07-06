package hamsoyamodels

import (
	"MF/db"
	"MF/helperfunc"
	"github.com/sirupsen/logrus"
	"time"
)

type HamsoyaTransactionType struct {
	Id          int64  `xml:"id"`
	Code        string `xml:"code"`
	Name        string `xml:"name"`
	Description string `xml:"description"`
	IsForJob    bool   `xml:"is_for_job"`
	IsPayment   bool   `xml:"is_payment"`
	IsActive    bool   `xml:"is_active"`
}

func (*HamsoyaTransactionType) TableName() string {
	return "transaction_type"
}

func (HamsoyaTransactionType *HamsoyaTransactionType) Save() HamsoyaTransactionType {
	postgresDb := db.GetHamsoyaPostgresDb()
	postgresDb.Create(&HamsoyaTransactionType)
	return *HamsoyaTransactionType
}
func (HamsoyaTransactionType *HamsoyaTransactionType) Update(transactionType HamsoyaTransactionType) HamsoyaTransactionType {
	postgresDb := db.GetHamsoyaPostgresDb()
	postgresDb.Model(&HamsoyaTransactionType).Update(transactionType)
	return *HamsoyaTransactionType
}

type HamsoyaTransaction struct {
	Id            int64     `xml:"id"`
	PreCheckId    int64     `xml:"pre_check_id"`
	StatusId      int64     `xml:"status_id"`
	TypeId        int64     `xml:"type_id"`
	ExtStatusId   int64     `xml:"ext_status_id"`
	ExtTransId    string    `xml:"ext_trans_id"`
	CreateDate    time.Time `xml:"create_date"`
	LastUpdate    time.Time `xml:"last_update"`
	Description   string    `xml:"description"`
	ClientPayerId int64     `xml:"client_payer_id"`
}

type ResponseHamsoyaTransactionsType struct {
	Error          error        `json:"error"`
	Page           int64        `json:"page"`
	TotalPage      int64        `json:"totalPage"`
	URL            string       `json:"url"`
	HamsoyaTransactionTypeList []HamsoyaTransactionType `json:"hamsoya_transaction_type_list"`
}
type ResponseHamsoyaTransactions struct {
	Error          error        `json:"error"`
	Page           int64        `json:"page"`
	TotalPage      int64        `json:"totalPage"`
	URL            string       `json:"url"`
	HamsoyaTransactionList []HamsoyaTransaction `json:"hamsoya_transaction_list"`
}

func (*HamsoyaTransaction) TableName() string {
	return "transactions"
}
func GetHamsoyaTransactionTypeById(id int64) (HamsoyaTransactionType HamsoyaTransactionType, err error) {
	postgresDb := db.GetHamsoyaPostgresDb()
	if err := postgresDb.Where("id = ?", id).First(&HamsoyaTransactionType).Error; err != nil {
		return HamsoyaTransactionType, err
	}
	return
}

func GetHamsoyaTransactionsType(transaction HamsoyaTransactionType, transactionSlice *ResponseHamsoyaTransactionsType, page int64) (transactionTypeSliceOver *ResponseHamsoyaTransactionsType) {
	postgresDb := db.GetHamsoyaPostgresDb()

	if err := postgresDb.Where(&transaction).Limit(100).Offset(100 * page).Order(`id`).Find(&transactionSlice.HamsoyaTransactionTypeList).Error; err != nil {
		transactionSlice.Error = err
		logrus.Println(" ", err)
	}
	return
}

func GetHamsoyaTransactionsTypeCount(transaction HamsoyaTransactionType) (transactionTypeSliceOver ResponseHamsoyaTransactionsType) {
	postgresDb := db.GetHamsoyaPostgresDb()

	if err := postgresDb.Table("transaction_type").Where(&transaction).Count(&transactionTypeSliceOver.TotalPage).Error; err != nil {
		transactionTypeSliceOver.Error = err
		logrus.Println(" ", err)
	}
	return
}

func GetHamsoyaTransactions(transaction HamsoyaTransaction, transactionSlice *ResponseHamsoyaTransactions, time helperfunc.TimeInterval, page int64) (HamsoyaTransactions *ResponseHamsoyaTransactions) {

	postgresDb := db.GetHamsoyaPostgresDb()
	if err := postgresDb.Where(&transaction).Where(`create_date > ? and create_date < ?`, time.From, time.To).Limit(100).Offset(page * 100).Find(&transactionSlice.HamsoyaTransactionList).Error; err != nil {
		transactionSlice.Error = err
	}
	return
}

func GetHamsoyaTransactionsCount(transaction HamsoyaTransaction, time helperfunc.TimeInterval) (HamsoyaTransactions ResponseHamsoyaTransactions) {

	postgresDb := db.GetHamsoyaPostgresDb()
	if err := postgresDb.Table("transactions").Where(&transaction).Where(`create_date > ? and create_date < ?`, time.From, time.To).Count(&HamsoyaTransactions.TotalPage).Error; err != nil {
		HamsoyaTransactions.Error = err
	}
	return
}

func GetHamsoyaTransactionById(id int64) (HamsoyaTransaction HamsoyaTransaction, err error) {
	postgresDb := db.GetHamsoyaPostgresDb()
	if err := postgresDb.Where("id = ?", id).First(&HamsoyaTransaction).Error; err != nil {
		return HamsoyaTransaction, err
	}
	return
}

