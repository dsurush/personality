package hamsoyamodels

import (
	"MF/db"
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


func (HamsoyaTransactionType *HamsoyaTransactionType) Save() HamsoyaTransactionType{
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

func GetHamsoyaTransactions(transaction HamsoyaTransaction, size, page int64) (HamsoyaTransactions []HamsoyaTransaction, err error) {
	page--
	if page < 0 {
		page = 0
	}
	postgresDb := db.GetHamsoyaPostgresDb()
	if err := postgresDb.Where(&transaction).Limit(size).Offset(page * size).Find(&HamsoyaTransactions).Error; err != nil {
		return HamsoyaTransactions, err
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