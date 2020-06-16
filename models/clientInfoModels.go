package models

import (
	"MF/db"
	"MF/helperfunc"
	"encoding/xml"
	"github.com/sirupsen/logrus"
	"regexp"
	"strings"
	"time"
)

// ClientInfo структура ...
type ClientInfo struct {
	ClientID            int       `xml:"-" gorm:"column:id"`
	Phone               string    `xml:"phone,omitempty" gorm:"column:phone"`
	Name                string    `xml:"name"  binding:"required" gorm:"column:name"`
	BirthDate           time.Time `xml:"-" gorm:"column:birth_date"`
	INN                 string    `xml:"inn" gorm:"column:inn"`
	PassportSeries      string    `xml:"passportSeries,omitempty" gorm:"column:doc_series"`
	PassportNumber      string    `xml:"passportNumber,omitempty" gorm:"column:doc_number"`
	PassportIssuingAuth string    `xml:"passportIssuingAuth,omitempty" gorm:"column:doc_iss_auth"`
	PassportIssuingDate time.Time `xml:"-" gorm:"column:doc_iss_date"`
	Address             string    `xml:"address,omitempty" gorm:"column:address"`
	Nationality         string    `xml:"nationality,omitempty" gorm:"column:nationality"`
	Sex                 string    `xml:"sex,omitempty" gorm:"column:sex"`
	IsActive            bool      `xml:"isActive" gorm:"column:active; default:true"`
	CreateDate			time.Time	`xml:"-" gorm:"column:create_date; default: CURRENT_TIMESTAMP"`
	IsIdentified        bool      `xml:"isIdentified" gorm:"column:identify; default:true"`
	IsBlackList         bool      `gorm:"column:black_list"`
	SendToCft           bool      `gorm:"column:send_to_cft"`
}

type ClientInfoRequest struct {
	RawXML
	ClientInfo
	SBirthDate           string `xml:"birthDate,omitempty"`
	SPassportIssuingDate string `xml:"passportIssuingDate,omitempty"`
}

type ClientInfoResponse struct {
	XMLName xml.Name `xml:"response"`
	ClientInfo
	SBirthDate           string `xml:"birthDate"`
	SPassportIssuingDate string `xml:"passportIssuingDate"`
	Result               int    `xml:"result"`
	Reason               string `xml:"reason"`
	HashSum              string `xml:"hashSum"`
}

//SaveModel saves ClientInfo model in db
func (clientInfo *ClientInfo) SaveMode() string {

	db := db.GetPostgresDb()
	if err := db.Create(&clientInfo).Error; err != nil {
		logrus.Println("ClientInfoSaveModel ", err.Error())

		if strings.Contains(err.Error(), "pq: duplicate key value violates unique constraint") {

			re := regexp.MustCompile("tb_client_(.*?)_key")
			match := re.FindStringSubmatch(err.Error())
			if len(match) > 0 {
				return "Duplicate param : " + match[1]
			}

		}

		return "Missing required params"
	}
	return ""
}

//TableName for changing struct name to db name
func (clientInfo ClientInfo) TableName() string {
	return "tb_client"
}

//Get ClientInfo by phone
func GetClientInfoById(id string) (clientInfo ClientInfo) {
	if err := db.GetPostgresDb().Where("id = ?", id).First(&clientInfo).Error; err != nil {
		logrus.Println("GetClientInfoById By Phone ", err)
	}
	return clientInfo
}

type ResponseClientsInfo struct {
	Error          error        `json:"error"`
	Page           int64        `json:"page"`
	TotalPage      int64        `json:"totalPage"`
	URL            string       `json:"url"`
	ClientInfoList []ClientInfo `json:"clientInfoList"`
}

//
//func GetClients(client ClientInfo, size, page int64) (clientsSlice ResponseClientsInfo) {
//
//	if err := db.GetPostgresDb().Where(&client).Limit(size).Offset(page * size).Order("create_date desc").Find(&clientsSlice.ClientInfoList).Error; err != nil {
//		clientsSlice.Error = err
//		logrus.Println(" ", err)
//	}
//	if err := db.GetPostgresDb().Table("tb_client").Where(&client).Count(&clientsSlice.TotalPage).Error; err != nil {
//		clientsSlice.Error = err
//		logrus.Println(" ", err)
//	}
//	return
//}
func GetClients(client ClientInfo, clientsSlice *ResponseClientsInfo, time helperfunc.TimeInterval, page int64) (clientsSliceOver *ResponseClientsInfo) {

	if err := db.GetPostgresDb().Where(&client).Where(`create_date > ? and create_date < ?`, time.From, time.To).Limit(100).Offset(page * 100).Order("create_date desc").Find(&clientsSlice.ClientInfoList).Error; err != nil {
		clientsSlice.Error = err
		logrus.Println(" ", err)
	}
	return
}

func GetClientsCount(client ClientInfo, time helperfunc.TimeInterval) (clientsSlice ResponseClientsInfo) {
	if err := db.GetPostgresDb().Table("tb_client").Where(&client).Where(`create_date > ? and create_date < ?`, time.From, time.To).Count(&clientsSlice.TotalPage).Error; err != nil {
		clientsSlice.Error = err
		logrus.Println(" ", err)
	}
	return
}