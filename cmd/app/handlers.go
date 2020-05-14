package app

import (
	"MF/hamsoyamodels"
	"MF/models"
	"MF/token"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)
//Handler for login // Get log and pass
func (server *MainServer) LoginHandler(writer http.ResponseWriter, request *http.Request, pr httprouter.Params) {
//	fmt.Println("login\n")
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	var requestBody token.RequestDTO
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(writer).Encode([]string{"err.json_invalid"})
		log.Print(err)
		return
	}
	//	log.Printf("login = %s, pass = %s\n", requestBody.Username, requestBody.Password)
	response, err := server.tokenSvc.Generate(request.Context(), &requestBody)
	//log.Println(response)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(writer).Encode([]string{"err.password_mismatch", err.Error()})
		if err != nil {
			log.Print(err)
		}
		return
	}
	err = json.NewEncoder(writer).Encode(&response)
	if err != nil {
		log.Print(err)
	}
}
//Get clients By Phone
func (server *MainServer) GetClientInfoByPhoneNumberHandler(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	fmt.Println("I am find client By number phone")
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	phone := param.ByName(`phone`)
	fmt.Println(phone)
	response := models.GetClientInfoByPhoneNumber(phone)
	if response.ClientID == 0{
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Println(response)
	err := json.NewEncoder(writer).Encode(&response)
	if err != nil {
		log.Print(err)
	}
}
//Get list clients Handler
func (server *MainServer) GetClientsInfoHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	fmt.Println("I am get all clients")
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	var clientDefault models.ClientInfo
	pageInt := 1
	rowsInt := 100
	page := request.URL.Query().Get(`page`)
	rows := request.URL.Query().Get(`rows`)
	IsActive, err := strconv.ParseBool(request.URL.Query().Get(`IsActive`))
	if err == nil {
		clientDefault.IsActive = IsActive
	}
	IsIdentified, err := strconv.ParseBool(request.URL.Query().Get(`IsIdentified`))
	if err == nil {
		clientDefault.IsIdentified = IsIdentified
	}
	IsBlackList, err := strconv.ParseBool(request.URL.Query().Get(`IsBlackList`))
	if err == nil {
		clientDefault.IsBlackList = IsBlackList
	}
	SendToCft, err := strconv.ParseBool(request.URL.Query().Get(`SendToCft`))
	if err == nil {
		clientDefault.SendToCft = SendToCft
	}
	Sex := request.URL.Query().Get(`Sex`)
	clientDefault.Sex = Sex

	pageInt, err = strconv.Atoi(page)
	if err != nil{
		pageInt = 1
	}
	rowsInt, err = strconv.Atoi(rows)
	if err != nil {
		rowsInt = 100
	}

	fmt.Println(clientDefault)
	response := models.GetClients(clientDefault, rowsInt, pageInt - 1)
	err = json.NewEncoder(writer).Encode(&response)
	if err != nil {
		log.Print(err)
	}
}
//UnUse Handler
func (server *MainServer) LoginHandler1(writer http.ResponseWriter, _*http.Request, _ httprouter.Params) {
	bytes, err := ioutil.ReadFile("web/templates/index.gohtml")
	if err != nil {
		log.Fatal("can't read from /web/templates.index.gohtml\nerr: ", err)
	}
	_, err = writer.Write(bytes)
	if err != nil {
		log.Fatal("can't send bytes to writer")
	}
}
//UnUse Handler
func (server *MainServer) MainPageHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
//	fmt.Println("Login\n")
	var requestBody token.RequestDTO
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(writer).Encode([]string{"err.json_invalid"})
		log.Print(err)
		return
	}
	//		log.Printf("login = %s, pass = %s\n", requestBody.Username, requestBody.Password)
	response, err := server.tokenSvc.Generate(request.Context(), &requestBody)
	//log.Println(response)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(writer).Encode([]string{"err.password_mismatch", err.Error()})
		if err != nil {
			log.Print(err)
		}
		return
	}
	//cookie := http.Cookie{
	//	//	Name:     "token",
	//	//	Value:    response.Token,
	//	//	Expires:  time.Now().Add(time.Hour),
	//	//	HttpOnly: true,
	//	//	Path:     "/",
	//	//	// Domain:   "localhost",
	//	//}
	//http.SetCookie(writer, &cookie)
	err = json.NewEncoder(writer).Encode(&response)
	if err != nil {
		log.Print(err)
	}
}
//UnUse Handler
func (server *MainServer) GetVendorCategoryHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	var vendor models.Vendor
	vendors := vendor.FindAll()
	//fmt.Println("Hello I am vendors\n", vendors)
	err := json.NewEncoder(writer).Encode(&vendors)
	if err != nil {
		log.Print(err)
	}
}
//GetVendorCategory
func (server *MainServer) GetVendorCategoryByPageSizeHandler(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	page := request.URL.Query().Get(`page`)
	rows := request.URL.Query().Get(`rows`)
	pageInt, _ := strconv.Atoi(page)
	rowsInt, _ := strconv.Atoi(rows)
	var vendor models.Vendor
	response := vendor.FindAllVendorsByPageSize(pageInt-1, rowsInt)
	err := json.NewEncoder(writer).Encode(&response)
	if err != nil {
		log.Print(err)
	}
}
// Get view Trans for report
func (server *MainServer) GetViewTransactionsHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	fmt.Println("I am view Transaction")
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	pageInt := 1
	rowsInt := 100
	page := request.URL.Query().Get(`page`)
	rows := request.URL.Query().Get(`rows`)
	pageInt, err := strconv.Atoi(page)
	if err != nil{
		pageInt = 1
	}
	rowsInt, err = strconv.Atoi(rows)
	if err != nil {
		rowsInt = 100
	}
	transaction := models.ViewTransaction{}
	RequestId, err := strconv.Atoi(request.URL.Query().Get(`RequestId`))
	if err == nil {
		transaction.RequestId = int64(RequestId)
	}
	PaymentID, err := strconv.Atoi(request.URL.Query().Get(`PaymentID`))
	if err == nil {
		transaction.PaymentID =  int64(PaymentID)
	}
	PreCheckQueueID, err := strconv.Atoi(request.URL.Query().Get(`PreCheckQueueID`))
	if err == nil {
		transaction.PreCheckQueueID =  int64(PreCheckQueueID)
	}
	Vendor, err := strconv.Atoi(request.URL.Query().Get(`Vendor`))
	if err == nil {
		transaction.Vendor =  Vendor
	}
	VendorName := request.URL.Query().Get(`VendorName`)
	transaction.VendorName = VendorName
	RequestType := request.URL.Query().Get(`RequestType`)
	transaction.RequestType = RequestType
	AccountPayer := request.URL.Query().Get(`AccountPayer`)
	transaction.AccountPayer = AccountPayer
	AccountReceiver := request.URL.Query().Get(`AccountReceiver`)
	transaction.AccountReceiver = AccountReceiver
	StateID := request.URL.Query().Get(`StateID`)
	transaction.StateID = StateID
	Aggregator := request.URL.Query().Get(`Aggregator`)
	transaction.Aggregator = Aggregator
	GateWay := request.URL.Query().Get(`GateWay`)
	transaction.GateWay = GateWay
	Amount, err := strconv.ParseFloat(request.URL.Query().Get(`Amount`), 64)
	if err == nil {
		transaction.Amount =  Amount
	}
	response := models.GetViewTransactions(transaction, int64(rowsInt), int64(pageInt - 1))
	err = json.NewEncoder(writer).Encode(&response)
	if err != nil {
		log.Print(err)
	}
}
// Get view report for report
func (server *MainServer) GetViewReportsHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	fmt.Println("I am get clients")
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	pageInt := 1
	rowsInt := 100
	page := request.URL.Query().Get(`page`)
	rows := request.URL.Query().Get(`rows`)
	pageInt, err := strconv.Atoi(page)
	if err != nil{
		pageInt = 1
	}
	rowsInt, err = strconv.Atoi(rows)
	if err != nil {
		rowsInt = 100
	}
	response, err := models.GetViewReport(int64(rowsInt), int64(pageInt - 1))
	if err != nil {
		err := json.NewEncoder(writer).Encode([]string{`error mismatch this raport'`})
		log.Print(err)
		return
	}
	err = json.NewEncoder(writer).Encode(&response)
	if err != nil {
		log.Print(err)
	}
}
// Save New Vendor
func (server *MainServer) SaveNewVendorHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params){
	var requestBody models.Vendor
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(writer).Encode([]string{"err.json_invalid"})
		log.Print(err)
		return
	}
	response := requestBody.Save()
	if response.ID <= 0{
		writer.WriteHeader(http.StatusNotFound)
		err := json.NewEncoder(writer).Encode([]string{"err.json_invalid"})
		log.Print(err)
		return
	}
	err = json.NewEncoder(writer).Encode(&response)
	if err != nil {
		log.Print(err)
	}
}
// Update  Vendor
func (server *MainServer) UpdateVendorHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params){
	var requestBody models.Vendor
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(writer).Encode([]string{"err.json_invalid"})
		log.Print(err)
		return
	}
	response := requestBody.Update(requestBody)
	fmt.Println(response)
	if response.ID <= 0{
		writer.WriteHeader(http.StatusNotFound)
		err := json.NewEncoder(writer).Encode([]string{"err.json_invalid"})
		log.Print(err)
		return
	}
	err = json.NewEncoder(writer).Encode(&response)
	if err != nil {
		log.Println(err)
	}
}
//Get one Vendor
func (server *MainServer) GetVendorHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params){
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	param := params.ByName(`id`)
	id, err := strconv.Atoi(param)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	response := models.GetVendorById(int64(id))
	err = json.NewEncoder(writer).Encode(&response)
	if err != nil {
		log.Print(err)
	}
}
//Get list Merchants
func (server *MainServer) GetMerchantsHandler(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	page := request.URL.Query().Get(`page`)
	rows := request.URL.Query().Get(`rows`)
	pageInt, _ := strconv.Atoi(page)
	rowsInt, _ := strconv.Atoi(rows)
	var merchant models.Merchant
	response := merchant.GetMerchants(int64(rowsInt), int64(pageInt-1))
	err := json.NewEncoder(writer).Encode(&response)
	if err != nil {
		log.Print(err)
	}
}
//Get on Merchant
func (server *MainServer) GetMerchantHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params){
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	param := params.ByName(`id`)
	id, err := strconv.Atoi(param)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	response := models.Merchant{}
	merchant := response.GetMerchantById(int64(id))
	err = json.NewEncoder(writer).Encode(&merchant)
	if err != nil {
		log.Print(err)
	}
}
//Update Merchant
func (server *MainServer) UpdateMerchantHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params){
	var requestBody models.Merchant
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(writer).Encode([]string{"err.json_invalid"})
		log.Print(err)
		return
	}
	response := requestBody.Update(requestBody)
	fmt.Println(response)
	if response.ID <= 0{
		writer.WriteHeader(http.StatusNotFound)
		err := json.NewEncoder(writer).Encode([]string{"err.json_invalid"})
		log.Print(err)
		return
	}
	err = json.NewEncoder(writer).Encode(&response)
	if err != nil {
		log.Println(err)
	}
}
//Get all ViewLog by page
func (server *MainServer) GetViewLogsHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	pageInt := 1
	rowsInt := 100
	page := request.URL.Query().Get(`page`)
	rows := request.URL.Query().Get(`rows`)
	pageInt, err := strconv.Atoi(page)
	if err != nil{
		pageInt = 1
	}
	rowsInt, err = strconv.Atoi(rows)
	if err != nil {
		rowsInt = 100
	}
	response, err := models.GetViewLogs(int64(rowsInt), int64(pageInt))
	if err != nil {
		err := json.NewEncoder(writer).Encode([]string{`error mismatch this view log'`})
		log.Print(err)
		return
	}
	err = json.NewEncoder(writer).Encode(&response)
	if err != nil {
		log.Print(err)
	}
}
//GetViewDTO
func (server *MainServer) GetViewLogsDTOHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	pageInt := 1
	rowsInt := 100
	page := request.URL.Query().Get(`page`)
	rows := request.URL.Query().Get(`rows`)
	pageInt, err := strconv.Atoi(page)
	if err != nil{
		pageInt = 1
	}
	rowsInt, err = strconv.Atoi(rows)
	if err != nil {
		rowsInt = 100
	}
	response, err := models.GetViewLogsDTO(int64(rowsInt), int64(pageInt))
	if err != nil {
		err := json.NewEncoder(writer).Encode([]string{`error mismatch this view log'`})
		log.Print(err)
		return
	}
	err = json.NewEncoder(writer).Encode(&response)
	if err != nil {
		log.Print(err)
	}
}
//
func (server *MainServer) GetViewLogHandler(writer http.ResponseWriter, request *http.Request, param httprouter.Params){
	id , err := strconv.Atoi(param.ByName("id"))
	if err != nil {
		err := json.NewEncoder(writer).Encode("invalid_id")
		if err != nil {
			log.Print(err)
			return
		}
		return
	}
	viewLog, err := models.GetViewLogById(int64(id))
	if err != nil {
		err := json.NewEncoder(writer).Encode("invalid_id")
		if err != nil {
			log.Print(err)
			return
		}
		return
	}
	err = json.NewEncoder(writer).Encode(viewLog)
	if err != nil {
		log.Print("invalid_viewlog")
		return
	}
}
//
func (server *MainServer) GetHamsoyaTransactionTypeHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	pageInt := 1
	rowsInt := 100
	page := request.URL.Query().Get(`page`)
	rows := request.URL.Query().Get(`rows`)
	pageInt, err := strconv.Atoi(page)
	if err != nil{
		pageInt = 1
	}
	rowsInt, err = strconv.Atoi(rows)
	if err != nil {
		rowsInt = 100
	}
	response, err := server.userSvc.GetHamsoyaTransactionsType(int64(rowsInt), int64(pageInt))

	if err != nil {
		err := json.NewEncoder(writer).Encode([]string{`error mismatch this transaction type'`})
		log.Print(err)
		return
	}
	err = json.NewEncoder(writer).Encode(&response)
	if err != nil {
		log.Print(err)
	}
}

