package app

import (
	"MF/hamsoyamodels"
	"MF/helperfunc"
	"MF/models"
	"MF/token"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"log"
	"math"
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
	log.Printf("login = %s, pass = %s\n", requestBody.Username, requestBody.Password)
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

//Get clients By ID
func (server *MainServer) GetClientInfoByIdHandler(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	fmt.Println("I am find client By number phone")
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	id := param.ByName(`id`)
	fmt.Println(id)
	response := models.GetClientInfoById(id)
	if response.ClientID == 0 {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(response)
	err := json.NewEncoder(writer).Encode(&response)
	if err != nil {
		log.Print(err)
	}
}

//Get list clients Handler ::: TODO CHANGE
func (server *MainServer) GetClientsInfoHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	var clientDefault models.ClientInfo
	URL := `http://127.0.0.1:8080/api/megafon/clients`
	PreURL := ``
	page := request.URL.Query().Get(`page`)
	pageInt, errPage := strconv.Atoi(page)
	if errPage != nil {
		pageInt = 1
		URL += `?page=1`
	}
	IsActive, err := strconv.ParseBool(request.URL.Query().Get(`IsActive`))
	if err == nil {
		clientDefault.IsActive = IsActive
		PreURL += fmt.Sprintf(`&IsActive=%s`, request.URL.Query().Get(`IsActive`))
	}
	IsIdentified, err := strconv.ParseBool(request.URL.Query().Get(`IsIdentified`))
	if err == nil {
		clientDefault.IsIdentified = IsIdentified
		PreURL += fmt.Sprintf(`&IsIdentified=%s`, request.URL.Query().Get(`IsIdentified`))
	}
	IsBlackList, err := strconv.ParseBool(request.URL.Query().Get(`IsBlackList`))
	if err == nil {
		clientDefault.IsBlackList = IsBlackList
		PreURL += fmt.Sprintf(`&IsBlackList=%s`, request.URL.Query().Get(`IsBlackList`))
	}
	SendToCft, err := strconv.ParseBool(request.URL.Query().Get(`SendToCft`))
	if err == nil {
		clientDefault.SendToCft = SendToCft
		PreURL += fmt.Sprintf(`&SendToCft=%s`, request.URL.Query().Get(`SendToCft`))
	}
	Sex := request.URL.Query().Get(`Sex`)
	if Sex != `` {
		if Sex == "W" {
			clientDefault.Sex = "Ж"
			PreURL += fmt.Sprintf(`&Sex=%s`, `W`)
		} else {
			clientDefault.Sex = "М"
			PreURL += fmt.Sprintf(`&Sex=%s`, `M`)
		}
	}

	//Default-ные значение времени. типа с начало до конца
	var interval helperfunc.TimeInterval
	unix := time.Unix(0, 0)
	interval.From = unix.Format(time.RFC3339)
	unixTimeNow := time.Now()
	interval.To = unixTimeNow.Format(time.RFC3339)
	from, err := strconv.Atoi(request.URL.Query().Get(`from`))
	if err == nil {
		from /= 1000
		i := time.Unix(int64(from), 0)
		ans := i.Format(time.RFC3339)
		interval.From = ans
		PreURL += `&from=` + request.URL.Query().Get(`from`)
	}
	to, err := strconv.Atoi(request.URL.Query().Get(`to`))
	if err == nil {
		from /= 1000
		i := time.Unix(int64(to), 0)
		ans := i.Format(time.RFC3339)
		interval.To = ans
		PreURL += `&to=` + request.URL.Query().Get(`to`)
	}

	var response models.ResponseClientsInfo
	response = models.GetClientsCount(clientDefault, interval)
	response.TotalPage = int64(math.Ceil(float64(response.TotalPage) / float64(int64(100))))
	response.Page = helperfunc.MinOftoInt(int64(pageInt), response.TotalPage)
	if errPage == nil {
		URL += `?page=` + fmt.Sprintf("%d", response.Page)
	}
	response.URL = URL + PreURL
	fmt.Println(response.URL)
	models.GetClients(clientDefault, &response, interval, response.Page-1)
	err = json.NewEncoder(writer).Encode(&response)
	if err != nil {
		log.Print(err)
	}
}

//UnUse Handler
func (server *MainServer) LoginHandler1(writer http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	fmt.Printf("Login\n")
	bytes, err := ioutil.ReadFile("web/templates/mpage.gohtml")
	if err != nil {
		log.Fatal("can't read from /web/templates/mgpage.gohtml\nerr: ", err)
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
	response, err := server.tokenSvc.Generate(request.Context(), &requestBody)
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

//GetVendorCategory
func (server *MainServer) GetVendorsHandler(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	URL := `http://127.0.0.1:8080/api/megafon/clients`
	//	PreURL := ``
	page := request.URL.Query().Get(`page`)
	pageInt, errPage := strconv.Atoi(page)
	if errPage != nil || pageInt <= 0 {
		pageInt = 1
		URL += `?page=1`
	}
	var vendor models.Vendor
	response := vendor.FindAll(pageInt - 1)
	err := json.NewEncoder(writer).Encode(&response)
	if err != nil {
		log.Print(err)
	}
}

// Get view Trans for report
func (server *MainServer) GetViewTransactionsHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	fmt.Println("I am view Transaction")
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	transaction := models.ViewTransaction{}
	PreURL := ``
	RequestId, err := strconv.Atoi(request.URL.Query().Get(`RequestId`))
	if err == nil {
		transaction.RequestId = int64(RequestId)
		PreURL += `&RequestId=` + request.URL.Query().Get(`RequestId`)
	}
	PaymentID, err := strconv.Atoi(request.URL.Query().Get(`PaymentID`))
	if err == nil {
		transaction.PaymentID = int64(PaymentID)
		PreURL += `&PaymentID=` + request.URL.Query().Get(`PaymentID`)
	}
	PreCheckQueueID, err := strconv.Atoi(request.URL.Query().Get(`PreCheckQueueID`))
	if err == nil {
		transaction.PreCheckQueueID = int64(PreCheckQueueID)
		PreURL += `&PreCheckQueueID=` + request.URL.Query().Get(`PreCheckQueueID`)
	}
	Vendor, err := strconv.Atoi(request.URL.Query().Get(`Vendor`))
	if err == nil {
		transaction.Vendor = Vendor
		PreURL += `&Vendor=` + request.URL.Query().Get(`Vendor`)
	}
	VendorName := request.URL.Query().Get(`VendorName`)
	if VendorName != `` {
		PreURL += `&VendorName=` + request.URL.Query().Get(`VendorName`)
		transaction.VendorName = VendorName
	}
	RequestType := request.URL.Query().Get(`RequestType`)
	if RequestType != `` {
		PreURL += `&RequestType=` + request.URL.Query().Get(`RequestType`)
		transaction.RequestType = RequestType
	}
	AccountPayer := request.URL.Query().Get(`AccountPayer`)
	if AccountPayer != `` {
		PreURL += `&AccountPayer=` + request.URL.Query().Get(`AccountPayer`)
		transaction.AccountPayer = AccountPayer
	}
	AccountReceiver := request.URL.Query().Get(`AccountReceiver`)
	if AccountReceiver != `` {
		PreURL += `&AccountReceiver=` + request.URL.Query().Get(`AccountReceiver`)
		transaction.AccountReceiver = AccountReceiver
	}
	StateID := request.URL.Query().Get(`StateID`)
	if StateID != `` {
		PreURL += `&StateID=` + request.URL.Query().Get(`StateID`)
		transaction.StateID = StateID
	}
	Aggregator := request.URL.Query().Get(`Aggregator`)
	if Aggregator != `` {
		PreURL += `&Aggregator=` + request.URL.Query().Get(`Aggregator`)
		transaction.Aggregator = Aggregator
	}
	GateWay := request.URL.Query().Get(`GateWay`)
	if GateWay != `` {
		PreURL += `&GateWay=` + request.URL.Query().Get(`GateWay`)
		transaction.GateWay = GateWay
	}

	Amount, err := strconv.ParseFloat(request.URL.Query().Get(`Amount`), 64)
	if err == nil {
		PreURL += `&Amount=` + request.URL.Query().Get(`Amount`)
		transaction.Amount = Amount
	}

	var interval helperfunc.TimeInterval
	unix := time.Unix(0, 0)
	interval.From = unix.Format(time.RFC3339)
	unixTimeNow := time.Now()
	interval.To = unixTimeNow.Format(time.RFC3339)
	from, err := strconv.Atoi(request.URL.Query().Get(`from`))
	if err == nil {
		i := time.Unix(int64(from), 0)
		ans := i.Format(time.RFC3339)
		interval.From = ans
		PreURL += `&from=` + request.URL.Query().Get(`from`)
	}
	to, err := strconv.Atoi(request.URL.Query().Get(`to`))
	if err == nil {
		i := time.Unix(int64(to), 0)
		ans := i.Format(time.RFC3339)
		interval.To = ans
		PreURL += `&to=` + request.URL.Query().Get(`to`)
	}

	var response models.ResponseViewTransactions
	page := request.URL.Query().Get(`page`)
	URL := `http://127.0.0.1:8080/api/megafon/transactions`
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 1
		URL += `?page=1`
	}
	response = models.GetViewTransactionsCount(transaction, interval)
	response.TotalPage = int64(math.Ceil(float64(response.TotalPage) / float64(int64(100))))
	response.Page = helperfunc.MinOftoInt(int64(pageInt), response.TotalPage)
	if err == nil {
		URL += `?page=` + fmt.Sprintf("%d", response.Page)
	}
	URL += PreURL
	response.URL = URL
	fmt.Println(URL)
	models.GetViewTransactions(transaction, &response, interval, response.Page-1)
	err = json.NewEncoder(writer).Encode(&response)
	if err != nil {
		log.Print(err)
	}
}

// Get view report for report
func (server *MainServer) GetViewReportsHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	fmt.Println("I am get clients")
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	viewReport := models.ViewReport{}
	PreURL := ``
	RequestId, err := strconv.Atoi(request.URL.Query().Get(`RequestId`))
	if err == nil {
		viewReport.RequestId = int64(RequestId)
		PreURL += `&RequestId=` + request.URL.Query().Get(`RequestId`)
	}
	PaymentID, err := strconv.Atoi(request.URL.Query().Get(`PaymentID`))
	if err == nil {
		viewReport.PaymentID = int64(PaymentID)
		PreURL += `&PaymentID=` + request.URL.Query().Get(`PaymentID`)
	}
	VendorID, err := strconv.Atoi(request.URL.Query().Get(`VendorId`))
	if err == nil {
		viewReport.VendorID = VendorID
		PreURL += `&VendorId=` + request.URL.Query().Get(`VendorId`)
	}
	VendorName := request.URL.Query().Get(`VendorName`)
	if VendorName != `` {
		PreURL += `&VendorName=` + request.URL.Query().Get(`VendorName`)
		viewReport.VendorName = VendorName
	}
	RequestType := request.URL.Query().Get(`RequestType`)
	if RequestType != `` {
		PreURL += `&RequestType=` + request.URL.Query().Get(`RequestType`)
		viewReport.RequestType = RequestType
	}
	AccountPayer := request.URL.Query().Get(`AccountPayer`)
	if AccountPayer != `` {
		PreURL += `&AccountPayer=` + request.URL.Query().Get(`AccountPayer`)
		viewReport.AccountPayer = AccountPayer
	}
	AccountReceiver := request.URL.Query().Get(`AccountReceiver`)
	if AccountReceiver != `` {
		PreURL += `&AccountReceiver=` + request.URL.Query().Get(`AccountReceiver`)
		viewReport.AccountReceiver = AccountReceiver
	}
	StateID := request.URL.Query().Get(`StateID`)
	if StateID != `` {
		PreURL += `&StateID=` + request.URL.Query().Get(`StateID`)
		viewReport.StateID = StateID
	}
	Aggregator := request.URL.Query().Get(`Aggregator`)
	if Aggregator != `` {
		PreURL += `&Aggregator=` + request.URL.Query().Get(`Aggregator`)
		viewReport.Aggregator = Aggregator
	}
	GateWay := request.URL.Query().Get(`GateWay`)
	if GateWay != `` {
		PreURL += `&GateWay=` + request.URL.Query().Get(`GateWay`)
		viewReport.GateWay = GateWay
	}

	Amount, err := strconv.ParseFloat(request.URL.Query().Get(`Amount`), 64)
	if err == nil {
		PreURL += `&Amount=` + request.URL.Query().Get(`Amount`)
		viewReport.Amount = Amount
	}
	fmt.Println(viewReport)

	var interval helperfunc.TimeInterval
	unix := time.Unix(0, 0)
	interval.From = unix.Format(time.RFC3339)
	unixTimeNow := time.Now()
	interval.To = unixTimeNow.Format(time.RFC3339)
	from, err := strconv.Atoi(request.URL.Query().Get(`from`))
	if err == nil {
		i := time.Unix(int64(from), 0)
		ans := i.Format(time.RFC3339)
		interval.From = ans
		PreURL += `&from=` + request.URL.Query().Get(`from`)
	}
	to, err := strconv.Atoi(request.URL.Query().Get(`to`))
	if err == nil {
		i := time.Unix(int64(to), 0)
		ans := i.Format(time.RFC3339)
		interval.To = ans
		PreURL += `&to=` + request.URL.Query().Get(`to`)
	}

	var response models.ResponseViewReports
	page := request.URL.Query().Get(`page`)
	URL := `http://127.0.0.1:8080/api/megafon/reports`
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 1
		URL += `?page=1`
	}
	response = models.GetViewReportCount(viewReport, interval)
	response.TotalPage = int64(math.Ceil(float64(response.TotalPage) / float64(int64(100))))
	response.Page = helperfunc.MinOftoInt(int64(pageInt), response.TotalPage)
	if err == nil {
		URL += `?page=` + fmt.Sprintf("%d", response.Page)
	}
	URL += PreURL
	response.URL = URL
	fmt.Println(URL)
	models.GetViewReport(viewReport, &response, interval, response.Page-1)
	err = json.NewEncoder(writer).Encode(&response)
	if err != nil {
		log.Print(err)
	}
}

