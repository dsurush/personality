package app

import (
	"MF/models"
	"MF/token"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type MainServer struct {
	router *httprouter.Router
	tokenSvc *token.TokenSvc
	userSvc *models.Usersvc
}

func NewMainServer(router *httprouter.Router, tokenSvc *token.TokenSvc, userSvc *models.Usersvc) *MainServer {
	return &MainServer{router: router, tokenSvc: tokenSvc, userSvc: userSvc}
}


func (server *MainServer) Start() {
//	clientInfo := models.GetClientInfo(`2326889`)
//fmt.Println(clientInfo)
	server.InitRoutes()
}

func (server *MainServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	// delegation////
	server.router.ServeHTTP(writer, request)
}