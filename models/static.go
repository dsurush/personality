package models

import (
	"MF/db"
	"MF/helperfunc"
	"fmt"
	"log"
	"sync"
	"time"
)

type StaticByAggregator struct {
	Aggregator string  `json:"aggregator"`
	Amount     float64 `json:"amount"`
	Commiss    float64 `json:"commiss"`
}

type StaticSumOfAmount struct {
	Amount float64 `json:"amount"`
	Count int64 `json:"count"`
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
	StaticSumOfAmount      StaticSumOfAmount           `json:"staticSumOfAmount"`
	ErrSumOfAmount         bool                        `json:"errSumOfAmount"`
	ShortListOfMaxPayer    []StaticShortListOfMaxPayer `json:"shortListOfMaxPayer"`
	ErrShortListOfMaxPayer bool                        `json:"errShortListOfMaxPayer"`
	VendorsTopList         []StaticVendorsTopList      `json:"vendorsTopList"`
	ErrVendorsTopList      bool                        `json:"errVendorsTopList"`
}
func GetMegafonStatic(interval helperfunc.TimeInterval) (MegafonStatic MegafonStatic) {
	defer timeTrack(time.Now(), "Get Megafon Static")
	postgresDb := db.GetPostgresDb()
	fmt.Println(interval.From)
	fmt.Println(interval.To)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := postgresDb.Table("view_transaction").Select("aggregator, sum(amount) as amount, sum(amount_with_commiss - amount) as commiss").
			Where("create_time > ? and create_time < ?", interval.From, interval.To).
			Group("aggregator").Order("amount desc").Find(&MegafonStatic.ByAggregator).Error; err != nil {
			MegafonStatic.ErrByAggregator = true
			fmt.Println("yes4")
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := postgresDb.Table("view_transaction").Select("account_payer, sum(amount) as amount").
			Where("create_time > ? and create_time < ?", interval.From, interval.To).
			Group("account_payer").Order("amount desc").Limit(10).Find(&MegafonStatic.ShortListOfMaxPayer).Error; err != nil {
			MegafonStatic.ErrShortListOfMaxPayer = true
			fmt.Println("yes3")
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		//inter := StaticSumOfAmount{}
		if err := postgresDb.Table("view_transaction").Select("sum(amount) as amount").
			Where("create_time > ? and create_time < ?", interval.From, interval.To).
			Find(&MegafonStatic.StaticSumOfAmount).Error; err != nil {
			MegafonStatic.ErrSumOfAmount = true
			fmt.Println("yes2")
		}
		//MegafonStatic.StaticSumOfAmount = inter.Amount
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := postgresDb.Table("view_transaction").Select("count(id) as count").Where(`id > 0`).
			Where("create_time > ? and create_time < ?", interval.From, interval.To).
			Count(&MegafonStatic.StaticSumOfAmount.Count).Error; err != nil {
			MegafonStatic.ErrSumOfAmount = true
			fmt.Println("yes1")
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := postgresDb.Table("public.view_transaction, public.tb_vendor_category").
			Where("create_time > ? and create_time < ?", interval.From, interval.To).
			Select("tb_vendor_category.name_rus, sum(view_transaction.amount) as amount").Where("tb_vendor_category.id = view_transaction.vendor_id").
			Group("tb_vendor_category.name_rus").Order("amount desc").Find(&MegafonStatic.VendorsTopList).Error; err != nil {
			MegafonStatic.ErrVendorsTopList = true
			fmt.Println("yes")
		}
	}()
	wg.Wait()
	return
}
func timeTrack(start time.Time, name string) {
	log.Printf("%s took %s", name, time.Since(start))
}