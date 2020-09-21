package hamsoyamodels

import "MF/db"

type HamsoyaStatus struct {
	Id           int64  `xml:"id" gorm:"column:id"`
	Name         string `xml:"name" gorm:"column:name"`
	Code         string `xml:"code" gorm:"column:code"`
	ExtCode      string `xml:"ext_code" gorm:"column:ext_code"`
	Description  string `xml:"description" gorm:"column:description"`
	ResultCode   int64  `xml:"result_code" gorm:"column:result_code"`
	Final        bool   `xml:"final" gorm:"column:final"`
	IsAmountHold bool   `xml:"is_amount_hold" gorm:"column:is_amount_hold"`
}

func (*HamsoyaStatus) TableName() string {
	return "status"
}

type ResponseHamsoyaStatuses struct {
	Error             error           `json:"error"`
	Page              int64           `json:"page"`
	TotalPage         int64           `json:"totalPage"`
	URL               string          `json:"url"`
	HamsoyaStatusList []HamsoyaStatus `json:"data"`
}

func GetHamsoyaStatusById(id int64) (HamsoyaStatus HamsoyaStatus, err error) {
	postgresDb := db.GetHamsoyaPostgresDb()
	postgresDb.Where("id = ?", id).First(&HamsoyaStatus)
	if err := postgresDb.Where("id = ?", id).First(&HamsoyaStatus).Error; err != nil {
		return HamsoyaStatus, err
	}
	return HamsoyaStatus, nil
}

func GetHamsoyaStatuses(hamsoyaStatus HamsoyaStatus, pages int64) (HamsoyaStatus ResponseHamsoyaStatuses) {
	pages--
	if pages < 0 {
		pages = 0
	}
	postgresDb := db.GetHamsoyaPostgresDb()
	if err := postgresDb.Where(&hamsoyaStatus).Limit(100).Offset(100 * pages).Find(&HamsoyaStatus.HamsoyaStatusList).Error; err != nil {
		HamsoyaStatus.Error = err
	}
	HamsoyaStatus.TotalPage = 0
	HamsoyaStatus.Page = 0
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
	err := postgresDb.Model(&HamsoyaStatus).Update(Status).Update(Status.IsAmountHold).Error
	if err != nil {
		return *HamsoyaStatus, err
	}
	return *HamsoyaStatus, nil
}
