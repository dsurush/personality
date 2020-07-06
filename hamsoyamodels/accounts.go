package hamsoyamodels

import (
	"MF/db"
	"MF/helperfunc"
	"github.com/sirupsen/logrus"
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
	Error          error        `json:"error"`
	Page           int64        `json:"page"`
	TotalPage      int64        `json:"totalPage"`
	URL            string       `json:"url"`
	HamsoyaAccountList []HamsoyaAccount `json:"hamsoya_account_list"`
}

func GetHamsoyaAccountById(id int64) (HamsoyaAccount HamsoyaAccount, err error) {
	postgresDb := db.GetHamsoyaPostgresDb()
	if err := postgresDb.Where("id = ?", id).First(&HamsoyaAccount).Error; err != nil {
		return HamsoyaAccount, err
	}
	return HamsoyaAccount, nil
}

func GetHamsoyaAccounts(account HamsoyaAccount, accountSlice *ResponseHamsoyaAccount, time helperfunc.TimeInterval, page int64) (AccountSliceOver *ResponseHamsoyaAccount) {

	if err := db.GetHamsoyaPostgresDb().Where(&account).Where(`create_date > ? and create_date < ?`, time.From, time.To).Limit(100).Offset(page * 100).Order("create_date desc").Find(&accountSlice.HamsoyaAccountList).Error; err != nil {
		accountSlice.Error = err
		logrus.Println(" ", err)
	}
	return
}

func GetHamsoyaAccountsCount(account HamsoyaAccount, time helperfunc.TimeInterval) (accountsSlice ResponseHamsoyaAccount) {

	if err := db.GetHamsoyaPostgresDb().Table("accounts").Where(&account).Where(`create_date > ? and create_date < ?`, time.From, time.To).Count(&accountsSlice.TotalPage).Error; err != nil {
		accountsSlice.Error = err
		logrus.Println(" ", err)
	}
	return
}