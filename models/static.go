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

type MegafonStatic struct {
	ByAggregator           []StaticByAggregator        `json:"byAggregator"`
	ErrByAggregator        bool                        `json:"errByAggregator"`
	SumOfAmount            StaticSumOfAmount           `json:"sumOfAmount"`
	ErrSumOfAmount         bool                        `json:"errSumOfAmount"`
	ShortListOfMaxPayer    []StaticShortListOfMaxPayer `json:"shortListOfMaxPayer"`
	ErrShortListOfMaxPayer bool                        `json:"errShortListOfMaxPayer"`
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
		return
	}
	if err := postgresDb.Table("view_transaction").Select("sum(amount) as Amount").Find(&MegafonStatic.SumOfAmount).Error; err != nil {
		MegafonStatic.ErrSumOfAmount = true
		return
	}
	return
}