// Save New Vendor
func (server *MainServer) SaveNewVendorHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
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
	if response.ID <= 0 {
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
func (server *MainServer) UpdateVendorHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
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
	if response.ID <= 0 {
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
func (server *MainServer) GetVendorHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
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
	PreURL := ``
	//
	var interval helperfunc.TimeInterval
	unix := time.Unix(0, 0)
	interval.From = unix.Format(time.RFC3339)
	unixTimeNow := time.Now()
	interval.To = unixTimeNow.Format(time.RFC3339)
	from, err := strconv.Atoi(request.URL.Query().Get(`from`))
	if err == nil {
		from /= 1000
		i := time.Unix(int64(from), 0)
		ans := i.Format(time.RFC3339)
		interval.From = ans
		PreURL += `&from=` + request.URL.Query().Get(`from`)
	}
	to, err := strconv.Atoi(request.URL.Query().Get(`to`))
	if err == nil {
		from /= 1000
		i := time.Unix(int64(to), 0)
		ans := i.Format(time.RFC3339)
		interval.To = ans
		PreURL += `&to=` + request.URL.Query().Get(`to`)
	}
	var response models.ResponseMerchants
	pageInt, err := strconv.Atoi(page)
	URL := `http://127.0.0.1:8080/api/megafon/merchants`
	if err != nil {
		pageInt = 1
		URL += `?page=1`
	}
	var merchant models.Merchant
	response = models.GetMerchantsCount(merchant, interval)
	response.TotalPage = int64(math.Ceil(float64(response.TotalPage) / float64(int64(100))))
	response.Page = helperfunc.MinOftoInt(int64(pageInt), response.TotalPage)
	if err == nil {
		URL += `?page=` + fmt.Sprintf("%d", response.Page)
	}
	URL += PreURL
	response.URL = URL
	fmt.Println(URL)
	models.GetMerchants(merchant, &response, interval, response.Page-1)
	err = json.NewEncoder(writer).Encode(&response)
	if err != nil {
		log.Print(err)
	}
}

//Get on Merchant
func (server *MainServer) GetMerchantHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
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
func (server *MainServer) UpdateMerchantHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	var requestBody models.MerchantDTO
	var reqBody models.Merchant
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	reqBody.ID = requestBody.ID
	reqBody.HumoOnlineID = requestBody.HumoOnlineID
	reqBody.NameENG = requestBody.NameENG
	reqBody.NameRUS = requestBody.NameRUS
	reqBody.QrCode = requestBody.QrCode
	reqBody.QrCodeNew = requestBody.QrCodeNew
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(writer).Encode([]string{"err.json_invalid"})
		log.Print(err)
		return
	}
	fmt.Println(reqBody)
	response := reqBody.Update(reqBody)
	fmt.Println(response)
	if response.ID <= 0 {
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
	if err != nil {
		pageInt = 1
	}
	rowsInt, err = strconv.Atoi(rows)
	if err != nil {
		rowsInt = 100
	}
	response := models.GetViewLogs(int64(rowsInt), int64(pageInt))
	if response.Error != nil {
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
	page := request.URL.Query().Get(`page`)
	PreURL := ``
	//
	var interval helperfunc.TimeInterval
	unix := time.Unix(0, 0)
	interval.From = unix.Format(time.RFC3339)
	unixTimeNow := time.Now()
	interval.To = unixTimeNow.Format(time.RFC3339)
	from, err := strconv.Atoi(request.URL.Query().Get(`from`))
	if err == nil {
		from /= 1000
		i := time.Unix(int64(from), 0)
		ans := i.Format(time.RFC3339)
		interval.From = ans
		PreURL += `&from=` + request.URL.Query().Get(`from`)
	}
	to, err := strconv.Atoi(request.URL.Query().Get(`to`))
	if err == nil {
		from /= 1000
		i := time.Unix(int64(to), 0)
		ans := i.Format(time.RFC3339)
		interval.To = ans
		PreURL += `&to=` + request.URL.Query().Get(`to`)
	}
	var response models.ResponseViewLogsList
	pageInt, err := strconv.Atoi(page)
	URL := `http://127.0.0.1:8080/api/megafon/logs`
	if err != nil {
		pageInt = 1
		URL += `?page=1`
	}

	var ViewLogDTO models.ViewLogDTO
	response = models.GetViewLogsDTOCount(ViewLogDTO, interval)
	response.TotalPage = int64(math.Ceil(float64(response.TotalPage) / float64(int64(100))))
	response.Page = helperfunc.MinOftoInt(int64(pageInt), response.TotalPage)
	if err == nil {
		URL += `?page=` + fmt.Sprintf("%d", response.Page)
	}
	URL += PreURL
	response.URL = URL
	fmt.Println(URL)
	models.GetViewLogsDTO(ViewLogDTO, &response, interval, response.Page-1)
	err = json.NewEncoder(writer).Encode(&response)
	if err != nil {
		log.Print(err)
	}
}

//
func (server *MainServer) GetViewLogHandler(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	id, err := strconv.Atoi(param.ByName("id"))
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
func (server *MainServer) GetHamsoyaTransactionsHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	fmt.Println("I am view Transaction")
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	transaction := hamsoyamodels.HamsoyaTransaction{}
	PreURL := ``
	Id, err := strconv.Atoi(request.URL.Query().Get(`Id`))
	if err == nil {
		transaction.Id = int64(Id)
		PreURL += `&Id=` + request.URL.Query().Get(`Id`)
	}
	PreCheckId, err := strconv.Atoi(request.URL.Query().Get(`PreCheckId`))
	if err == nil {
		transaction.PreCheckId = int64(PreCheckId)
		PreURL += `&PreCheckId=` + request.URL.Query().Get(`PreCheckId`)
	}
	StatusId, err := strconv.Atoi(request.URL.Query().Get(`StatusId`))
	if err == nil {
		transaction.StatusId = int64(StatusId)
		PreURL += `&StatusId=` + request.URL.Query().Get(`StatusId`)
	}
	TypeId, err := strconv.Atoi(request.URL.Query().Get(`TypeId`))
	if err == nil {
		transaction.TypeId = int64(TypeId)
		PreURL += `&TypeId=` + request.URL.Query().Get(`TypeId`)
	}
	ExtStatusId, err := strconv.Atoi(request.URL.Query().Get(`ExtStatusId`))
	if err == nil {
		transaction.ExtStatusId = int64(ExtStatusId)
		PreURL += `&ExtStatusId=` + request.URL.Query().Get(`ExtStatusId`)
	}
	ClientPayerId, err := strconv.Atoi(request.URL.Query().Get(`ClientPayerId`))
	if err == nil {
		transaction.ClientPayerId = int64(ClientPayerId)
		PreURL += `&ClientPayerId=` + request.URL.Query().Get(`ClientPayerId`)
	}
	ExtTransId := request.URL.Query().Get(`ExtTransId`)
	if ExtTransId != `` {
		transaction.ExtTransId = ExtTransId
		PreURL += `&ExtTransId=` + request.URL.Query().Get(`ExtTransId`)
	}

	var interval helperfunc.TimeInterval
	unix := time.Unix(0, 0)
	interval.From = unix.Format(time.RFC3339)
	unixTimeNow := time.Now()
	interval.To = unixTimeNow.Format(time.RFC3339)
	from, err := strconv.Atoi(request.URL.Query().Get(`from`))
	if err == nil {
		i := time.Unix(int64(from), 0)
		ans := i.Format(time.RFC3339)
		interval.From = ans
		PreURL += `&from=` + request.URL.Query().Get(`from`)
	}
	to, err := strconv.Atoi(request.URL.Query().Get(`to`))
	if err == nil {
		i := time.Unix(int64(to), 0)
		ans := i.Format(time.RFC3339)
		interval.To = ans
		PreURL += `&to=` + request.URL.Query().Get(`to`)
	}

	var response hamsoyamodels.ResponseHamsoyaTransactions
	page := request.URL.Query().Get(`page`)
	URL := `http://127.0.0.1:8080/api/hamsoya/transactions`
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 1
		URL += `?page=1`
	}
	response = hamsoyamodels.GetHamsoyaTransactionsCount(transaction, interval)
	response.TotalPage = int64(math.Ceil(float64(response.TotalPage) / float64(int64(100))))
	response.Page = helperfunc.MinOftoInt(int64(pageInt), response.TotalPage)
	if err == nil {
		URL += `?page=` + fmt.Sprintf("%d", response.Page)
	}
	URL += PreURL
	response.URL = URL
	fmt.Println(URL)
	hamsoyamodels.GetHamsoyaTransactions(transaction, &response, interval, response.Page-1)
	err = json.NewEncoder(writer).Encode(&response)
	if err != nil {
		log.Print(err)
	}
}

