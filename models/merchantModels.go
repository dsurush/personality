package models

import (
	"encoding/xml"
	"MF/db"
	"time"
)

var getMerchProc = "BEGIN ibs.z$dbo_doocat_megafon_fin_lib.GET_MERCHANTS(:merchants, :errorCode, :errorDescription); END;"

//SyncPMerchantsWithOMerchants gets CFT merchants and inserts to postgres merchants table if exist update
//func SyncPMerchantsWithOMerchants() {
//	_, _, ses := db.GetOracleDB()
//	stmtProcCall, err := ses.Prep(getMerchProc)
//	defer stmtProcCall.Close()
//	if err != nil {
//		log.Warn(err)
//		return
//	}
//
//	merchants := &ora.Rset{}
//	var errorCode int64
//	var errorDescription string
//
//	_, err = stmtProcCall.Exe(merchants, &errorCode, &errorDescription)
//
//	if err != nil {
//		log.Warn(err)
//		return
//	}
//
//	fieldNames := merchants.ColumnIndex()
//	if merchants.IsOpen() {
//		for merchants.Next() {
//			var merchant Merchant
//			merchant.HumoOnlineID, _ = strconv.ParseInt(merchants.Row[fieldNames["ID_NUMBER"]].(string), 10, 64)
//			merchant.NameRUS = merchants.Row[fieldNames["NAME"]].(string)
//			log.Println(merchant.HumoOnlineID)
//			merchant.CreateTime = merchants.Row[fieldNames["DATE_OPEN"]].(time.Time)
//			if merchant.HumoOnlineID == 2337 || merchant.HumoOnlineID == 2781 || merchant.HumoOnlineID == 4694 || merchant.HumoOnlineID == 4834 || merchant.HumoOnlineID == 5205 || merchant.HumoOnlineID == 6191 {
//				continue
//			}
//
//			if merchants.Row[fieldNames["QR_CODE"]] != nil {
//				reader := bytes.NewReader(merchants.Row[fieldNames["QR_CODE"]].([]byte))
//				qrmatrix, err := qrcode.Decode(reader)
//				if err != nil {
//					log.Println("error while getting value from " + merchant.NameRUS + "`s qr-code " + err.Error())
//					continue
//				}
//				merchant.QrCode = qrmatrix.Content
//				merchant.save()
//			}
//		}
//		if err := merchants.Err(); err != nil {
//			log.Warn(err)
//		}
//	}
//}

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
