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
	err := postgresDb.Model(&HamsoyaStatus).Update(Status).Error
	if err != nil {
		return *HamsoyaStatus, err
	}
	return *HamsoyaStatus, nil
}
