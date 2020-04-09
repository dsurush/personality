package main

import (
	"MF/cmd/app"
	_ "MF/db"
	"MF/token"
	"github.com/julienschmidt/httprouter"
	"MF/models"
)

func main() {
	router := httprouter.New()
	tokenSvc := token.NewTokenSvc([]byte(`surush`))
	//db := models.GetPostgresDb()
	usersvc := models.NewUsersvc()
	server := app.NewMainServer(router, tokenSvc, usersvc)
	server.Start()

}