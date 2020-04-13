package token

import (
	"MF/models"
	"context"
	"errors"
	"github.com/dsurush/jwt/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type TokenSvc struct {
	secret []byte
}

func NewTokenSvc(secret []byte) *TokenSvc {
	return &TokenSvc{secret: secret}
}


type Payload struct {
	Id    int64    `json:"id"`
	Exp   int64    `json:"exp"`
	Role string    `json:"role"`
}

type RequestDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ResponseDTO struct {
	Token string `json:"token"`
}

//var ErrInvalidLogin = errors.New("invalid login or password")
var ErrInvalidPasswordOrLogin = errors.New("invalid password")


func (receiver *TokenSvc) Generate(context context.Context, request *RequestDTO) (response ResponseDTO, err error) {
	user, err := models.FindUserByLogin(request.Username)
	if err != nil {
		err = ErrInvalidPasswordOrLogin
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		err = ErrInvalidPasswordOrLogin
		return
	}

	response.Token, err = jwt.Encode(Payload{
		Id:    user.Id,
		Exp:   time.Now().Add(time.Hour).Unix(),
		Role:  user.Role,
	}, receiver.secret)
	return
}
