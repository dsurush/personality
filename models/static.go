package models

import "MF/db"

type StaticByAggregator struct {
	Aggregator string  `json:"aggregator"`
	Amount     float64 `json:"amount"`
	Commiss    float64 `json:"commiss"`
}

type StaticSumOfAmount struct {
	Amount float64 `json:"amount"`
}

type StaticShortListOfMaxPayer struct {
	AccountPayer int64   `json:"account_payer"`
	Amount       float64 `json:"amount"`
}

type StaticVendorsTopList struct {
	Name   string  `json:"name" sql:"column:name_rus"`
	Amount float64 `json:"amount"`
}

type MegafonStatic struct {
	ByAggregator           []StaticByAggregator        `json:"byAggregator"`
	ErrByAggregator        bool                        `json:"errByAggregator"`
	StaticSumOfAmount      float64                     `json:"amount"`
	ErrSumOfAmount         bool                        `json:"errSumOfAmount"`
	ShortListOfMaxPayer    []StaticShortListOfMaxPayer `json:"shortListOfMaxPayer"`
	ErrShortListOfMaxPayer bool                        `json:"errShortListOfMaxPayer"`
	VendorsTopList         []StaticVendorsTopList      `json:"vendorsTopList"`
	ErrVendorsTopList      bool                        `json:"err_vendorsTopList"`
}

func GetMegafonStatic() (MegafonStatic MegafonStatic) {
	postgresDb := db.GetPostgresDb()
	if err := postgresDb.Table("view_transaction").Select("aggregator, sum(amount) as amount, sum(amount_with_commiss - amount) as commiss").
		Group("aggregator").Order("amount desc").Find(&MegafonStatic.ByAggregator).Error; err != nil {
		MegafonStatic.ErrByAggregator = true
		return
	}
	if err := postgresDb.Table("view_transaction").Select("account_payer, sum(amount) as amount").
		Group("account_payer").Order("amount desc").Limit(10).Find(&MegafonStatic.ShortListOfMaxPayer).Error; err != nil {
		MegafonStatic.ErrShortListOfMaxPayer = true
		return
	}
	inter := StaticSumOfAmount{}
	if err := postgresDb.Table("view_transaction").Select("sum(amount) as Amount").First(&inter).Error; err != nil {
		MegafonStatic.ErrSumOfAmount = true
		return
	}
	MegafonStatic.StaticSumOfAmount = inter.Amount
	if err := postgresDb.Table("public.view_transaction, public.tb_vendor_category").
		Select("tb_vendor_category.name_rus, sum(view_transaction.amount) as amount").Where("tb_vendor_category.id = view_transaction.vendor_id").
		Group("tb_vendor_category.name_rus").Order("amount desc").Find(&MegafonStatic.VendorsTopList).Error; err != nil {
		MegafonStatic.ErrVendorsTopList = true
		return
	}
	return
}
