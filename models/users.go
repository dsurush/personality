package models

import (
	"MF/db"
	"MF/helperfunc"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
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

func FindUserByLogin(login string) (user User, err error) {
	if err := db.GetPostgresDb().Where("login = ?", login).First(&user).Error; err != nil {
		logrus.Warn("Find User By LoginHandler ", err.Error())
		return user, err
	}
	//	fmt.Println("I AM = ", user)
	return user, nil
}

func FindUserByID(id string) (user User, err error) {
	if err := db.GetPostgresDb().Where("id = ?", id).First(&user).Error; err != nil {
		logrus.Warn("Find User By id ", err.Error())
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

func (receiver *Usersvc) GetUsers() (Users []User, err error) {
	if err := db.GetPostgresDb().Limit(100).Offset(0).Find(&Users).Error; err != nil {
		logrus.Warn("GetUsers:", err.Error())
		return nil, err
	}
	return Users, nil
}

func (receiver *Usersvc) ChangeUserPass(id string, pass string, newPass string) (err error) {
	//Находим пользователя по ID
	user, err := FindUserByID(id)
	if err != nil {
	//	err = token.ErrInvalidPasswordOrLogin
		return
	}
	//Проверяем сопадают ли пароли
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass))
	if err != nil {
	//	err = token.ErrInvalidPasswordOrLogin
		return
	}
	//Обновляем пароль
	postgresDb := db.GetPostgresDb()
	HashNewPass, err := helperfunc.HashPassword(newPass)
	if err != nil {
		//err = token.ErrInvalidPasswordOrLogin
		return
	}
	user.Password = HashNewPass
	err = postgresDb.Model(&user).Update(user).Error
	if err != nil {
	//	err = token.ErrInvalidPasswordOrLogin
		return
	}
	return
}