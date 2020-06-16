package hamsoyamodels

import "MF/db"

type HamsoyaStatus struct {
	Id           int64  `xml:"id"`
	Name         string `xml:"name"`
	Code         string `xml:"code"`
	ExtCode      string `xml:"ext_code"`
	Description  string `xml:"description"`
	ResultCode   int64  `xml:"result_code"`
	Final        bool   `xml:"final"`
	IsAmountHold bool   `xml:"is_amount_hold"`
}

func (*HamsoyaStatus) TableName() string {
	return "status"
}

type ResponseHamsoyaStatuses struct {
	Error             error
	Count             int64
	HamsoyaStatusList []HamsoyaStatus
}

func GetHamsoyaStatusById(id int64) (HamsoyaStatus HamsoyaStatus, err error) {
	postgresDb := db.GetHamsoyaPostgresDb()
	postgresDb.Where("id = ?", id).First(&HamsoyaStatus)
	if err := postgresDb.Where("id = ?", id).First(&HamsoyaStatus).Error; err != nil {
		return HamsoyaStatus, err
	}
	return HamsoyaStatus, nil
}

func GetHamsoyaStatuses(hamsoyaStatus HamsoyaStatus, rows, pages int64) (HamsoyaStatus ResponseHamsoyaStatuses) {
	pages--
	if pages < 0 {
		pages = 0
	}
	postgresDb := db.GetHamsoyaPostgresDb()
	if err := postgresDb.Where(&hamsoyaStatus).Limit(rows).Offset(rows * pages).Find(&HamsoyaStatus.HamsoyaStatusList).Count(&HamsoyaStatus.Count).Error; err != nil {
		HamsoyaStatus.Error = err
	}
	return
}

func (HamsoyaStatus *HamsoyaStatus) Save() (HamsoyaStatus, error) {
	postgresDb := db.GetHamsoyaPostgresDb()
	err := postgresDb.Create(&HamsoyaStatus).Error
	if err != nil {
		return *HamsoyaStatus, err
	}
	return *HamsoyaStatus, nil
}

func (HamsoyaStatus *HamsoyaStatus) Update(Status HamsoyaStatus) (HamsoyaStatus, error) {
	postgresDb := db.GetHamsoyaPostgresDb()
	err := postgresDb.Model(&HamsoyaStatus).Update(Status).Error
	if err != nil {
		return *HamsoyaStatus, err
	}
	return *HamsoyaStatus, nil
}
