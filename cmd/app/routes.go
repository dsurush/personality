package app

import (
	"MF/middleware/authorized"
	"MF/middleware/corss"
	"MF/settings"
	"net/http"
	//"MF/middleware/corss"
	"MF/middleware/jwt"
	"MF/middleware/logger"
	"MF/token"
	"fmt"
	"reflect"
)

func (server *MainServer) InitRoutes() {
	fmt.Println("Init routes")
	test(server)

	server.router.POST("/api/login", logger.Logger(`Create Token for user:`)(corss.Middleware(server.LoginHandler)))

	//Static
	server.router.GET(`/api/megafon`, logger.Logger(`Get static`)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`, `user`}, jwt.FromContext)(server.GetMegafonStaticHandler)))))

	server.router.GET("/api/megafon/clients/:id", logger.Logger(`Get client by id: `)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`, `user`}, jwt.FromContext)(server.GetClientInfoByIdHandler)))))
	//server.router.GET("/api/megafon/clients", logger.Logger(`Get all clients By Page and Rows`)(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))((authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetClientsInfoHandler)))))
	server.router.GET("/api/megafon/clients", logger.Logger(`Get all clients By Page and Rows`)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`, `user`}, jwt.FromContext)(server.GetClientsInfoHandler)))))

	server.router.GET("/api/megafon/vendors", logger.Logger(`Get vendors `)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`, `user`}, jwt.FromContext)(server.GetVendorsHandler)))))
	server.router.GET(`/api/megafon/vendors/:id`, logger.Logger(`Get vendor by id `)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`, `user`}, jwt.FromContext)(server.GetVendorHandler)))))
	server.router.POST(`/api/megafon/vendors/:id/edit`, logger.Logger(`Change vendor by id `)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.UpdateVendorHandler)))))
	server.router.GET(`/api/megafon/vendors/:id/edit`, logger.Logger(`Get vendor for edit `)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetVendorHandler)))))
	server.router.POST(`/api/megafon/vendors/:id/add`, logger.Logger(`Save vendor `)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.SaveNewVendorHandler)))))

	server.router.GET(`/api/megafon/view-transactions`, logger.Logger(`Get view - transactions `)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`, `user`}, jwt.FromContext)(server.GetViewTransactionsHandler)))))
	server.router.GET(`/api/megafon/view-transactions/:id`, logger.Logger(`Get view - transactions by id `)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`, `user`}, jwt.FromContext)(server.GetViewTransactionByIdHandler)))))

	server.router.GET(`/api/megafon/transactions`, logger.Logger(`Get transactions `)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`, `user`}, jwt.FromContext)(server.GetTableTransactionsHandler)))))
	server.router.GET(`/api/megafon/transactions/:id`, logger.Logger(`Get transactions by id `)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`, `user`}, jwt.FromContext)(server.GetTableTransactionByIdHandler)))))
	server.router.GET(`/api/megafon/transactions/:id/edit`, logger.Logger(`Get transactions by id for edit`)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetTableTransactionByIdStatusHandler)))))
	server.router.POST(`/api/megafon/transactions/:id/edit`, logger.Logger(`Set transactions by id for edit`)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.SetNewTransactionStatus)))))
	server.router.POST(`/api/megafon/transactions/:id/cancel`, logger.Logger(`Cancel Transaction`)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.CancelMegafonTransactionHandler)))))


	server.router.GET(`/api/megafon/reports`, logger.Logger(`Get reports `)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`, `user`}, jwt.FromContext)(server.GetViewReportsHandler)))))
	server.router.GET(`/api/megafon/reports/:id`, logger.Logger(`Get reports by id `)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`, `user`}, jwt.FromContext)(server.GetViewReportsByIdHandler)))))
	server.router.GET(`/api/megafon/downloads-reports`, logger.Logger(`Get reports to excel file`)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetViewReportsForExcelHandler)))))

	server.router.GET(`/api/megafon/merchants`, logger.Logger(`Get merchants `)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`, `user`}, jwt.FromContext)(server.GetMerchantsHandler)))))
	server.router.GET(`/api/megafon/merchants/:id`, logger.Logger(`Get merchant `)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`, `user`}, jwt.FromContext)(server.GetMerchantHandler)))))
	server.router.POST(`/api/megafon/merchants/:id/edit`, logger.Logger(`Change merchant by id `)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.UpdateMerchantHandler)))))
	server.router.GET(`/api/megafon/merchants/:id/edit`, logger.Logger(`Get merchant for edit `)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetMerchantHandler)))))

	//server.router.GET(`/api/megafon/logs`, logger.Logger(`Get Megafon logs `)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetViewLogsHandler)))))
	server.router.GET(`/api/megafon/logs/:id`, logger.Logger(`Get by id MegafonFin log`)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`, `user`}, jwt.FromContext)(server.GetViewLogHandler)))))
	//This router for not full logs form (DataTransferObject)
	server.router.GET(`/api/megafon/logs`, logger.Logger(`Change Megafon logs DTO `)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`, `user`}, jwt.FromContext)(server.GetViewLogsDTOHandler)))))

	///Hamsoya
	server.router.GET(`/api/hamsoya/transactionstype`, logger.Logger(`Get Hamsoya TransactionTypeTypes`)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`, `user`}, jwt.FromContext)(server.GetHamsoyaTransactionTypeHandler)))))
	server.router.GET(`/api/hamsoya/transactionstype/:id`, logger.Logger(`Get Hamsoya TransactionType by id `)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`, `user`}, jwt.FromContext)(server.GetHamsoyaTransactionTypeByIdHandler)))))
	server.router.GET(`/api/hamsoya/transactionstype/:id/edit`, logger.Logger(`Get Hamsoya TransactionType by id `)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetHamsoyaTransactionTypeByIdHandler)))))
	server.router.POST(`/api/hamsoya/transactionstype/:id/add`, logger.Logger(`save new Hamsoya TransactionType`)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.SaveHamsoyaTransactionType)))))
	server.router.POST(`/api/hamsoya/transactionstype/:id/edit`, logger.Logger(`Edit Hamsoya TransactionType`)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.UpdateHamsoyaTransactionTypeHandler)))))

	server.router.GET(`/api/hamsoya/transactions`, logger.Logger(`Get Hamsoya Transactions`)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`, `user`}, jwt.FromContext)(server.GetHamsoyaTransactionsHandler)))))
	server.router.GET(`/api/hamsoya/transactions/:id`, logger.Logger(`Get Hamsoya Transaction by id`)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`, `user`}, jwt.FromContext)(server.GetHamsoyaTransactionByIdHandler)))))

	server.router.GET(`/api/hamsoya/configs`, logger.Logger(`Get Hamsoya configs`)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`, `user`}, jwt.FromContext)(server.GetHamosyaConfigsHandler)))))
	server.router.GET(`/api/hamsoya/configs/:id`, logger.Logger(`Get Hamsoya config by id`)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`, `user`}, jwt.FromContext)(server.GetHamsoyaConfigByIdHandler)))))
	server.router.GET(`/api/hamsoya/configs/:id/edit`, logger.Logger(`Get Hamsoya config by id`)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetHamsoyaConfigByIdHandler)))))
	server.router.POST(`/api/hamsoya/configs/:id/add`, logger.Logger(`Save Hamsoya config`)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.SaveHamsoyaConfigHandler)))))
	server.router.POST(`/api/hamsoya/configs/:id/edit`, logger.Logger(`Edit Hamsoya config`)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.UpdateHamsoyaConfigHandler)))))

	server.router.GET(`/api/hamsoya/accounttypes`, logger.Logger(`Get Hamsoya accounttypes`)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`, `user`}, jwt.FromContext)(server.GetHamosyaAccountTypesHandler)))))
	server.router.GET(`/api/hamsoya/accounttypes/:id`, logger.Logger(`Get Hamsoya accounttype by id`)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`, `user`}, jwt.FromContext)(server.GetHamsoyaAccountTypeByIdHandler)))))
	server.router.GET(`/api/hamsoya/accounttypes/:id/edit`, logger.Logger(`Get Hamsoya accounttype by id`)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetHamsoyaAccountTypeByIdHandler)))))
	server.router.POST(`/api/hamsoya/accounttypes/:id/edit`, logger.Logger(`Edit Hamsoya accounttype`)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.UpdateHamsoyaAccountTypeHandler)))))
	server.router.POST(`/api/hamsoya/accounttypes/:id/add`, logger.Logger(`Save Hamsoya accounttype`)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.SaveHamsoyaAccountTypeHandler)))))

	server.router.GET(`/api/hamsoya/statuses`, logger.Logger(`Get Hamsoya statuses`)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`, `user`}, jwt.FromContext)(server.GetHamsoyaStatusesHandler)))))
	server.router.GET(`/api/hamsoya/statuses/:id`, logger.Logger(`Get Hamsoya status by id `)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`, `user`}, jwt.FromContext)(server.GetHamsoyaStatusHandler)))))
	server.router.GET(`/api/hamsoya/statuses/:id/edit`, logger.Logger(`Get Hamsoya status by id `)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetHamsoyaStatusHandler)))))
	server.router.POST(`/api/hamsoya/statuses/:id/edit`, logger.Logger(`Edit Hamsoya status`)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.UpdateHamsoyaStatusHandler)))))
	server.router.POST(`/api/hamsoya/statuses/:id/add`, logger.Logger(`Save Hamsoya status`)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.SaveHamsoyaStatusHandler)))))

	//
	server.router.GET(`/api/hamsoya/viewtransactions`, logger.Logger(`Get Hamsoya viewtransactions`)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`, `user`}, jwt.FromContext)(server.GetHamsoyaViewTransactionsHandler)))))
	server.router.GET(`/api/hamsoya/viewtransactions/:id`, logger.Logger(`Get Hamsoya viewtransaction by id`)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`, `user`}, jwt.FromContext)(server.GetHamsoyaViewTransactionHandler)))))

	server.router.GET(`/api/hamsoya/viewtranses`, logger.Logger(`Get Hamsoya viewtrans by id `)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`, `user`}, jwt.FromContext)(server.GetHamsoyaViewTransesHandler)))))
	server.router.GET(`/api/hamsoya/viewtranses/:id`, logger.Logger(`Get Hamsoya viewtrans by id `)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`, `user`}, jwt.FromContext)(server.GetHamsoyaViewTransHandler)))))

	server.router.GET(`/api/hamsoya/documents`, logger.Logger(`Get Hamsoya Documents`)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`, `user`}, jwt.FromContext)(server.GetHamsoyaDocumentsHandler)))))
	server.router.GET(`/api/hamsoya/documents/:id`, logger.Logger(`Get Hamsoya Document by id`)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`, `user`}, jwt.FromContext)(server.GetHamsoyaDocumentByIdHandler)))))

	server.router.GET(`/api/hamsoya/records`, logger.Logger(`Get Hamsoya records`)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`, `user`}, jwt.FromContext)(server.GetHamsoyaRecordsHandler)))))
	server.router.GET(`/api/hamsoya/records/:id`, logger.Logger(`Get Hamsoya record by id`)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`, `user`}, jwt.FromContext)(server.GetHamsoyaRecordByIdHandler)))))

	server.router.GET(`/api/hamsoya/prechecks`, logger.Logger(`Get Hamsoya prechecks`)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`, `user`}, jwt.FromContext)(server.GetHamsoyaPrechecksHandler)))))
	server.router.GET(`/api/hamsoya/prechecks/:id`, logger.Logger(`Get Hamsoya precheck by id`)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`, `user`}, jwt.FromContext)(server.GetHamsoyaPrecheckByIdHandler)))))

	server.router.GET(`/api/hamsoya/accounts`, logger.Logger(`Get Hamsoya accounts`)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`, `user`}, jwt.FromContext)(server.GetHamsoyaAccountsHandler)))))
	server.router.GET(`/api/hamsoya/accounts/:id`, logger.Logger(`Get Hamsoya account by id`)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`, `user`}, jwt.FromContext)(server.GetHamsoyaAccountByIdHandler)))))

	server.router.GET(`/api/hamsoya/clients`, logger.Logger(`Get Hamsoya clients`)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`, `user`}, jwt.FromContext)(server.GetHamsoyaClientsHandler)))))

	server.router.GET(`/api/hamsoya/clients/:id`,
		logger.Logger(`Get Hamsoya client by id`)(
		corss.Middleware(
		jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(
		authorized.Authorized([]string{`admin`, `user`}, jwt.FromContext)(
		server.GetHamsoyaClientByIdHandler)))))


	server.router.POST(`/api/me/change-pass`, logger.Logger(`Get Hamsoya client by id`)(
		corss.Middleware(
		jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(
		authorized.Authorized([]string{`admin`, `user`}, jwt.FromContext)(
		server.ChangePasswordHandler)))))

	server.router.GET(`/api/users`, logger.Logger(`Get user list`)(
		corss.Middleware(
		jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(
		authorized.Authorized([]string{`admin`}, jwt.FromContext)(
		server.GetUsersHandler)))))

	server.router.POST(`/api/users/add`, logger.Logger(`add new user`)(
		corss.Middleware(
		jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(
		authorized.Authorized([]string{`admin`}, jwt.FromContext)(
		server.AddUserHandler)))))

//	server.router.POST(``)

	//	server.router.
	//handler := cors.Default().Handler(server)
	//panic(http.ListenAndServe("127.0.0.1:8080", handler))
	port := fmt.Sprintf(":%d", settings.AppSettings.AppParams.PortRun)
	fmt.Printf("Server is listening in port %s%s ...", settings.AppSettings.AppParams.ServerName, port)
	panic(http.ListenAndServe(port, server))
}

func test(server *MainServer) {
	//user := models.User{
	//	Id:       0,
	//	Name:     "tim",
	//	Surname:  "tim",
	//	Sex:      "tim",
	//	Login:    "tim",
	//	Password: "tim",
	//	Address:  "tim",
	//	Phone:    "tim",
	//	Team:     "tim",
	//	Role:     "admin",
	//	Remove:   false,
	//}
	//err := server.userSvc.AddNewUser(user)
	//if err != nil {
	//	fmt.Println("pizda")
	//} else{
	//	fmt.Println("anus")
	//}
}