//Get transaction By id
func (server *MainServer) GetHamsoyaTransactionByIdHandler(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	id, err := strconv.Atoi(param.ByName("id"))
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

//
func (server *MainServer) GetHamsoyaTransactionTypeHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	page := request.URL.Query().Get(`page`)
	PreURL := ``
	pageInt, err := strconv.Atoi(page)
	URL := `http://localhost:3000/hamsoya/transactionstype`
	if err != nil {
		pageInt = 1
		PreURL = `?page=1`
	} else {
		PreURL = fmt.Sprintf(`?page=%s`, pageInt)
	}
	var transactionTypeDefault hamsoyamodels.HamsoyaTransactionType
	IsActive, err := strconv.ParseBool(request.URL.Query().Get(`IsActive`))
	if err == nil {
		transactionTypeDefault.IsActive = IsActive
		PreURL += fmt.Sprintf(`&IsActive=%s`, request.URL.Query().Get(`IsActive`))
	}
	IsForJob, err := strconv.ParseBool(request.URL.Query().Get(`IsForJob`))
	if err == nil {
		transactionTypeDefault.IsForJob = IsForJob
		PreURL += fmt.Sprintf(`&IsForJob=%s`, request.URL.Query().Get(`IsForJob`))
	}
	IsPayment, err := strconv.ParseBool(request.URL.Query().Get(`IsPayment`))
	if err == nil {
		transactionTypeDefault.IsPayment = IsPayment
		PreURL += fmt.Sprintf(`&IsPayment=%s`, request.URL.Query().Get(`IsPayment`))
	}

	fmt.Println(transactionTypeDefault)

	var response hamsoyamodels.ResponseHamsoyaTransactionsType
	response = hamsoyamodels.GetHamsoyaTransactionsTypeCount(transactionTypeDefault)
	response.TotalPage = int64(math.Ceil(float64(response.TotalPage) / float64(int64(100))))
	response.Page = helperfunc.MinOftoInt(int64(pageInt), response.TotalPage)
	response.URL = URL + PreURL
	fmt.Println(response.URL)
	hamsoyamodels.GetHamsoyaTransactionsType(transactionTypeDefault, &response, int64(pageInt) - 1)

	err = json.NewEncoder(writer).Encode(&response)
	if err != nil {
		log.Print(err)
	}
}

