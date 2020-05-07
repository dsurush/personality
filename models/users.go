package models

import (
	"MF/db"
	"MF/hamsoyamodels"
	"github.com/sirupsen/logrus"
)

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
	Remove   bool   `xml:"remove" gorm:"column:remove; default:false"`
}


func (*User) TableName() string {
	return "tb_users"
}

func FindUserByLogin(login string) (user User, err error){
	if err := db.GetPostgresDb().Where("login = ?", login).First(&user).Error; err != nil{
		logrus.Warn("Find User By LoginHandler ", err.Error())
		return user, err
	}
//	fmt.Println("I AM = ", user)
	return user, nil
}


type Usersvc struct {
}

func NewUsersvc() *Usersvc {
	return &Usersvc{}
}


func (receiver *Usersvc) GetClientsInfo() (clientsInfo []ClientInfo) {
	if err := db.GetPostgresDb().Limit(100).Offset(0).Find(&clientsInfo).Error; err != nil{
		logrus.Warn("GetClientsInfo:", err.Error())
		return nil
	}
	return clientsInfo
}

func (receiver *Usersvc) GetHamsoyaTransactionsType(size, page int64) (HamsoyaTransactionsType []hamsoyamodels.HamsoyaTransactionType, err error){
	postgresDb := db.GetHamsoyaPostgresDb()
	if err := postgresDb.Limit(10).Offset(0).Find(&HamsoyaTransactionsType).Error; err !=  nil{
		return HamsoyaTransactionsType, err
	}
	return
}

func GetHamsoyaTransactionsType(size, page int64) (HamsoyaTransactionsType []hamsoyamodels.HamsoyaTransactionType, err error){
	postgresDb := db.GetHamsoyaPostgresDb()
	if err := postgresDb.Limit(10).Offset(0).Find(&HamsoyaTransactionsType).Error; err !=  nil{
		return HamsoyaTransactionsType, err
	}
	return
}
