package hamsoyamodels

import (
	"MF/db"
	"fmt"
	"time"
)

type HamsoyaAccount struct {
	Id         int64     `xml:"id"`
	AccNum     string    `xml:"acc_num"`
	Name       string    `xml:"name"`
	ClientId   int64     `xml:"client_id"`
	Overdraft  float64   `xml:"overdraft"`
	CurrencyId int64     `xml:"currency_id"`
	TypeId     int64     `xml:"type_id"`
	IsDefault  bool      `xml:"is_default"`
	IsActive   bool      `xml:"is_active"`
	CreateDate time.Time `xml:"create_date"`
	LastUpdate time.Time `xml:"last_update"`
	Saldo      float64   `xml:"saldo"`
}

func (*HamsoyaAccount) TableName() string {
	return "accounts"
}

type ResponseHamsoyaAccount struct {
	Error              error
	Count              int64
	HamsoyaAccountList []HamsoyaAccount
}

///TODO: DELETE ME I AM TEST FUNCTION
func TESTTIME(time string) (HamsoyaAccount ResponseHamsoyaAccount) {
	postgresDb := db.GetHamsoyaPostgresDb()
	if err := postgresDb.Where("create_date > ?", time).Offset(0).Find(&HamsoyaAccount.HamsoyaAccountList).Count(&HamsoyaAccount.Count).Error; err != nil {
		HamsoyaAccount.Error = err
	}
	return
}

//TODO: DELETE ME TOO, I AM TEST FOR INTERVAL
func GetHamsoyaAccountsTEST(hamsoyaAccount HamsoyaAccount, rows, pages int64, leftTime string) (HamsoyaAccount ResponseHamsoyaAccount) {
	fmt.Println("TETST")
	pages--
	if pages < 0 {
		pages = 0
	}
	postgresDb := db.GetHamsoyaPostgresDb()
	if err := postgresDb.Where(&hamsoyaAccount).Where("create_date > ?", leftTime).Limit(rows).Offset(rows * pages).Find(&HamsoyaAccount.HamsoyaAccountList).Count(&HamsoyaAccount.Count).Error; err != nil {
		HamsoyaAccount.Error = err
	}
	return
}

//

func GetHamsoyaAccountById(id int64) (HamsoyaAccount HamsoyaAccount, err error) {
	postgresDb := db.GetHamsoyaPostgresDb()
	if err := postgresDb.Where("id = ?", id).First(&HamsoyaAccount).Error; err != nil {
		return HamsoyaAccount, err
	}
	return HamsoyaAccount, nil
}

func GetHamsoyaAccounts(hamsoyaAccount HamsoyaAccount, rows, pages int64) (HamsoyaAccount ResponseHamsoyaAccount) {
	pages--
	if pages < 0 {
		pages = 0
	}
	postgresDb := db.GetHamsoyaPostgresDb()
	if err := postgresDb.Where(&hamsoyaAccount).Limit(rows).Offset(rows * pages).Find(&HamsoyaAccount.HamsoyaAccountList).Count(&HamsoyaAccount.Count).Error; err != nil {
		HamsoyaAccount.Error = err
	}
	return
}