//Get TransactionType By Id
func (server *MainServer) GetHamsoyaTransactionTypeByIdHandler(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	id, err := strconv.Atoi(param.ByName("id"))
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
	if response.Id <= 0 {
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
	if response.Id <= 0 {
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
func (server *MainServer) GetHamosyaConfigsHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	pageInt := 1
	rowsInt := 100
	page := request.URL.Query().Get(`page`)
	pageInt, err := strconv.Atoi(page)
	if err != nil {
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

	HamsoyaConfig := hamsoyamodels.GetHamsoyaConfig(newHamsoyaConfig, int64(rowsInt), int64(pageInt))
	if HamsoyaConfig.Error != nil {
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
	//	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
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

// GET ONE config
func (server *MainServer) GetHamsoyaConfigByIdHandler(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	id, err := strconv.Atoi(param.ByName("id"))
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

// Get AccountTypes
func (server *MainServer) GetHamosyaAccountTypesHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	pageInt := 1
	rowsInt := 100
	page := request.URL.Query().Get(`page`)
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 1
	}
	rows := request.URL.Query().Get(`rows`)
	rowsInt, err = strconv.Atoi(rows)
	if err != nil {
		rowsInt = 100
	}
	var newHamsoyaAccountType hamsoyamodels.HamsoyaAccountType
	id, err := strconv.Atoi(request.URL.Query().Get(`id`))
	if err == nil {
		newHamsoyaAccountType.Id = int64(id)
	}
	code := request.URL.Query().Get(`code`)
	newHamsoyaAccountType.Code = code

	Type := request.URL.Query().Get(`type`)
	newHamsoyaAccountType.Type = Type

	Name := request.URL.Query().Get(`name`)
	newHamsoyaAccountType.Name = Name

	prefix, err := strconv.Atoi(request.URL.Query().Get(`prefix`))
	if err == nil {
		newHamsoyaAccountType.Prefix = int64(prefix)
	}

	HamsoyaAccountTypes := hamsoyamodels.GetHamsoyaAccountTypes(newHamsoyaAccountType, int64(rowsInt), int64(pageInt))

	if HamsoyaAccountTypes.Error != nil {
		writer.WriteHeader(http.StatusNotFound)
		err := json.NewEncoder(writer).Encode(`mismatch_hamsoyaConfig`)
		if err != nil {
			log.Println(err)
			return
		}
		return
	}
	err = json.NewEncoder(writer).Encode(HamsoyaAccountTypes)
	if err != nil {
		log.Println(err)
		return
	}
	return
}

// Get AccountType
func (server *MainServer) GetHamsoyaAccountTypeByIdHandler(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	id, err := strconv.Atoi(param.ByName("id"))
	if err != nil {
		err := json.NewEncoder(writer).Encode("invalid_id")
		if err != nil {
			log.Print(err)
			return
		}
		return
	}
	AccountTypes, err := hamsoyamodels.GetHamsoyaAccountTypeById(int64(id))
	if err != nil {
		err := json.NewEncoder(writer).Encode("invalid_id")
		if err != nil {
			log.Print(err)
			return
		}
		return
	}
	err = json.NewEncoder(writer).Encode(&AccountTypes)
	if err != nil {
		log.Print("invalid_config.")
		return
	}
}

// Save AccountType
func (server *MainServer) SaveHamsoyaAccountTypeHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	var requestBody hamsoyamodels.HamsoyaAccountType
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	fmt.Println(requestBody)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(writer).Encode("err.json_invalid")
		log.Println(err)
		return
	}
	//	requestBody.CreateDate = time.Now()
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

// Edit AccountType
func (server *MainServer) UpdateHamsoyaAccountTypeHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	//	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	var requestBody hamsoyamodels.HamsoyaAccountType
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

// Get Statuses
func (server *MainServer) GetHamsoyaStatusesHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	pageInt := 1
	page := request.URL.Query().Get(`page`)
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 1
	}
	var newHamsoyaStatus hamsoyamodels.HamsoyaStatus
	id, err := strconv.Atoi(request.URL.Query().Get(`id`))
	if err == nil {
		newHamsoyaStatus.Id = int64(id)
	}
	code := request.URL.Query().Get(`code`)
	newHamsoyaStatus.Code = code

	Name := request.URL.Query().Get(`name`)
	newHamsoyaStatus.Name = Name

	ExtCode := request.URL.Query().Get(`extcode`)
	newHamsoyaStatus.ExtCode = ExtCode

	resultCode, err := strconv.Atoi(request.URL.Query().Get(`resultcode`))
	if err == nil {
		newHamsoyaStatus.ResultCode = int64(resultCode)
	}

	final, err := strconv.ParseBool(request.URL.Query().Get(`final`))
	if err == nil {
		newHamsoyaStatus.Final = final
	}

	IsAmountHold, err := strconv.ParseBool(request.URL.Query().Get(`is_amount_hold`))
	if err == nil {
		newHamsoyaStatus.IsAmountHold = IsAmountHold
	}

	HamsoyaStatus := hamsoyamodels.GetHamsoyaStatuses(newHamsoyaStatus, int64(pageInt))

	if HamsoyaStatus.Error != nil {
		writer.WriteHeader(http.StatusNotFound)
		err := json.NewEncoder(writer).Encode(`mismatch_hamsoyaStatus`)
		if err != nil {
			log.Println(err)
			return
		}
		return
	}
	err = json.NewEncoder(writer).Encode(HamsoyaStatus)
	if err != nil {
		log.Println(err)
		return
	}
	return
}

// Get Status
func (server *MainServer) GetHamsoyaStatusHandler(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	id, err := strconv.Atoi(param.ByName("id"))
	if err != nil {
		err := json.NewEncoder(writer).Encode("invalid_id")
		if err != nil {
			log.Print(err)
			return
		}
		return
	}
	HamsoyaStatus, err := hamsoyamodels.GetHamsoyaStatusById(int64(id))
	if err != nil {
		err := json.NewEncoder(writer).Encode("invalid_id")
		if err != nil {
			log.Print(err)
			return
		}
		return
	}
	err = json.NewEncoder(writer).Encode(&HamsoyaStatus)
	if err != nil {
		log.Print("invalid_HamsoyaStatus.")
		return
	}
}

// Save Status
func (server *MainServer) SaveHamsoyaStatusHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	var requestBody hamsoyamodels.HamsoyaStatus
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	fmt.Println(requestBody)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(writer).Encode("err.json_invalid")
		log.Println(err)
		return
	}
	//	requestBody.CreateDate = time.Now()
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

