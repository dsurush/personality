package hamsoyamodels

import (
	"MF/db"
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
	Error error
	Count int64
	HamsoyaAccountList []HamsoyaAccount
}

func GetHamsoyaAccountById(id int64) (HamsoyaAccount HamsoyaAccount, err error) {
	postgresDb := db.GetHamsoyaPostgresDb()
	if err := postgresDb.Where("id = ?", id).First(&HamsoyaAccount).Error; err != nil{
		return HamsoyaAccount, err
	}
	return HamsoyaAccount, nil
}
