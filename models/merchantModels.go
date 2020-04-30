package models

import (
	"encoding/xml"
	"MF/db"
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

//TableName for changing struct name to db name
func (merchant *Merchant) TableName() string {
	return "tb_ho_merchant"
}

//GetList gets list of merchants
func (merchant *Merchant) GetList() []Merchant {
	merchants := []Merchant{}
	db := db.GetPostgresDb()
	db.Find(&merchants)
	return merchants
}
// Get merchantlist by page and rowssize
func (merchant *Merchant) GetMerchants(size, page int64) (merchants []Merchant) {
	postgresDb := db.GetPostgresDb()
	postgresDb.Table(`tb_ho_merchant`).Select("tb_ho_merchant.*").Limit(size).Offset(page * size).Scan(&merchants)
	return merchants
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
	postgresDb.Model(&Merchant).Updates(merchant)
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
