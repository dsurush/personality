package app

import (
	"MF/logger"
	"fmt"
	"log"
	"net/http"
	"os"
)

func (server *MainServer) InitRoutes() {
	//router := httprouter.New()
	//bytes, _ := bcrypt.GenerateFromPassword([]byte(`surush`), bcrypt.DefaultCost)
	//fmt.Println(string(bytes))
		//user, err := models.FindUserByLogin("surush")
		//if err != nil {
		//	log.Fatalf("НЕТ ТАКОГО ПОЛЬЗОВАТЕЛЯ %e\n", err)
		//}
		//fmt.Println(user)
		test()
	fmt.Println("Init routes")
	server.router.POST("/api/login", logger.Logger(`Create Token for user: "`)(server.CreateTokenHandler))
	server.router.GET("/api/client/:phone", logger.Logger(`Get client: `)(server.GetClientInfoHandler))
	server.router.GET("/api/clients", logger.Logger(`Get all clients: `)(server.GetClientsInfoHandler))
	//TODO: GET ALL TRANSACTION
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
	file, err := os.OpenFile(`test`, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Panic("Failed to log to file", err)
		panic(err)
	}
	defer func() {
		err2 := file.Close()
		if err2 != nil {
			fmt.Println("ошибка закрытие файла")
		}
	}()
	n, err := file.Write([]byte(`Loging`))
	fmt.Println(n)
}