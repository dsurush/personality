package app

import (
	"MF/models"
	"MF/token"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func (server *MainServer) LoginHandler(writer http.ResponseWriter, request *http.Request, pr httprouter.Params) {
	fmt.Println("login\n")
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

func (server *MainServer) GetClientsInfoByFilterHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
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
//UnUse
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
//UnUse
func (server *MainServer) MainPageHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	fmt.Println("Login\n")
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
//UnUse
func (server *MainServer) GetVendorCategoryHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	var vendor models.Vendor
	vendors := vendor.FindAll()
	//fmt.Println("Hello I am vendors\n", vendors)
	err := json.NewEncoder(writer).Encode(&vendors)
	if err != nil {
		log.Print(err)
	}
}

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
//UNUSE
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