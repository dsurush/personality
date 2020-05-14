package hamsoyamodels

import (
	"MF/db"
	"time"
)

type HamsoyaConfig struct {
	Id          int64     `xml:"id"`
	Code        string    `xml:"code"`
	Value       int64     `xml:"value"`
	Description string    `xml:"description"`
	CreateDate  time.Time `xml:"create_date"`
	ValueStr    string    `xml:"value_str"`
}

func (*HamsoyaConfig) TableName() string{
	return "config"
}

func GetHamsoyaConfig(config HamsoyaConfig, rows, pages int64) (HamsoyaConfig []HamsoyaConfig, err error) {
	pages--
	if pages < 0 {
		pages = 0
	}
	postgresDb := db.GetHamsoyaPostgresDb()
	if err := postgresDb.Where(&config).Limit(rows).Offset(pages * rows).Find(&HamsoyaConfig).Error; err != nil {
		return HamsoyaConfig, err
	}
	return
}

func (HamsoyaConfig *HamsoyaConfig) Save() (HamsoyaConfig, error){
	postgresDb := db.GetHamsoyaPostgresDb()
	err:= postgresDb.Create(&HamsoyaConfig).Error
	if err != nil {
		return *HamsoyaConfig, err
	}
	return *HamsoyaConfig, nil
}

func (HamsoyaConfig *HamsoyaConfig) Update(Config HamsoyaConfig) (HamsoyaConfig, error) {
	postgresDb := db.GetHamsoyaPostgresDb()
	err := postgresDb.Model(&HamsoyaConfig).Update(Config).Error
	if err != nil {
		return *HamsoyaConfig, err
	}
	return *HamsoyaConfig, nil
}

func GetHamsoyaConfigById(id int64) (HamsoyaConfig HamsoyaConfig, err error) {
	postgresDb := db.GetHamsoyaPostgresDb()
	if err := postgresDb.Where("id = ?", id).First(&HamsoyaConfig).Error; err != nil {
		return HamsoyaConfig, err
	}
	return
}