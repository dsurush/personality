package app

import (
	"MF/models"
	"MF/token"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

//MainServer. Структура которая реализует метод ServerHTTP. В его параметры входят все зависимости.
type MainServer struct {
	router   *httprouter.Router
	tokenSvc *token.TokenSvc
	userSvc  *models.Usersvc
}

func NewMainServer(router *httprouter.Router, tokenSvc *token.TokenSvc, userSvc *models.Usersvc) *MainServer {
	return &MainServer{router: router, tokenSvc: tokenSvc, userSvc: userSvc}
}

// In start all starting things
func (server *MainServer) Start() {
	server.InitRoutes()
}

func (server *MainServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	// delegation
	server.router.ServeHTTP(writer, request)
}