// Edit Status
func (server *MainServer) UpdateHamsoyaStatusHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	var requestBody hamsoyamodels.HamsoyaStatus
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

// Get Hamsoya view Trans
func (server *MainServer) GetHamsoyaViewTransHandler(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	id, err := strconv.Atoi(param.ByName("id"))
	if err != nil {
		err := json.NewEncoder(writer).Encode("invalid_id")
		if err != nil {
			log.Print(err)
			return
		}
		return
	}
	HamsoyaViewTrans, err := hamsoyamodels.GetHamsoyaViewTransById(int64(id))
	if err != nil {
		err := json.NewEncoder(writer).Encode("invalid_id")
		if err != nil {
			log.Print(err)
			return
		}
		return
	}
	err = json.NewEncoder(writer).Encode(&HamsoyaViewTrans)
	if err != nil {
		log.Print("invalid_HamsoyaViewTrans.")
		return
	}
}

// Get Hamsoya view Transes

// Get Hamsoya view Transaction
func (server *MainServer) GetHamsoyaViewTransactionHandler(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	id, err := strconv.Atoi(param.ByName("id"))
	if err != nil {
		err := json.NewEncoder(writer).Encode("invalid_id")
		if err != nil {
			log.Print(err)
			return
		}
		return
	}
	HamsoyaViewTransaction, err := hamsoyamodels.GetHamsoyaViewTransactionById(int64(id))
	if err != nil {
		err := json.NewEncoder(writer).Encode("invalid_id")
		if err != nil {
			log.Print(err)
			return
		}
		return
	}
	err = json.NewEncoder(writer).Encode(&HamsoyaViewTransaction)
	if err != nil {
		log.Print("invalid_HamsoyaViewTransaction.")
		return
	}
}

// Get Hamsoya view Transactions TODO: check this route
func (server *MainServer) GetHamsoyaViewTransactionsHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	pageInt := 1
	rowsInt := 100
	page := request.URL.Query().Get(`page`)
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 1
	}
	rows := request.URL.Query().Get(`rows`)
	rowsInt, err = strconv.Atoi(rows)
	if err != nil {
		rowsInt = 100
	}
	var newHamsoyaViewTransaction hamsoyamodels.HamsoyaViewTransaction

	HamsoyaViewTransaction := hamsoyamodels.GetHamsoyaViewTransactions(newHamsoyaViewTransaction, int64(rowsInt), int64(pageInt))

	err = json.NewEncoder(writer).Encode(HamsoyaViewTransaction)
	if err != nil {
		log.Println(err)
		return
	}
	return
}

// Get Hamsoya Document
func (server *MainServer) GetHamsoyaDocumentByIdHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(writer).Encode("invalid_id")
		if err != nil {
			log.Print(err)
			return
		}
		return
	}
	Document, err := hamsoyamodels.GetHamsoyaDocument(int64(id))
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		err := json.NewEncoder(writer).Encode("server wrong")
		if err != nil {
			log.Print(err)
			return
		}
		return
	}
	err = json.NewEncoder(writer).Encode(&Document)
	if err != nil {
		log.Print("invalid_HamsoyaViewTransaction.")
		return
	}
}

