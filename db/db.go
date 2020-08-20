package db

import (
	"MF/settings"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	log "github.com/sirupsen/logrus"
	"time"
)

var postgresDbCon *gorm.DB
var postgresHamsoyaDbCon *gorm.DB

// InitDb postgresMegafonDB init
func init() {
	fmt.Println("DB INIT")
	var err error
	settings.AppSettings = settings.ReadSettings("./settings-dev.json")
	//settings.HamsoyaSettings = settings.ReadSettings("./settings-hamsoya.json")

	//	postgresHamsoyadbParams := settings.HamsoyaSettings.PostgresMegafonDbParams

	postgresMegafondbParams := settings.AppSettings.PostgresMegafonDbParams
	//test
	postgresHamsoyadb := settings.AppSettings.PostgresHamsoyaDbParams

	//fmt.Printf("DATABASE = %s\n", postgresHamsoyadb.Database)

	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s database=%s sslmode=disable",
		postgresMegafondbParams.Server, postgresMegafondbParams.Port,
		postgresMegafondbParams.User, postgresMegafondbParams.Password,
		postgresMegafondbParams.Database)

	connHamsoyaString := fmt.Sprintf("host=%s port=%d user=%s password=%s database=%s sslmode=disable",
		postgresHamsoyadb.Server, postgresHamsoyadb.Port,
		postgresHamsoyadb.User, postgresHamsoyadb.Password,
		postgresHamsoyadb.Database)

	//connHamsoyaString := fmt.Sprintf("host=%s port=%d user=%s password=%s database=%s sslmode=disable",
	//	postgresMegafondbParams.Server, postgresMegafondbParams.Port,
	//	postgresMegafondbParams.User, postgresMegafondbParams.Password,
	//	"hamsoya")

	fmt.Printf("%s\n%s\n", connHamsoyaString, connString)
	// Opening con
	//
	//
	// nection
	//	fmt.Println("I am here = ", settings.AppSettings.LinkForCancelTransaction)
	postgresDbCon, err = gorm.Open("postgres", connString)

	postgresDbCon.LogMode(true)
	// Error
	if err != nil {
		fmt.Println("error ошибка  obj - ", err)
		fmt.Println("error ошибка - ", err.Error())
		time.Sleep(time.Second * 10)
		log.Warn("Database init error", err.Error())
		panic(err)
	}
	postgresHamsoyaDbCon, err = gorm.Open("postgres", connHamsoyaString)
	postgresHamsoyaDbCon.LogMode(true)

	if err != nil {
		//fmt.Println("error ошибка - ", err)
		//log.Warn("Database init error", err.Error())
		//time.Sleep(time.Second * 10)
		log.Fatalf("Can't connect to Hamsoya DB %e", err)
		//panic(err)
	}
	/*postgresDbCon.AutoMigrate(&models.ClientInfo{}, &models.Merchant{}, &models.TablePreeCheck{},
	&models.TableTransaction{}, &models.RefundedCardTransactions{}, &models.VendorListReqRawXML{},
	&models.PaymentReqRawXML{}, models.ResponseLog{} , &models.Vendor{}, &models.User{}, &models.Role{})
	*/
}

// GetPostgresDb is func for create one global connection
func GetPostgresDb() *gorm.DB {
	return postgresDbCon
}
func GetHamsoyaPostgresDb() *gorm.DB {
	return postgresHamsoyaDbCon
}
