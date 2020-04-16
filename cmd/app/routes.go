package app

import (
	"MF/middleware/authorized"
	"MF/middleware/jwt"
	"MF/middleware/logger"
	"MF/token"
	"fmt"
	"net/http"
	"reflect"
)

func (server *MainServer) InitRoutes() {
	//bytes, _ := bcrypt.GenerateFromPassword([]byte(`alisher`), bcrypt.DefaultCost)
	//fmt.Println(string(bytes))
	//	test()
	fmt.Println("Init routes")
	//server.router.GET("/", server.LoginHandler1)
	server.router.POST("/api/login", logger.Logger(`Create Token for user: "`)(server.LoginHandler))
	server.router.GET("/api/client/:phone", logger.Logger(`Get client: `)(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))((authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetClientInfoHandler)))))
	server.router.GET("/api/clients", logger.Logger(`Get all clients: `)(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))((authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetClientsInfoHandler)))))
	server.router.GET("/api/vendors", logger.Logger(`Get all vendors`)(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))((authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetVendorCategoryHandler)))))
	//TODO:
	panic(http.ListenAndServe("127.0.0.1:8080", server))
}

func test()  {
//	user, _ := models.FindUserByLogin("surush")
//	data := string(time.Now())
//	file := ioutil.WriteFile("account.txt", data, 0666)
//	if file != nil {
	//	log.Fatalf("Xuyovo vsyo")
	//}
//		log.SetOutput(fmt.Println(`sas`))
//	s := "13.213.321:132"
//	suffix := strings.Split(s, `:`)
//	fmt.Println(suffix[0])
//	user, err := models.FindUserByLogin(`alisher`)
//	if err != nil {
//		log.Fatalf("Я не нашел Алишера")
	//}
	//fmt.Println(user)
	//file, err := os.OpenFile(`test`, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	//if err != nil {
	//	log.Panic("Failed to log to file", err)
	//	panic(err)
	//}
	//defer func() {
	//	err2 := file.Close()
	//	if err2 != nil {
	//		fmt.Println("ошибка закрытие файла")
	//	}
	//}()
	//n, err := file.Write([]byte(`Loging`))
	//fmt.Println(n)
}