package app

import (
	"MF/hamsoyamodels"
	"MF/middleware/authorized"
	"MF/middleware/corss"
	"net/http"
	"time"

	//"MF/middleware/corss"
	"MF/middleware/jwt"
	"MF/middleware/logger"
	"MF/token"
	"fmt"
	"reflect"
)

func (server *MainServer) InitRoutes() {
	fmt.Println("Init routes")
	test()

	//server.router.POST("/api/login", logger.Logger(`Create Token for user:`)(server.LoginHandler))
	server.router.POST("/api/login", logger.Logger(`Create Token for user:`)(corss.Middleware(server.LoginHandler)))

	//test
	server.router.GET(`/api/surush`, server.LoginHandler1)

	server.router.GET("/api/megafon/client/:id", logger.Logger(`Get client by id: `)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetClientInfoByIdHandler)))))
	//server.router.GET("/api/megafon/clients", logger.Logger(`Get all clients By Page and Rows`)(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))((authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetClientsInfoHandler)))))
	server.router.GET("/api/megafon/clients", logger.Logger(`Get all clients By Page and Rows`)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetClientsInfoHandler)))))
//	server.router.GET("/api/megafon/clients", corss.Middleware(server.GetClientsInfoHandler))

	server.router.GET("/api/megafon/vendors", logger.Logger(`Get vendors `)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetVendorCategoryByPageSizeHandler)))))
	server.router.POST(`/api/megafon/vendors/save`, logger.Logger(`Save vendor `)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.SaveNewVendorHandler)))))
	server.router.GET(`/api/megafon/vendors/vendor/:id`, logger.Logger(`Get vendor by id `)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetVendorHandler)))))
	server.router.POST(`/api/megafon/vendors/vendor/:id/edit`, logger.Logger(`Change vendor by id `)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.UpdateVendorHandler)))))

	server.router.GET(`/api/megafon/transactions`, logger.Logger(`Get transactions `)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetViewTransactionsHandler)))))
	server.router.GET(`/api/megafon/reports`, logger.Logger(`Get reports `)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetViewReportsHandler)))))

	server.router.GET(`/api/megafon/merchants`, logger.Logger(`Get merchants `)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetMerchantsHandler)))))
	server.router.GET(`/api/megafon/merchants/merchant/:id`, logger.Logger(`Get merchant `)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetMerchantHandler)))))
	server.router.POST(`/api/megafon/merchants/merchant/:id/edit`, logger.Logger(`Change merchant by id `)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.UpdateMerchantHandler)))))

	server.router.GET(`/api/megafon/logs`, logger.Logger(`Get Megafon logs `)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetViewLogsHandler)))))
	server.router.GET(`/api/megafon/logs/log/:id`, logger.Logger(`Get by id Megafing log`)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetViewLogHandler)))))
	//This router for not full logs form (DataTransferObject)
	server.router.GET(`/api/megafon/logs/DTO`, logger.Logger(`Change Megafon logs DTO `)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetViewLogsDTOHandler)))))

	///Hamsoya
	server.router.GET(`/api/hamsoya/transactionstype`, logger.Logger(`Get Hamsoya TransactionTypeTypes`)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetHamsoyaTransactionTypeHandler)))))
	server.router.GET(`/api/hamsoya/transactionstype/transactiontype/:id`, logger.Logger(`Get Hamsoya TransactionType by id `)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetHamsoyaTransactionTypeByIdHandler)))))
	server.router.POST(`/api/hamsoya/transactionstype/save`,  logger.Logger(`save new Hamsoya TransactionType`)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.SaveHamsoyaTransactionType)))))
	server.router.POST(`/api/hamsoya/transactionstype/transactiontype/:id/edit`, logger.Logger(`Edit Hamsoya TransactionType`)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.UpdateHamsoyaTransactionTypeHandler)))))

	server.router.GET(`/api/hamsoya/transactions`, logger.Logger(`Get Hamsoya Transactions`)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetHamsoyaTransactionsHandler)))))
	server.router.GET(`/api/hamsoya/transactions/transaction/:id`, logger.Logger(`Get Hamsoya Transaction by id`)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetHamsoyaTransactionByIdHandler)))))

	server.router.GET(`/api/hamsoya/configs`, logger.Logger(`Get Hamsoya configs`)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetHamosyaConfigsHandler)))))
	server.router.GET(`/api/hamsoya/configs/config/:id`, logger.Logger(`Get Hamsoya config by id`)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetHamsoyaConfigByIdHandler)))))
	server.router.POST(`/api/hamsoya/configs/save`, logger.Logger(`Save Hamsoya config`)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.SaveHamsoyaConfigHandler)))))
	server.router.POST(`/api/hamsoya/configs/config/:id/edit`, logger.Logger(`Edit Hamsoya config`)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.UpdateHamsoyaConfigHandler)))))

	server.router.GET(`/api/hamsoya/acoounttypes`, logger.Logger(`Get Hamsoya accounttypes`)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetHamosyaAccountTypesHandler)))))
	server.router.GET(`/api/hamsoya/acoounttypes/accounttype/:id`, logger.Logger(`Get Hamsoya accounttype by id`)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetHamsoyaAccountTypeByIdHandler)))))
	server.router.POST(`/api/hamsoya/acoounttypes/accounttype/:id/edit`, logger.Logger(`Edit Hamsoya accounttype`)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.UpdateHamsoyaAccountTypeHandler)))))
	server.router.POST(`/api/hamsoya/acoounttypes/save`, logger.Logger(`Save Hamsoya accounttype`)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.SaveHamsoyaAccountTypeHandler)))))

	server.router.GET(`/ap/hamsoya/statuses`, logger.Logger(`Get Hamsoya statuses`)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetHamsoyaStatusesHandler)))))
	server.router.GET(`/ap/hamsoya/statuses/status/:id`, logger.Logger(`Get Hamsoya status by id `)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetHamsoyaStatusHandler)))))
	server.router.POST(`/ap/hamsoya/statuses/status/:id/edit`, logger.Logger(`Edit Hamsoya status`)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.UpdateHamsoyaStatusHandler)))))
	server.router.POST(`/ap/hamsoya/statuses/save`, logger.Logger(`Save Hamsoya status`)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.SaveHamsoyaStatusHandler)))))

	//TODO: filter by time
	server.router.GET(`/api/hamsoya/viewtransactions`, logger.Logger(`Get Hamsoya viewtransactions`)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetHamsoyaViewTransactionsHandler)))))
	server.router.GET(`/api/hamsoya/viewtransactions/transaction/:id`, logger.Logger(`Get Hamsoya viewtransaction by id`)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetHamsoyaViewTransactionHandler)))))

	//server.router.GET(`/api/hamsoya/viewtranses`, server)
	server.router.GET(`/api/hamsoya/viewtranses/trans/:id`, logger.Logger(`Get Hamsoya viewtrans by id `)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetHamsoyaViewTransHandler)))))

	server.router.GET(`/api/hamsoya/documents`, logger.Logger(`Get Hamsoya Documents`)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetHamsoyaDocumentsHandler)))))
	server.router.GET(`/api/hamsoya/documents/document/:id`, logger.Logger(`Get Hamsoya Document by id`)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetHamsoyaDocumentByIdHandler)))))

	server.router.GET(`/api/hamsoya/records`, logger.Logger(`Get Hamsoya records`)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetHamsoyaRecordsHandler)))))
	server.router.GET(`/api/hamsoya/records/record/:id`, logger.Logger(`Get Hamsoya record by id`)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetHamsoyaRecordByIdHandler)))))

	server.router.GET(`/api/hamsoya/prechecks`, logger.Logger(`Get Hamsoya prechecks`)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetHamsoyaPrechecksHandler)))))
	server.router.GET(`/api/hamsoya/prechecks/precheck/:id`, logger.Logger(`Get Hamsoya precheck by id`)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetHamsoyaPrecheckByIdHandler)))))

	server.router.GET(`/api/hamsoya/accounts`, logger.Logger(`Get Hamsoya accounts`)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetHamsoyaAccountsHandler)))))
	server.router.GET(`/api/hamsoya/accounts/account/:id`, logger.Logger(`Get Hamsoya account by id`)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetHamsoyaAccountByIdHandler)))))

	server.router.GET(`/api/hamsoya/clients`, logger.Logger(`Get Hamsoya clients`)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetHamsoyaClientsHandler)))))
	server.router.GET(`/api/hamsoya/clients/client/:id`, logger.Logger(`Get Hamsoya client by id`)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetHamsoyaClientByIdHandler)))))
	//handler := cors.Default().Handler(server)
	//panic(http.ListenAndServe("127.0.0.1:8080", handler))

	// TODO: DELETE ME
	server.router.GET(`/api/test`, server.TESTGetHamsoyaAccountsHandler)

	panic(http.ListenAndServe("127.0.0.1:8080", server))
}

func test(){
//	var a hamsoyamodels.HamsoyaClientType
//	myDataString := 1591635497
//	unix := time.Now().Unix()
//	fmt.Println(unix)
//	i := time.Unix(unix, 0)
//	fmt.Println(i.Format(time.RFC3339))

	//myDate, err := time.Parse(time.RFC3339, myDataString)
	//if err != nil {
	//	panic(err)
	//}
	myDataString := 1591635497
	i := time.Unix(int64(myDataString), 0)
	ans := i.Format(time.RFC3339)
	HamsoyaAccount := hamsoyamodels.TESTTIME(ans)
	fmt.Println(HamsoyaAccount)
	fmt.Println(time.Now().Unix())
}
