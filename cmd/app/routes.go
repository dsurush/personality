package app

import (
	"MF/hamsoyamodels"
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

	server.router.GET(`/api/megafon/logs`, logger.Logger(`Get Megafon logs `)(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))((authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetViewLogsHandler)))))
	server.router.GET(`/api/megafon/logs/log/:id`, logger.Logger(`Get by id Megafing log`)(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))((authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetViewLogHandler)))))
	//This router for not full logs form (DataTransferObject)
	server.router.GET(`/api/megafon/logs/DTO`, logger.Logger(`Change Megafon logs DTO `)(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))((authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetViewLogsDTOHandler)))))
	///Hamsoya

	server.router.GET(`/api/hamsoya/transactionstype`, logger.Logger(`Get Hamsoya TransactionTypeTypes`)(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))((authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetHamsoyaTransactionTypeHandler)))))
	server.router.GET(`/api/hamsoya/transactionstype/transactiontype/:id`, logger.Logger(`Get Hamsoya TransactionType by id `)(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))((authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetHamsoyaTransactionTypeByIdHandler)))))
	server.router.POST(`/api/hamsoya/transactionstype/save`,  logger.Logger(`save new Hamsoya TransactionType`)(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))((authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.SaveHamsoyaTransactionType)))))
	server.router.POST(`/api/hamsoya/transactionstype/transactiontype/:id/edit`, logger.Logger(`Edit Hamsoya TransactionType`)(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))((authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.UpdateHamsoyaTransactionTypeHandler)))))

	server.router.GET(`/api/hamsoya/transactions`, logger.Logger(`Get Hamsoya Transactions`)(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))((authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetHamsoyaTransactionsHandler)))))
	server.router.GET(`/api/hamsoya/transactions/transaction/:id`, logger.Logger(`Get Hamsoya Transaction by id`)(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))((authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetHamsoyaTransactionByIdHandler)))))

	server.router.GET(`/api/hamsoya/configs`, logger.Logger(`Get Hamsoya configs`)(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))((authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetHamosyaConfigsHandler)))))
	server.router.GET(`/api/hamsoya/configs/config/:id`, logger.Logger(`Get Hamsoya config by id`)(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))((authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetHamsoyaConfigByIdHandler)))))
	server.router.POST(`/api/hamsoya/configs/save`, logger.Logger(`Save Hamsoya config`)(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))((authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.SaveHamsoyaConfigHandler)))))
	server.router.POST(`/api/hamsoya/configs/config/:id/edit`, logger.Logger(`Edit Hamsoya config`)(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))((authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.UpdateHamsoyaConfigHandler)))))

	server.router.GET(`/api/hamsoya/acoounttypes`, logger.Logger(`Get Hamsoya accounttypes`)(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))((authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetHamosyaAccountTypesHandler)))))
	server.router.GET(`/api/hamsoya/acoounttypes/accounttype/:id`, logger.Logger(`Get Hamsoya accounttype by id`)(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))((authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetHamsoyaAccountTypeByIdHandler)))))
	server.router.POST(`/api/hamsoya/acoounttypes/accounttype/:id/edit`, logger.Logger(`Edit Hamsoya accounttype`)(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))((authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.UpdateHamsoyaAccountTypeHandler)))))
	server.router.POST(`/api/hamsoya/acoounttypes/save`, logger.Logger(`Save Hamsoya accounttype`)(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))((authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.SaveHamsoyaAccountTypeHandler)))))

	server.router.GET(`/ap/hamsoya/statuses`, logger.Logger(`Get Hamsoya statuses`)(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))((authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetHamsoyaStatusesHandler)))))
	server.router.GET(`/ap/hamsoya/statuses/status/:id`, logger.Logger(`Get Hamsoya status by id `)(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))((authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetHamsoyaStatusHandler)))))
	server.router.POST(`/ap/hamsoya/statuses/status/:id/edit`, logger.Logger(`Edit Hamsoya status`)(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))((authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.UpdateHamsoyaStatusHandler)))))
	server.router.POST(`/ap/hamsoya/statuses/save`, logger.Logger(`Save Hamsoya status`)(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))((authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.SaveHamsoyaStatusHandler)))))

	//TODO: filter by time
	server.router.GET(`/api/hamsoya/viewtransactions`, logger.Logger(`Get Hamsoya viewtransactions`)(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))((authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetHamsoyaViewTransactionsHandler)))))
	server.router.GET(`/api/hamsoya/viewtransactions/transaction/:id`, logger.Logger(`Get Hamsoya viewtransaction by id`)(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))((authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetHamsoyaViewTransactionHandler)))))

	//server.router.GET(`/api/hamsoya/viewtranses`, server)
	server.router.GET(`/api/hamsoya/viewtranses/trans/:id`, logger.Logger(`Get Hamsoya viewtrans by id `)(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))((authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetHamsoyaViewTransHandler)))))

	server.router.GET(`/api/hamsoya/documents`, logger.Logger(`Get Hamsoya Documents`)(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))((authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetHamsoyaDocumentsHandler)))))
	server.router.GET(`/api/hamsoya/documents/document/:id`, logger.Logger(`Get Hamsoya Document by id`)(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))((authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetHamsoyaDocumentByIdHandler)))))

	server.router.GET(`/api/hamsoya/records`, logger.Logger(`Get Hamsoya records`)(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))((authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetHamsoyaRecordsHandler)))))
	server.router.GET(`/api/hamsoya/records/record/:id`, logger.Logger(`Get Hamsoya record by id`)(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))((authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetHamsoyaRecordByIdHandler)))))

	server.router.GET(`/api/hamsoya/prechecks`, logger.Logger(`Get Hamsoya prechecks`)(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))((authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetHamsoyaPrechecksHandler)))))
	server.router.GET(`/api/hamsoya/prechecks/precheck/:id`, logger.Logger(`Get Hamsoya precheck by id`)(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))((authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetHamsoyaPrecheckByIdHandler)))))

	server.router.GET(`/api/hamsoya/accounts`, server.GetHamsoyaAccountsHandler)
	server.router.GET(`/api/hamsoya/accounts/account/:id`, logger.Logger(`Get Hamsoya account by id`)(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))((authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetHamsoyaAccountByIdHandler)))))

	panic(http.ListenAndServe("127.0.0.1:8080", server))
}

func test(){
//	var a hamsoyamodels.HamsoyaClientType
//	unix := time.Now().Unix()
//	fmt.Println(unix)
//	i := time.Unix(unix, 0)
//	fmt.Println(i.Format(time.RFC3339))
//	myDataString := "2020-02-14T11:54:14.186066+00:00"
//	myDate, err := time.Parse(time.RFC3339, myDataString)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(myDate.Unix())
//	HamsoyaTransactions := hamsoyamodels.GetTEST(1581681254, 1, 20)
//	fmt.Println(HamsoyaTransactions)
	HamsoyaPreCheck, err := hamsoyamodels.GetHamsoyaPreCheckById(8)
	fmt.Println(HamsoyaPreCheck, err);
	HamsoyaRecord, err := hamsoyamodels.GetHamsoyaRecordById(8)
	fmt.Println(HamsoyaRecord, err);
}
