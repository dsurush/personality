package models

type User struct {
	Id       int64  `xml:"id" gorm:"column:id"`
	Name     string `xml:"name" gorm:"column:name"`
	Surname  string `xml:"surname" gorm:"column:surname"`
	Sex      string `xml:"sex" gorm:"column:sex"`
	Login    string `xml:"login" gorm:"column:login"`
	Password string `xml:"pass" gorm:"column:password"`
	Address  string `xml:"address" gorm:"column:address"`
	Phone    string `xml:"phone" gorm:"column:phone"`
	Team     string `xml:"team" gorm:"column:team"`
	Role     string `xml:"role" gorm:"column:role"`
	Remove   bool   `xml:"remove" gorm:"column:remove"`
}

func (User) TableName() string {
	return "tb_users"
}