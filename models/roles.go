package models

type Role struct {
	Id   int64  `xml:"id" gorm:"column:id"`
	Name string `xml:"name" gorm:"column:name"`
}

func (*Role) TableName() string {
	return "tb_roles"
}
