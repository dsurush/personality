package main

import (
	"MF/cmd/app"
	_ "MF/db"
	"MF/models"
	"MF/token"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	router := httprouter.New()
	tokenSvc := token.NewTokenSvc([]byte(`surush`))
	usersvc := models.NewUsersvc()
	router.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header.Get("Access-Control-Request-Method") != "" {
			// Set CORS headers
			w.Header().Set("Content-Type", "application/json, text/html")
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, RefreshToken")
			w.Header().Set("Accept", "*/*")
			w.WriteHeader(http.StatusOK)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	})

	server := app.NewMainServer(router, tokenSvc, usersvc)
	server.Start()
}
