package hamsoyamodels

type HamsoyaTransactionType struct {
	Id int64
	Code string
	Name string
	Description string
	IsForJob bool
	IsPayment bool
	IsActive bool
}

func (*HamsoyaTransactionType) TableName() string {
	return "transaction_type"
}