// Get Hamsoya Documents
func (server *MainServer) GetHamsoyaDocumentsHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	page := request.URL.Query().Get(`page`)
	PreURL := ``
	pageInt, err := strconv.Atoi(page)
	URL := `http://localhost:3000/hamsoya/documents`
	if err != nil {
		pageInt = 1
		PreURL = `?page=1`
	} else {
		PreURL = fmt.Sprintf(`?page=%s`, pageInt)
	}
	var documentDefault hamsoyamodels.HamsoyaDocument
	Id, err := strconv.Atoi(request.URL.Query().Get(`Id`))
	if err == nil {
		documentDefault.Id = int64(Id)
		PreURL += fmt.Sprintf(`&Id=%s`, request.URL.Query().Get(`Id`))
	}
	AccountDt, err := strconv.Atoi(request.URL.Query().Get(`AccountDt`))
	if err == nil {
		documentDefault.AccountDt = int64(AccountDt)
		PreURL += fmt.Sprintf(`&AccountDt=%s`, request.URL.Query().Get(`AccountDt`))
	}
	AccountCt, err := strconv.Atoi(request.URL.Query().Get(`AccountCt`))
	if err == nil {
		documentDefault.AccountCt = int64(AccountCt)
		PreURL += fmt.Sprintf(`&AccountCt=%s`, request.URL.Query().Get(`AccountCt`))
	}
	TransId, err := strconv.Atoi(request.URL.Query().Get(`TransId`))
	if err == nil {
		documentDefault.TransId = int64(TransId)
		PreURL += fmt.Sprintf(`&TransId=%s`, request.URL.Query().Get(`TransId`))
	}
	StatusId, err := strconv.Atoi(request.URL.Query().Get(`StatusId`))
	if err == nil {
		documentDefault.StatusId = int64(StatusId)
		PreURL += fmt.Sprintf(`&StatusId=%s`, request.URL.Query().Get(`StatusId`))
	}
	CancelDocId, err := strconv.Atoi(request.URL.Query().Get(`CancelDocId`))
	if err == nil {
		documentDefault.CancelDocId = int64(CancelDocId)
		PreURL += fmt.Sprintf(`&CancelDocId=%s`, request.URL.Query().Get(`CancelDocId`))
	}

	Amount, err := strconv.ParseFloat(request.URL.Query().Get(`Amount`), 64)
	if err == nil {
		documentDefault.Amount = Amount
		PreURL += fmt.Sprintf(`&Amount=%s`, request.URL.Query().Get(`Amount`))
	}
	var interval helperfunc.TimeInterval
	unix := time.Unix(0, 0)
	interval.From = unix.Format(time.RFC3339)
	unixTimeNow := time.Now()
	interval.To = unixTimeNow.Format(time.RFC3339)
	from, err := strconv.Atoi(request.URL.Query().Get(`from`))
	if err == nil {
		i := time.Unix(int64(from), 0)
		ans := i.Format(time.RFC3339)
		interval.From = ans
		PreURL += `&from=` + request.URL.Query().Get(`from`)
	}
	to, err := strconv.Atoi(request.URL.Query().Get(`to`))
	if err == nil {
		i := time.Unix(int64(to), 0)
		ans := i.Format(time.RFC3339)
		interval.To = ans
		PreURL += `&to=` + request.URL.Query().Get(`to`)
	}
	fmt.Println(documentDefault)

	var response hamsoyamodels.ResponseHamsoyaDocuments
	response = hamsoyamodels.GetHamsoyaDocumentsCount(documentDefault, interval)
	response.TotalPage = int64(math.Ceil(float64(response.TotalPage) / float64(int64(100))))
	response.Page = helperfunc.MinOftoInt(int64(pageInt), response.TotalPage)
	response.URL = URL + PreURL
	fmt.Println(response.URL)
	hamsoyamodels.GetHamsoyaDocuments(documentDefault, &response, interval, int64(pageInt) - 1)

	err = json.NewEncoder(writer).Encode(&response)
	if err != nil {
		log.Print(err)
	}
}

// Get Hamsoya Record
func (server *MainServer) GetHamsoyaRecordByIdHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(writer).Encode("invalid_id")
		if err != nil {
			log.Print(err)
			return
		}
		return
	}
	Record, err := hamsoyamodels.GetHamsoyaRecordById(int64(id))
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		err := json.NewEncoder(writer).Encode("server wrong")
		if err != nil {
			log.Print(err)
			return
		}
		return
	}
	err = json.NewEncoder(writer).Encode(&Record)
	if err != nil {
		log.Print("invalid_HamsoyaRecord.")
		return
	}
}

//Get Hamsoya Records
func (server *MainServer) GetHamsoyaRecordsHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	page := request.URL.Query().Get(`page`)
	PreURL := ``
	pageInt, err := strconv.Atoi(page)
	URL := `http://localhost:3000/hamsoya/records`
	if err != nil {
		pageInt = 1
		PreURL = `?page=1`
	} else {
		PreURL = fmt.Sprintf(`?page=%s`, pageInt)
	}
	var recordDefault hamsoyamodels.HamsoyaRecord
	Id, err := strconv.Atoi(request.URL.Query().Get(`Id`))
	if err == nil {
		recordDefault.Id = int64(Id)
		PreURL += fmt.Sprintf(`&Id=%s`, request.URL.Query().Get(`Id`))
	}
	AccountId, err := strconv.Atoi(request.URL.Query().Get(`AccountId`))
	if err == nil {
		recordDefault.AccountId = int64(AccountId)
		PreURL += fmt.Sprintf(`&AccountId=%s`, request.URL.Query().Get(`AccountId`))
	}

	DocumentId, err := strconv.Atoi(request.URL.Query().Get(`DocumentId`))
	if err == nil {
		recordDefault.DocumentId = int64(DocumentId)
		PreURL += fmt.Sprintf(`&DocumentId=%s`, request.URL.Query().Get(`DocumentId`))
	}
	Amount, err := strconv.ParseFloat(request.URL.Query().Get(`Amount`), 64)
	if err == nil {
		recordDefault.Amount = Amount
		PreURL += fmt.Sprintf(`&Amount=%s`, request.URL.Query().Get(`Amount`))
	}
	StartSaldo, err := strconv.ParseFloat(request.URL.Query().Get(`StartSaldo`), 64)
	if err == nil {
		recordDefault.StartSaldo = StartSaldo
		PreURL += fmt.Sprintf(`&StartSaldo=%s`, request.URL.Query().Get(`StartSaldo`))
	}
	Type := request.URL.Query().Get(`Type`)
	if Type != ``{
		recordDefault.Type = Type
		PreURL += fmt.Sprintf(`&Type=%s`, request.URL.Query().Get(`Type`))

	}

	var interval helperfunc.TimeInterval
	unix := time.Unix(0, 0)
	interval.From = unix.Format(time.RFC3339)
	unixTimeNow := time.Now()
	interval.To = unixTimeNow.Format(time.RFC3339)
	from, err := strconv.Atoi(request.URL.Query().Get(`from`))
	if err == nil {
		i := time.Unix(int64(from), 0)
		ans := i.Format(time.RFC3339)
		interval.From = ans
		PreURL += `&from=` + request.URL.Query().Get(`from`)
	}
	to, err := strconv.Atoi(request.URL.Query().Get(`to`))
	if err == nil {
		i := time.Unix(int64(to), 0)
		ans := i.Format(time.RFC3339)
		interval.To = ans
		PreURL += `&to=` + request.URL.Query().Get(`to`)
	}
	fmt.Println(recordDefault)

	var response hamsoyamodels.ResponseHamsoyaRecords
	response = hamsoyamodels.GetHamsoyaRecordsCount(recordDefault, interval)
	response.TotalPage = int64(math.Ceil(float64(response.TotalPage) / float64(int64(100))))
	response.Page = helperfunc.MinOftoInt(int64(pageInt), response.TotalPage)
	response.URL = URL + PreURL
	fmt.Println(response.URL)
	hamsoyamodels.GetHamsoyaRecords(recordDefault, &response, interval, int64(pageInt) - 1)

	err = json.NewEncoder(writer).Encode(&response)
	if err != nil {
		log.Print(err)
	}
}

// Get Hamsoya Precheck
func (server *MainServer) GetHamsoyaPrecheckByIdHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(writer).Encode("invalid_id")
		if err != nil {
			log.Print(err)
			return
		}
		return
	}
	Precheck, err := hamsoyamodels.GetHamsoyaPreCheckById(int64(id))
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		err := json.NewEncoder(writer).Encode("server wrong")
		if err != nil {
			log.Print(err)
			return
		}
		return
	}
	err = json.NewEncoder(writer).Encode(&Precheck)
	if err != nil {
		log.Print("invalid_HamsoyaPrecheck.")
		return
	}
}

