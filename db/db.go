package db

import (
	"MF/settings"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	log "github.com/sirupsen/logrus"
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

	fmt.Println(connHamsoyaString, '\n', connString, '\n')
	// Opening con
	//
	//
	// nection
	postgresDbCon, err = gorm.Open("postgres", connString)

	postgresDbCon.LogMode(true)
	// Error
	if err != nil {
		//fmt.Println("error ошибка - ", err)
		log.Warn("Database init error", err.Error())
		panic(err)
	}
	postgresHamsoyaDbCon, err = gorm.Open("postgres", connHamsoyaString)
	postgresHamsoyaDbCon.LogMode(true)
	if err != nil {
		//fmt.Println("error ошибка - ", err)
		//log.Warn("Database init error", err.Error())
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