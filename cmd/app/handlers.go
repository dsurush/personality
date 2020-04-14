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
)

func (server *MainServer) LoginHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	fmt.Println("login\n")
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

func (server *MainServer) GetClientInfoHandler(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	fmt.Println("I am find client By number phone")
	phone := param.ByName(`phone`)
	fmt.Println(phone)
	response := models.GetClientInfo(phone)
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