//Get Hamsoya Prechecks
func (server *MainServer) GetHamsoyaPrechecksHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	pageInt := 1
	rowsInt := 100
	page := request.URL.Query().Get(`page`)
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 1
	}
	rows := request.URL.Query().Get(`rows`)
	rowsInt, err = strconv.Atoi(rows)
	if err != nil {
		rowsInt = 100
	}
	var newHamsoyaPreCheck hamsoyamodels.HamsoyaPreCheck
	PreChecks := hamsoyamodels.GetHamsoyaPreChecks(newHamsoyaPreCheck, int64(rowsInt), int64(pageInt))

	if PreChecks.Error != nil {
		writer.WriteHeader(http.StatusNotFound)
		err := json.NewEncoder(writer).Encode(`mismatch_hamsoyaDocuments`)
		if err != nil {
			log.Println(err)
			return
		}
		return
	}
	err = json.NewEncoder(writer).Encode(PreChecks)
	if err != nil {
		log.Println(err)
		return
	}
	return
}

// Get Hamsoya Account
func (server *MainServer) GetHamsoyaAccountByIdHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(writer).Encode("invalid_id")
		if err != nil {
			log.Print(err)
			return
		}
		return
	}
	Account, err := hamsoyamodels.GetHamsoyaAccountById(int64(id))
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		err := json.NewEncoder(writer).Encode("server wrong")
		if err != nil {
			log.Print(err)
			return
		}
		return
	}
	err = json.NewEncoder(writer).Encode(&Account)
	if err != nil {
		log.Print("invalid_HamsoyaPrecheck.")
		return
	}
}

// Get Hamsoya Accounts
func (server *MainServer) GetHamsoyaAccountsHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	page := request.URL.Query().Get(`page`)
	PreURL := ``
	pageInt, err := strconv.Atoi(page)
	URL := `http://localhost:3000/hamsoya/accounts`
	if err != nil {
		pageInt = 1
		PreURL = `?page=1`
	} else {
		PreURL = fmt.Sprintf(`?page=%s`, pageInt)
	}
	var accountDefault hamsoyamodels.HamsoyaAccount
	Id, err := strconv.Atoi(request.URL.Query().Get(`Id`))
	if err == nil {
		accountDefault.Id = int64(Id)
		PreURL += fmt.Sprintf(`&Id=%s`, request.URL.Query().Get(`Id`))
	}
	ClientId, err := strconv.Atoi(request.URL.Query().Get(`ClientId`))
	if err == nil {
		accountDefault.ClientId = int64(ClientId)
		PreURL += fmt.Sprintf(`&ClientId=%s`, request.URL.Query().Get(`ClientId`))
	}
	Overdraft, err := strconv.ParseFloat(request.URL.Query().Get(`Overdraft`), 64)
	if err == nil {
		accountDefault.Overdraft = Overdraft
		PreURL += fmt.Sprintf(`&Overdraft=%s`, request.URL.Query().Get(`Overdraft`))
	}
	CurrencyId, err := strconv.Atoi(request.URL.Query().Get(`CurrencyId`))
	if err == nil {
		accountDefault.CurrencyId = int64(CurrencyId)
		PreURL += fmt.Sprintf(`&CurrencyId=%s`, request.URL.Query().Get(`CurrencyId`))
	}
	TypeId, err := strconv.Atoi(request.URL.Query().Get(`TypeId`))
	if err == nil {
		accountDefault.TypeId = int64(TypeId)
		PreURL += fmt.Sprintf(`&TypeId=%s`, request.URL.Query().Get(`TypeId`))
	}
	Saldo, err := strconv.ParseFloat(request.URL.Query().Get(`Saldo`), 64)
	if err == nil {
		accountDefault.Saldo = Saldo
		PreURL += fmt.Sprintf(`&Saldo=%s`, request.URL.Query().Get(`Saldo`))
	}
	AccNum := request.URL.Query().Get(`AccNum`)
	if AccNum != ``{
		accountDefault.AccNum = AccNum
		PreURL += fmt.Sprintf(`&AccNum=%s`, request.URL.Query().Get(`AccNum`))
	}
	IsActive, err := strconv.ParseBool(request.URL.Query().Get(`IsActive`))
	if err == nil {
		accountDefault.IsActive = IsActive
		PreURL += fmt.Sprintf(`&IsActive=%s`, request.URL.Query().Get(`IsActive`))
	}
	IsDefault, err := strconv.ParseBool(request.URL.Query().Get(`IsDefault`))
	if err == nil {
		accountDefault.IsDefault = IsDefault
		PreURL += fmt.Sprintf(`&IsDefault=%s`, request.URL.Query().Get(`IsDefault`))
	}

	var interval helperfunc.TimeInterval
	unix := time.Unix(0, 0)
	interval.From = unix.Format(time.RFC3339)
	unixTimeNow := time.Now()
	interval.To = unixTimeNow.Format(time.RFC3339)
	from, err := strconv.Atoi(request.URL.Query().Get(`from`))
	if err == nil {
		i := time.Unix(int64(from), 0)
		ans := i.Format(time.RFC3339)
		interval.From = ans
		PreURL += `&from=` + request.URL.Query().Get(`from`)
	}
	to, err := strconv.Atoi(request.URL.Query().Get(`to`))
	if err == nil {
		i := time.Unix(int64(to), 0)
		ans := i.Format(time.RFC3339)
		interval.To = ans
		PreURL += `&to=` + request.URL.Query().Get(`to`)
	}
	fmt.Println(accountDefault)

	var response hamsoyamodels.ResponseHamsoyaAccount
	response = hamsoyamodels.GetHamsoyaAccountsCount(accountDefault, interval)
	response.TotalPage = int64(math.Ceil(float64(response.TotalPage) / float64(int64(100))))
	response.Page = helperfunc.MinOftoInt(int64(pageInt), response.TotalPage)
	response.URL = URL + PreURL
	fmt.Println(response.URL)
	hamsoyamodels.GetHamsoyaAccounts(accountDefault, &response, interval, int64(pageInt) - 1)

	err = json.NewEncoder(writer).Encode(&response)
	if err != nil {
		log.Print(err)
	}
}

// Get Hamsoya Client
func (server *MainServer) GetHamsoyaClientByIdHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(writer).Encode("invalid_id")
		if err != nil {
			log.Print(err)
			return
		}
		return
	}
	Client, err := hamsoyamodels.GetHamsoyaClientById(int64(id))
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		err := json.NewEncoder(writer).Encode("server wrong")
		if err != nil {
			log.Print(err)
			return
		}
		return
	}
	err = json.NewEncoder(writer).Encode(&Client)
	if err != nil {
		log.Print("invalid_HamsoyaPrecheck.")
		return
	}
}

