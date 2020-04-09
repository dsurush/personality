package app

import (
	"acuser/middleware/logger"
	"fmt"
	"net/http"
)

func (server *MainServer) InitRoutes() {
	//router := httprouter.New()
	//bytes, _ := bcrypt.GenerateFromPassword([]byte(`surush`), bcrypt.DefaultCost)
	//fmt.Println(string(bytes))
		//user, err := models.FindUserByLogin("surush")
		//if err != nil {
		//	log.Fatalf("НЕТ ТАКОГО ПОЛЬЗОВАТЕЛЯ %e\n", err)
		//}
		//fmt.Println(user)
	fmt.Println("Init routes")
	server.router.POST("/api/login", logger.Logger(`Create Token for user:: "`)(server.CreateTokenHandler))
	server.router.GET("/api/client/:phone", server.GetClientInfoHandler)
	server.router.GET("/api/clients", server.GetClientsInfoHandler)
	//TODO: GET ALL TRANSACTION
	//TODO:
	panic(http.ListenAndServe("127.0.0.1:8080", server))
}