// TODO: найти проблему как распарсить время, чтобы парсилась из string в time.time
func (server *MainServer) GetHamsoyaTransactionsHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	pageInt := 1
	rowsInt := 100
	var transaction hamsoyamodels.HamsoyaTransaction
	page := request.URL.Query().Get(`page`)
	pageInt, err := strconv.Atoi(page)
	if err != nil{
		pageInt = 1
	}

	rows := request.URL.Query().Get(`rows`)
	rowsInt, err = strconv.Atoi(rows)
	if err != nil {
		rowsInt = 100
	}
	id, err := strconv.Atoi(request.URL.Query().Get(`id`))
	if err == nil {
		transaction.Id = int64(id)
	}
	ClientPayerId, err := strconv.Atoi(request.URL.Query().Get(`clientpayerid`))
	if err == nil {
		transaction.ClientPayerId = int64(ClientPayerId)
		fmt.Println(ClientPayerId)
	}
	//PreCheckId    int64     `xml:"pre_check_id"`
	//StatusId      int64     `xml:"status_id"`
	//TypeId        int64     `xml:"type_id"`
	//ExtStatusId   int64     `xml:"ext_status_id"`
	//ExtTransId    string    `xml:"ext_trans_id"`
	//CreateDate    time.Time `xml:"create_date"`
	//LastUpdate    time.Time `xml:"last_update"`
	//Description   string    `xml:"description"`+
	//ClientPayerId int64     `xml:"client_payer_id"`
	//myDateString := "2019-10-30T01:07:39.085082+05:00"
	//myDateString := request.URL.Query().Get(`createdata`)
	//fmt.Println("My Starting Date:\t", myDateString)
	//	myDate, err := time.Parse( "2019-10-30T01:07:39.085082+05:00", myDateString)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("My Date Reformatted:\t", myDate)
	//transaction.CreateDate = myDate

	response, err := hamsoyamodels.GetHamsoyaTransactions(transaction, int64(rowsInt), int64(pageInt))

	if err != nil {
		err := json.NewEncoder(writer).Encode([]string{`error mismatch this transaction type`})
		log.Println(err)
		return
	}
	err = json.NewEncoder(writer).Encode(&response)
	if err != nil {
		log.Println(err)
	}
}
//Get transaction By id
func (server *MainServer) GetHamsoyaTransactionByIdHandler(writer http.ResponseWriter, request *http.Request, param httprouter.Params){
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	id , err := strconv.Atoi(param.ByName("id"))
	if err != nil {
		err := json.NewEncoder(writer).Encode("invalid_id")
		if err != nil {
			log.Print(err)
			return
		}
		return
	}
	transaction, err := hamsoyamodels.GetHamsoyaTransactionById(int64(id))
	if err != nil {
		err := json.NewEncoder(writer).Encode("invalid_id")
		if err != nil {
			log.Print(err)
			return
		}
		return
	}
	err = json.NewEncoder(writer).Encode(&transaction)
	if err != nil {
		log.Print("invalid_transaction")
		return
	}
}
//Get TransactionType By Id
func (server *MainServer) GetHamsoyaTransactionTypeByIdHandler(writer http.ResponseWriter, request *http.Request, param httprouter.Params){
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	id , err := strconv.Atoi(param.ByName("id"))
	if err != nil {
		err := json.NewEncoder(writer).Encode("invalid_id")
		if err != nil {
			log.Print(err)
			return
		}
		return
	}
	transaction, err := hamsoyamodels.GetHamsoyaTransactionTypeById(int64(id))
	if err != nil {
		err := json.NewEncoder(writer).Encode("invalid_id")
		if err != nil {
			log.Print(err)
			return
		}
		return
	}
	err = json.NewEncoder(writer).Encode(&transaction)
	if err != nil {
		log.Print("invalid_transaction")
		return
	}
}
//Save transactionType
func (server *MainServer) SaveHamsoyaTransactionType(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	var requestBody hamsoyamodels.HamsoyaTransactionType
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(writer).Encode("wrong_body")
		log.Print(err)
		return
	}
	fmt.Println(requestBody)
	response := requestBody.Save()
	if response.Id <= 0{
		writer.WriteHeader(http.StatusNotFound)
		err := json.NewEncoder(writer).Encode([]string{"err.json_invalid"})
		log.Print(err)
		return
	}
	err = json.NewEncoder(writer).Encode(&response)
	if err != nil {
		log.Print(err)
	}
}
//Edit transactionType
func (server *MainServer) UpdateHamsoyaTransactionTypeHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	fmt.Println("here")
	var requestBody hamsoyamodels.HamsoyaTransactionType
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(writer).Encode([]string{"err.json_invalid"})
		log.Print(err)
		return
	}
	response := requestBody.Update(requestBody)
	fmt.Println(response)
	if response.Id <= 0{
		writer.WriteHeader(http.StatusNotFound)
		err := json.NewEncoder(writer).Encode([]string{"err.json_invalid"})
		log.Print(err)
		return
	}
	err = json.NewEncoder(writer).Encode(&response)
	if err != nil {
		log.Println(err)
	}
}
// Get Configs
func (server *MainServer) GetHamosyaConfigsHandler (writer http.ResponseWriter, request *http.Request, _ httprouter.Params){
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	pageInt := 1
	rowsInt := 100
	page := request.URL.Query().Get(`page`)
	pageInt, err := strconv.Atoi(page)
	if err != nil{
		pageInt = 1
	}

	rows := request.URL.Query().Get(`rows`)
	rowsInt, err = strconv.Atoi(rows)
	if err != nil {
		rowsInt = 100
	}
	var newHamsoyaConfig hamsoyamodels.HamsoyaConfig
	id, err := strconv.Atoi(request.URL.Query().Get(`id`))
	if err == nil {
		newHamsoyaConfig.Id = int64(id)
	}
	code := request.URL.Query().Get(`code`)
	newHamsoyaConfig.Code = code

	value, err := strconv.Atoi(request.URL.Query().Get(`value`))
	if err == nil {
		newHamsoyaConfig.Value = int64(value)
	}

	valueStr := request.URL.Query().Get(`valuestr`)
	newHamsoyaConfig.ValueStr = valueStr

	HamsoyaConfig, err := hamsoyamodels.GetHamsoyaConfig(newHamsoyaConfig, int64(rowsInt), int64(pageInt))
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		err := json.NewEncoder(writer).Encode(`mismatch_hamsoyaConfig`)
		if err != nil {
			log.Println(err)
			return
		}
		return
	}
	err = json.NewEncoder(writer).Encode(HamsoyaConfig)
	if err != nil {
		log.Println(err)
		return
	}
	return
}
// Save Configs
func (server *MainServer) SaveHamsoyaConfigHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	var requestBody hamsoyamodels.HamsoyaConfig
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	fmt.Println(requestBody)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(writer).Encode("err.json_invalid")
		log.Println(err)
		return
	}
	requestBody.CreateDate = time.Now()
	response, err := requestBody.Save()
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		err := json.NewEncoder(writer).Encode("wrong_date")
		if err != nil {
			log.Println(err)
			return
		}
	}
	err = json.NewEncoder(writer).Encode(&response)
	if err != nil {
		log.Println(err)
		return
	}
	return
}
// TODO: Edit Configs check for id and configs.id
func (server *MainServer) UpdateHamsoyaConfigHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	var requestBody hamsoyamodels.HamsoyaConfig
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	fmt.Println(requestBody)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(writer).Encode("err.json_invalid")
		log.Println(err)
		return
	}
	response, err := requestBody.Update(requestBody)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		err := json.NewEncoder(writer).Encode("wrong_date")
		if err != nil {
			log.Println(err)
			return
		}
	}
	err = json.NewEncoder(writer).Encode(&response)
	if err != nil {
		log.Println(err)
		return
	}
	return
}
// TODO: GET ONE config
func (server *MainServer) GetHamsoyaConfigByIdHandler(writer http.ResponseWriter, request *http.Request, param httprouter.Params){
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	id , err := strconv.Atoi(param.ByName("id"))
	if err != nil {
		err := json.NewEncoder(writer).Encode("invalid_id")
		if err != nil {
			log.Print(err)
			return
		}
		return
	}
	transaction, err := hamsoyamodels.GetHamsoyaConfigById(int64(id))
	if err != nil {
		err := json.NewEncoder(writer).Encode("invalid_id")
		if err != nil {
			log.Print(err)
			return
		}
		return
	}
	err = json.NewEncoder(writer).Encode(&transaction)
	if err != nil {
		log.Print("invalid_config.")
		return
	}
}