// Get Hamsoya Clients
func (server *MainServer) GetHamsoyaClientsHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	pageInt := 1
	page := request.URL.Query().Get(`page`)
	PreURL := ``
	URL := `http://127.0.0.1:8080/api/hamsoya/clients`
	pageInt, errPage := strconv.Atoi(page)
	if errPage != nil {
		pageInt = 1
		URL += `?page=1`
	}
	var clientDefault hamsoyamodels.HamsoyaClient
	IsActive, err := strconv.ParseBool(request.URL.Query().Get(`IsActive`))
	if err == nil {
		clientDefault.IsActive = IsActive
		PreURL += fmt.Sprintf(`&IsActive=%s`, request.URL.Query().Get(`IsActive`))
	}
	Identify, err := strconv.ParseBool(request.URL.Query().Get(`Identify`))
	if err == nil {
		clientDefault.Identify = Identify
		PreURL += fmt.Sprintf(`&Identify=%s`, request.URL.Query().Get(`Identify`))
	}
	PhoneNum := request.URL.Query().Get(`PhoneNum`)
	if PhoneNum != `` {
		clientDefault.PhoneNum = PhoneNum
		PreURL += `&PhoneNum=` + PhoneNum
	}
	Name := request.URL.Query().Get(`Name`)
	if Name != `` {
		clientDefault.Name = Name
		PreURL += `&Name=` + Name
	}
	id, err := strconv.Atoi(request.URL.Query().Get(`id`))
	if err == nil {
		clientDefault.Id = int64(id)
		PreURL += fmt.Sprintf(`&id=%s`, request.URL.Query().Get(`id`))
	}
	AgentId, err := strconv.Atoi(request.URL.Query().Get(`AgentId`))
	if err == nil {
		clientDefault.AgentId = int64(AgentId)
		PreURL += fmt.Sprintf(`&AgentId=%s`, request.URL.Query().Get(`AgentId`))
	}
	TypeId, err := strconv.Atoi(request.URL.Query().Get(`TypeId`))
	if err == nil {
		clientDefault.TypeId = int64(TypeId)
		PreURL += fmt.Sprintf(`&TypeId=%s`, request.URL.Query().Get(`TypeId`))
	}

	//Default-ные значение времени. типа с начало до конца
	var interval helperfunc.TimeInterval
	unix := time.Unix(0, 0)
	interval.From = unix.Format(time.RFC3339)
	unixTimeNow := time.Now()
	interval.To = unixTimeNow.Format(time.RFC3339)
	from, err := strconv.Atoi(request.URL.Query().Get(`from`))
	if err == nil {
		from /= 1000
		i := time.Unix(int64(from), 0)
		ans := i.Format(time.RFC3339)
		interval.From = ans
		PreURL += `&from=` + request.URL.Query().Get(`from`)
	}
	to, err := strconv.Atoi(request.URL.Query().Get(`to`))
	if err == nil {
		from /= 1000
		i := time.Unix(int64(to), 0)
		ans := i.Format(time.RFC3339)
		interval.To = ans
		PreURL += `&to=` + request.URL.Query().Get(`to`)
	}

	var response hamsoyamodels.ResponseHamsoyaClients
	response = hamsoyamodels.GetHamsoyaClientsCount(clientDefault, interval)
	response.TotalPage = int64(math.Ceil(float64(response.TotalPage) / float64(int64(100))))
	response.Page = helperfunc.MinOftoInt(int64(pageInt), response.TotalPage)
	if errPage == nil {
		URL += `?page=` + fmt.Sprintf("%d", response.Page)
	}
	response.URL = URL + PreURL
	fmt.Println(response.URL)
	hamsoyamodels.GetHamsoyaClients(clientDefault, &response, interval, response.Page-1)
	err = json.NewEncoder(writer).Encode(&response)
	if err != nil {
		log.Print(err)
	}
}
func (server *MainServer) GetHamsoyaViewTransesHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	fmt.Println("I am view Trans")
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	transaction := hamsoyamodels.HamsoyaViewTrans{}
	PreURL := ``
	Id, err := strconv.Atoi(request.URL.Query().Get(`Id`))
	if err == nil {
		transaction.ID = int64(Id)
		PreURL += `&Id=` + request.URL.Query().Get(`Id`)
	}
	VendorID, err := strconv.Atoi(request.URL.Query().Get(`VendorID`))
	if err == nil {
		transaction.VendorID = int64(VendorID)
		PreURL += `&VendorID=` + request.URL.Query().Get(`VendorID`)
	}

	ExtTransId := request.URL.Query().Get(`ExtTransId`)
	if ExtTransId != `` {
		transaction.ExtTransId = ExtTransId
		PreURL += `&ExtTransId=` + request.URL.Query().Get(`ExtTransId`)
	}
	RequestType := request.URL.Query().Get(`RequestType`)
	if ExtTransId != `` {
		transaction.RequestType = RequestType
		PreURL += `&RequestType=` + request.URL.Query().Get(`RequestType`)
	}
	PhoneNum := request.URL.Query().Get(`PhoneNum`)
	if ExtTransId != `` {
		transaction.PhoneNum = PhoneNum
		PreURL += `&PhoneNum=` + request.URL.Query().Get(`PhoneNum`)
	}
	ClientReceiver := request.URL.Query().Get(`ClientReceiver`)
	if ExtTransId != `` {
		transaction.ClientReceiver = ClientReceiver
		PreURL += `&ClientReceiver=` + request.URL.Query().Get(`ClientReceiver`)
	}
	Amount, err := strconv.ParseFloat(request.URL.Query().Get(`Amount`), 64)
	if err == nil {
		transaction.Amount = Amount
		PreURL += `&Amount=` + request.URL.Query().Get(`Amount`)
	}
	TotalAmount, err := strconv.ParseFloat(request.URL.Query().Get(`TotalAmount`), 64)
	if err == nil {
		transaction.TotalAmount = TotalAmount
		PreURL += `&TotalAmount=` + request.URL.Query().Get(`TotalAmount`)
	}
	ExternalFee, err := strconv.ParseFloat(request.URL.Query().Get(`ExternalFee`), 64)
	if err == nil {
		transaction.ExternalFee = ExternalFee
		PreURL += `&ExternalFee=` + request.URL.Query().Get(`ExternalFee`)
	}

	var interval helperfunc.TimeInterval
	unix := time.Unix(0, 0)
	interval.From = unix.Format(time.RFC3339)
	unixTimeNow := time.Now()
	interval.To = unixTimeNow.Format(time.RFC3339)
	from, err := strconv.Atoi(request.URL.Query().Get(`from`))
	if err == nil {
		i := time.Unix(int64(from), 0)
		ans := i.Format(time.RFC3339)
		interval.From = ans
		PreURL += `&from=` + request.URL.Query().Get(`from`)
	}
	to, err := strconv.Atoi(request.URL.Query().Get(`to`))
	if err == nil {
		i := time.Unix(int64(to), 0)
		ans := i.Format(time.RFC3339)
		interval.To = ans
		PreURL += `&to=` + request.URL.Query().Get(`to`)
	}

	var response hamsoyamodels.ResponseHamsoyaTranses
	page := request.URL.Query().Get(`page`)
	URL := `http://127.0.0.1:8080/api/hamsoya/viewtranses`
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 1
		URL += `?page=1`
	}
	response = hamsoyamodels.GetHamsoyaViewTransesCount(transaction, interval)
	response.TotalPage = int64(math.Ceil(float64(response.TotalPage) / float64(int64(100))))
	response.Page = helperfunc.MinOftoInt(int64(pageInt), response.TotalPage)
	if err == nil {
		URL += `?page=` + fmt.Sprintf("%d", response.Page)
	}
	URL += PreURL
	response.URL = URL
	fmt.Println(URL)
	hamsoyamodels.GetHamsoyaViewTranses(transaction, &response, interval, response.Page-1)
	err = json.NewEncoder(writer).Encode(&response)
	if err != nil {
		log.Print(err)
	}
}