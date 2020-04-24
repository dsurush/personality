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

func (server *MainServer) GetClientsInfoHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	fmt.Println("I am get all clients")
	response := server.userSvc.GetClientsInfo()
	err := json.NewEncoder(writer).Encode(&response)
	if err != nil {
		log.Print(err)
	}
}

func (server *MainServer) GetClientsInfoByFilterHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	fmt.Println("I am get all clients")
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
	if err != nil {
		clientDefault.IsIdentified = IsIdentified
	}
	IsBlackList, err := strconv.ParseBool(request.URL.Query().Get(`IsBlackList`))
	if err != nil {
		clientDefault.IsBlackList = IsBlackList
	}
	SendToCft, err := strconv.ParseBool(request.URL.Query().Get(`SendToCft`))
	if err != nil {
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
/// UnUse
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
	fmt.Println("I am find client By number phone")
//	page := param.ByName("page")
//	rows := param.ByName(`rows`)
	page := request.URL.Query().Get(`page`)
	rows := request.URL.Query().Get(`rows`)
	//fmt.Println(get1, get2, "eeeeee")
	pageInt, _ := strconv.Atoi(page)
	rowsInt, _ := strconv.Atoi(rows)
	var vendor models.Vendor
	response := vendor.FindAllVendorsByPageSize(pageInt, rowsInt)
	err := json.NewEncoder(writer).Encode(&response)
	if err != nil {
		log.Print(err)
	}
}

func (server *MainServer) GetViewTransactionsHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	fmt.Println("I am get all clients")
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
	response := models.GetViewTransactions(int64(rowsInt), int64(pageInt - 1))
	err = json.NewEncoder(writer).Encode(&response)
	if err != nil {
		log.Print(err)
	}
}