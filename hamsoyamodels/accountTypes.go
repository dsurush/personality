package hamsoyamodels

import "MF/db"

type HamsoyaAccountType struct {
	Id          int64  `xml:"id"`
	Type        string `xml:"type"`
	Code        string `xml:"code"`
	Name        string `xml:"name"`
	Prefix      int64  `xml:"prefix"`
	Description string `xml:"description"`
}

func (*HamsoyaAccountType) TableName() string {
	return "account_type"
}

type ResponseHamsoyaAccountTypes struct {
	Error                  error                `json:"error"`
	Page                   int64                `json:"page"`
	TotalPage              int64                `json:"totalPage"`
	URL                    string               `json:"url"`
	HamsoyaAccountTypeList []HamsoyaAccountType `json:"data"`
}

func GetHamsoyaAccountTypeById(id int64) (AccountType HamsoyaAccountType, err error) {
	postgresDb := db.GetHamsoyaPostgresDb()
	postgresDb.Where("id = ?", id).First(&AccountType)
	if err := postgresDb.Where("id = ?", id).First(&AccountType).Error; err != nil {
		return AccountType, err
	}
	return AccountType, nil
}

func GetHamsoyaAccountTypes(accountType HamsoyaAccountType, rows, pages int64) (AccountTypes ResponseHamsoyaAccountTypes) {
	pages--
	if pages < 0 {
		pages = 0
	}
	postgresDb := db.GetHamsoyaPostgresDb()
	if err := postgresDb.Where(&accountType).Limit(rows).Offset(rows * pages).Find(&AccountTypes.HamsoyaAccountTypeList).Error; err != nil {
		AccountTypes.Error = err
	}
	AccountTypes.Page = 0
	AccountTypes.TotalPage = 0
	return
}

func (HamsoyaAccountType *HamsoyaAccountType) Save() (HamsoyaAccountType, error) {
	postgresDb := db.GetHamsoyaPostgresDb()
	err := postgresDb.Create(&HamsoyaAccountType).Error
	if err != nil {
		return *HamsoyaAccountType, err
	}
	return *HamsoyaAccountType, nil
}

func (HamsoyaAccountType *HamsoyaAccountType) Update(AccountType HamsoyaAccountType) (HamsoyaAccountType, error) {
	postgresDb := db.GetHamsoyaPostgresDb()
	err := postgresDb.Model(&HamsoyaAccountType).Update(AccountType).Error
	if err != nil {
		return *HamsoyaAccountType, err
	}
	return *HamsoyaAccountType, nil
}
