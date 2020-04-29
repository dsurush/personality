package app

import (
	"MF/middleware/authorized"
	"MF/middleware/jwt"
	"MF/middleware/logger"
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
	server.router.POST(`/api/vendors/save`, server.SaveNewVendorHandler)
	server.router.GET(`/api/transactions`, logger.Logger(`Get transactions `)(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))((authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetViewTransactionsHandler)))))
	server.router.GET(`/api/reports`, logger.Logger(`Get reports `)(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))((authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetViewReportsHandler)))))
	panic(http.ListenAndServe("127.0.0.1:8080", server))
}

func test(){
	//vendor := models.Vendor{
	//	ID:           1002,
	//	LatName:      "surush",
	//	Name:         "surush",
	//	CatID:        1,
	//	Feept:        1,
	//	Prefix:       "surush",
	//	HumoPayID:    1,
	//	TajPayID:     1,
	//	ExpressPayID: "surush",
	//	AmonatBonkID: 1,
	//	HumoPayNewID: 1,
	//	Route:        0,
	//	MinSum:       0,
	//	CreateTime:   time.Now(),
	////	Type:         "",
	//	IsActive:     false,
	//}
	//save := vendor.Save()
	//fmt.Println(save)
}
