package app

import (
	"MF/middleware/authorized"
	"MF/middleware/jwt"
	"MF/middleware/logger"
	"MF/models"
	"time"

	//	"MF/models"
	"MF/token"
	"fmt"
	"net/http"
	"reflect"
)

func (server *MainServer) InitRoutes() {
	fmt.Println("Init routes")
	test()

	server.router.POST("/api/login", logger.Logger(`Create Token for user:`)(server.LoginHandler))
	server.router.GET("/api/client/:phone", logger.Logger(`Get client: `)(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))((authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetClientInfoByPhoneNumberHandler)))))
	server.router.GET("/api/clients", logger.Logger(`Get all clients By Page and Rows`)(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))((authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetClientsInfoByFilterHandler)))))
	server.router.GET("/api/vendors", logger.Logger(`Get vendors `)(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))((authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetVendorCategoryByPageSizeHandler)))))
	server.router.GET(`/api/transactions`, )

	panic(http.ListenAndServe("127.0.0.1:8080", server))
}

func test(){
	clientsSlice := models.GetClients(models.ClientInfo{
		ClientID:            0,
		Phone:               "",
		Name:                "",
		BirthDate:           time.Time{},
		INN:                 "",
		PassportSeries:      "",
		PassportNumber:      "",
		PassportIssuingAuth: "",
		PassportIssuingDate: time.Time{},
		Address:             "",
		Nationality:         "",
		Sex:                 "W",
		IsActive:            false,
		IsIdentified:        false,
		IsBlackList:         false,
		SendToCft:           false,
	}, 10, 0)
	fmt.Println(clientsSlice)
}