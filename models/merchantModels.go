package models

import (
	"MF/db"
	"MF/helperfunc"
	"encoding/xml"
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
)

var getMerchProc = "BEGIN ibs.z$dbo_doocat_megafon_fin_lib.GET_MERCHANTS(:merchants, :errorCode, :errorDescription); END;"

//Merchant struct for tb_merchant
type Merchant struct {
	ID           uint      `gorm:"column:id" xml:"id"`
	HumoOnlineID int64     `gorm:"column:humo_online_id" xml:"-"`
	NameENG      string    `gorm:"column:name_eng" xml:"latName"`
	NameRUS      string    `gorm:"column:name_rus" xml:"name"`
	QrCode       string    `gorm:"column:qr_code" xml:"qr"`
	QrCodeNew    string    `gorm:"column:qr_code_new" xml:"qr_new"`
	CreateTime   time.Time `gorm:"column:create_time" xml:"-"`
	UpdateTime   time.Time `gorm:"column:update_time" xml:"-"`
}

type MerchantDTO struct {
	ID           uint      `gorm:"column:id" xml:"id"`
	HumoOnlineID int64     `gorm:"column:humo_online_id" xml:"-"`
	NameENG      string    `gorm:"column:name_eng" xml:"latName"`
	NameRUS      string    `gorm:"column:name_rus" xml:"name"`
	QrCode       string    `gorm:"column:qr_code" xml:"qr"`
	QrCodeNew    string    `gorm:"column:qr_code_new" xml:"qr_new"`
	//CreateTime   time.Time `gorm:"column:create_time" xml:"-"`
	//UpdateTime   time.Time `gorm:"column:update_time" xml:"-"`
}

type ResponseMerchants struct {
	Error          error        `json:"error"`
	Page           int64        `json:"page"`
	TotalPage      int64        `json:"totalPage"`
	URL            string       `json:"url"`
	MerchantList []Merchant `json:"data"`
}

//TableName for changing struct name to db name
func (merchant *Merchant) TableName() string {
	return "tb_ho_merchant"
}

// Get merchantlist by page and rowssize
func GetMerchants(merchant Merchant, merchantsSlice *ResponseMerchants, time helperfunc.TimeInterval, page int64) (merchantsSliceOver *ResponseMerchants) {
	postgresDb := db.GetPostgresDb()
	if err := postgresDb.Where(&merchant).Where(`create_time > ? and create_time < ?`, time.From, time.To).Limit(100).Offset(page * 100).Order("create_time desc").Find(&merchantsSlice.MerchantList).Error; err != nil {
		merchantsSlice.Error = err
		fmt.Println("Can't get Merchats from db")
	}
	return
}

// Get Merchants Count
func GetMerchantsCount(merchant Merchant, time helperfunc.TimeInterval) (merchantsSlice ResponseMerchants) {
	if err := db.GetPostgresDb().Table("tb_ho_merchant").Where(&merchant).Where(`create_time > ? and create_time < ?`, time.From, time.To).Count(&merchantsSlice.TotalPage).Error; err != nil {
		merchantsSlice.Error = err
		logrus.Println(" ", err)
	}
	return
}

//Get merchant by ID
func (merchant *Merchant) GetMerchantById(id int64) (merchantsById Merchant) {
	postgresDb := db.GetPostgresDb()
	postgresDb.First(&merchantsById, id)
	return merchantsById
}

//
func (Merchant *Merchant) Update(merchant Merchant) Merchant {
	postgresDb := db.GetPostgresDb()
	if err := postgresDb.Model(&Merchant).Updates(merchant).Error; err != nil {
		fmt.Println("can't change this pole")
	}
	return *Merchant
}

//
func (merchant *Merchant) save() {
	db := db.GetPostgresDb()
	db.Where(Merchant{HumoOnlineID: merchant.HumoOnlineID}).Assign(Merchant{QrCode: merchant.QrCode}).FirstOrCreate(merchant)
}

//
//Find finds merchant by qr
func (merchant *Merchant) Find(qr string) {
	db := db.GetPostgresDb()
	//db.Where(Merchant{QrCode: qr}).Find(merchant)
	db.Where(Merchant{QrCode: qr}).Or(Merchant{QrCodeNew: qr}).Find(merchant)
}

// MerchantListResXML ...
type MerchantListResXML struct {
	XMLName   xml.Name   `xml:"response"`
	Command   string     `xml:"command"`
	Merchants []Merchant `xml:"merchant"`
	HashSum   string     `xml:"hashSum"`
}
