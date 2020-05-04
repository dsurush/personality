package models

import (
	"MF/db"
	"encoding/xml"
	"github.com/sirupsen/logrus"
	"time"
)

// VendorListReqRawXML ...
type VendorListReqRawXML struct {
	XMLName   xml.Name  `xml:"request"`
	ID        int       `xml:"-" gorm:"column:id"`
	Command   string    `xml:"command" gorm:"column:command"`
	HashSum   string    `xml:"hashSum" gorm:"column:hash_sum"`
	CreatedAt time.Time `xml:"-" gorm:"column:create_time;type:timestamp"`
	IPAddress string    `xml:"-" gorm:"column:ip_address"`
}
//
////SaveModel saves VendorListReqRawXML model in db
func (vendorListReq *VendorListReqRawXML) SaveModel() {
	db := db.GetPostgresDb()
	db.Create(vendorListReq)
}

//TableName for changing struct name to db name
func (vendorListReq VendorListReqRawXML) TableName() string {
	return "tb_request_log"
}

// VendorListResXML ...
type VendorListResXML struct {
	XMLName xml.Name `xml:"response"`
	Command string   `xml:"command"`
	Vendors []Vendor `xml:"vendor"`
	HashSum string   `xml:"hashSum"`
}

//Vendor ...
type Vendor struct {
	ID           int       `xml:"id" gorm:"column:id"`
	LatName      string    `xml:"latName" gorm:"column:name_eng"`
	Name         string    `xml:"name" gorm:"column:name_rus"`
	CatID        int       `xml:"catID" gorm:"column:category_id"`
	Feept        float64   `xml:"feept" gorm:"column:feept"`
	Prefix       string    `xml:"prefix" gorm:"column:prefix"`
	HumoPayID    int       `xml:"-" gorm:"column:humopay_id"`
	TajPayID     int       `xml:"-" gorm:"column:tajpay_id"`
	ExpressPayID string    `xml:"-" gorm:"column:expresspay_id"`
	AmonatBonkID int       `xml:"-" gorm:"column:amonatbonk_id"`
	HumoPayNewID int       `xml:"-" gorm:"column:humopay_new_id"`
	Route        int       `xml:"-" gorm:"column:route"`
	MinSum       float64   `xml:"minSum" gorm:"column:min_sum"`
	CreateTime   time.Time `xml:"-" gorm:"column:create_time"`
//	Type         string    `xml:"type,omitempty"`
	IsActive     bool      `xml:"-" gorm:"column:is_active"`
}

//TableName for changing struct name to db name
func (vendor *Vendor) TableName() string {
	return "tb_vendor"
}

//Find finds vendor by id
func (vendor *Vendor) Find(id int) {
	db := db.GetPostgresDb()
	db.Find(vendor, "id = ?", id)
}

// FindAll returns slice of vendors
func (vendor *Vendor) FindAll() []Vendor {
	vendors := []Vendor{}
	db := db.GetPostgresDb()
	db.Table("tb_vendor").Select("tb_vendor.*, vc.name_rus as category").
		Joins("inner join tb_vendor_category as vc on vc.id = tb_vendor.category_id").Where("tb_vendor.is_active = true").
		Order("tb_vendor.category_id").Scan(&vendors)
	return vendors
}

func (vendor *Vendor) FindAllVendors() (vendors []Vendor) {
	postgresDb := db.GetPostgresDb()
	postgresDb.Table(`tb_vendor`).Select("tb_vendor.*")
	return vendors
}

func (Vendor *Vendor) FindAllVendorsByPageSize(page int, size int) (vendors []Vendor) {
	postgresDb := db.GetPostgresDb()
	postgresDb.Table(`tb_vendor`).Select("tb_vendor.*").Limit(size).Offset(page * size).Scan(&vendors)
	return vendors
}

func (Vendor *Vendor) Save() Vendor{
	postgresDb := db.GetPostgresDb()
	postgresDb.Create(&Vendor)
	return *Vendor
}
// Create Method for update
func (Vendor *Vendor) Update(vendor Vendor) Vendor {
	postgresDb := db.GetPostgresDb()
	postgresDb.Model(&Vendor).Updates(vendor)
	return *Vendor
}
// Create func for Get Vendor
func GetVendorById(id int64) (vendor Vendor) {
	if err := db.GetPostgresDb().Where("id = ?", id).First(&vendor).Error; err != nil {
		logrus.Println("GetVendorById ", err)
	}
	return vendor
}