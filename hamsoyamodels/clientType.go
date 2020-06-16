package hamsoyamodels

import "MF/db"

type HamsoyaClientType struct {
	Id          int64  `xml:"id"`
	Code        string `xml:"code"`
	Name        string `xml:"name"`
	Description string `xml:"description"`
}

func (*HamsoyaClientType) TableName() string {
	return "client_type"
}

func (a *HamsoyaClientType) Count() (count int64) {
	postgresDb := db.GetHamsoyaPostgresDb()
	postgresDb.Table("client_type").Count(&count)
	//	postgresDb.Where("id > 0").First(&a)
	return
}
