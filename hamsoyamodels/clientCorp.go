package hamsoyamodels

type HamsoyaClientCorp struct {
	Id      int64  `xml:"id"`
	Name    string `xml:"name"`
	Address string `xml:"address"`
	INN     string `xml:"inn"`
}

func (*HamsoyaClientCorp) TableName() string {
	return "client_corp"
}
