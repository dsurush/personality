package app

import (
	"MF/middleware/authorized"
	"MF/middleware/jwt"
	"MF/middleware/logger"
	"MF/models"
	"MF/token"
	"fmt"
	"net/http"
	"reflect"
	"time"
)

func (server *MainServer) InitRoutes() {
	fmt.Println("Init routes")
	test()

	server.router.POST("/api/login", logger.Logger(`Create Token for user:`)(server.LoginHandler))
	server.router.GET("/api/client/:phone", logger.Logger(`Get client: `)(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))((authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetClientInfoByPhoneNumberHandler)))))
	server.router.GET("/api/clients", logger.Logger(`Get all clients By Page and Rows`)(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))((authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetClientsInfoByFilterHandler)))))
	server.router.GET("/api/vendors", logger.Logger(`Get vendors `)(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))((authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetVendorCategoryByPageSizeHandler)))))
	server.router.POST(`/api/vendors/save`, server.SaveNewVendorHandler)
	server.router.GET(`/api/vendors/vendor/:id`, server.GetVendorHandler)
	server.router.POST(`/api/vendors/vendor/:id/edit`, server.UpdateVendorHandler)
	server.router.GET(`/api/transactions`, logger.Logger(`Get transactions `)(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))((authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetViewTransactionsHandler)))))
	server.router.GET(`/api/reports`, logger.Logger(`Get reports `)(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))((authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetViewReportsHandler)))))
	server.router.GET(`/api/merchants`, server.GetMerchantsHandler)
	server.router.GET(`/api/merchants/merchant/:id`, server.GetMerchantHandler)
	server.router.POST(`/api/merchants/merchant/:id/edit`, server.UpdateMerchantHandler)
	panic(http.ListenAndServe("127.0.0.1:8080", server))
}

func test(){
	merchant := models.Merchant{}
	merchants := merchant.GetMerchantById(1060)
	update := merchants.Update(models.Merchant{
		ID:           1060,
		HumoOnlineID: 16085,
		NameENG:      "My Name Is Impire",
		NameRUS:      "ЧДММ \"По Империя\"",
		QrCode:       "eyJ0eXBlIjoiTVVMVElfTUVSQ0hBTlQiLCAiZW1haWwiOiI5MDEyMjg4NTVAdmVyeWZha2UxLnRqIiwgIm1vYmlsZU51bWJlciI6IjAwOTkyOTAxMjI4ODU1In0=",
		QrCodeNew:    "eyJ0eXBlIjoiTVVMVElfTUVSQ0hBTlQiLCAiZW1haWwiOiI5MDEyMjg4NTVAdmVyeWZha2UxLnRqIiwgIm1vYmlsZU51bWJlciI6IjAwOTkyOTAxMjI4ODU1In0=	",
		CreateTime:   time.Time{},
		UpdateTime:   time.Now(),
	})
	fmt.Println(update)
}
