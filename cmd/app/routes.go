package app

import (
	"MF/middleware/authorized"
	"MF/middleware/jwt"
	"MF/middleware/logger"
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
	server.router.GET("/api/megafon/client/:phone", logger.Logger(`Get client: `)(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))((authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetClientInfoByPhoneNumberHandler)))))
	server.router.GET("/api/megafon/clients", logger.Logger(`Get all clients By Page and Rows`)(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))((authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetClientsInfoHandler)))))
	server.router.GET("/api/megafon/vendors", logger.Logger(`Get vendors `)(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))((authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetVendorCategoryByPageSizeHandler)))))
	server.router.POST(`/api/megafon/vendors/save`, logger.Logger(`Save vendor `)(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))((authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.SaveNewVendorHandler)))))
	server.router.GET(`/api/megafon/vendors/vendor/:id`, logger.Logger(`Get vendor by id `)(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))((authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetVendorHandler)))))
	server.router.POST(`/api/megafon/vendors/vendor/:id/edit`, logger.Logger(`Change vendor by id `)(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))((authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.UpdateVendorHandler)))))
	server.router.GET(`/api/megafon/transactions`, logger.Logger(`Get transactions `)(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))((authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetViewTransactionsHandler)))))
	server.router.GET(`/api/megafon/reports`, logger.Logger(`Get reports `)(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))((authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetViewReportsHandler)))))
	server.router.GET(`/api/megafon/merchants`, logger.Logger(`Get merchants `)(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))((authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetMerchantsHandler)))))
	server.router.GET(`/api/megafon/merchants/merchant/:id`, logger.Logger(`Get merchant `)(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))((authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetMerchantHandler)))))
	server.router.POST(`/api/megafon/merchants/merchant/:id/edit`, logger.Logger(`Change merchant by id `)(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))((authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.UpdateMerchantHandler)))))
	server.router.GET(`/api/megafon/logs`, server.GetViewLogsHandler)
	server.router.GET(`/api/megafon/logs/log/:id`, server.GetViewLogHandler)
	//This router for not full logs form (DataTransferObject)
	server.router.GET(`/api/megafon/logs/DTO`, server.GetViewLogsDTOHandler)
	///Hamsoya
	server.router.GET(`/api/hamsoya/transactionstype`, server.GetHamsoyaTransactionTypeHandler)
	server.router.GET(`/api/hamsoya/transactionstype/transactiontype/:id`, server.GetHamsoyaTransactionTypeByIdHandler)
	server.router.POST(`/api/hamsoya/transactionstype/save`,  server.SaveHamsoyaTransactionType)
	server.router.POST(`/api/hamsoya/transactionstype/transactiontype/:id/edit`, server.UpdateHamsoyaTransactionTypeHandler)

	server.router.GET(`/api/hamsoya/transactions`, server.GetHamsoyaTransactionsHandler)
	server.router.GET(`/api/hamsoya/transactions/transaction/:id`, server.GetHamsoyaTransactionByIdHandler)

	server.router.GET(`/api/hamsoya/configs`, server.GetHamosyaConfigsHandler)
	server.router.GET(`/api/hamsoya/configs/config/:id`, server.GetHamsoyaConfigByIdHandler)
	server.router.POST(`/api/hamsoya/configs/save`, server.SaveHamsoyaConfigHandler)
	server.router.POST(`/api/hamsoya/configs/config/:id/edit`, server.UpdateHamsoyaConfigHandler)

	server.router.GET(`/api/hamsoya/acoounttypes`, server.GetHamosyaAccountTypesHandler)
	server.router.GET(`/api/hamsoya/acoounttypes/accounttype/:id`, server.GetHamsoyaAccountTypeByIdHandler)
	server.router.POST(`/api/hamsoya/acoounttypes/accounttype/:id/edit`, server.UpdateHamsoyaAccountTypeHandler)
	server.router.POST(`/api/hamsoya/acoounttypes/save`, server.SaveHamsoyaAccountTypeHandler)

	panic(http.ListenAndServe("127.0.0.1:8080", server))
}

func test(){

	myDateString := "2019-10-30T01:07:39.085082+05:00"
	fmt.Println("My Starting Date:\t", myDateString)
	myDate, err := time.Parse( time.RFC3339, myDateString)
	if err != nil {
		panic(err)
	}
	fmt.Println("My Date Reformatted:\t", myDate)
}
