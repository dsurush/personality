package db

import (
	"MF/models"
	"MF/settings"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	log "github.com/sirupsen/logrus"
)

var (
	postgresDbCon *gorm.DB
)

// InitDb postgresMegafonDB init
func InitDb() *gorm.DB {
	settings.AppSettings = settings.ReadSettings()

	postgresMegafondbParams := settings.AppSettings.PostgresMegafonDbParams
	fmt.Println(postgresMegafondbParams)
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s database=%s sslmode=disable",
		postgresMegafondbParams.Server, postgresMegafondbParams.Port,
		postgresMegafondbParams.User, postgresMegafondbParams.Password,
		postgresMegafondbParams.Database)

	// Opening connection
	db, err := gorm.Open("postgres", connString)

	db.LogMode(true)
	// Error
	if err != nil {
		//fmt.Println("error ошибка - ", err)
		log.Warn("Database init error", err)
		panic(err)
	}
	db.AutoMigrate(&models.ClientInfo{}, &models.Merchant{}, &models.TablePreeCheck{}, &models.TableTransaction{}, &models.RefundedCardTransactions{}, &models.VendorListReqRawXML{},
	&models.PaymentReqRawXML{}, models.ResponseLog{} ,&models.User{}, &models.Role{}, &models.Vendor{})
	return db
}

// InitPostgresDatabase initializes database
func InitPostgresDatabase() {
	postgresDbCon = InitDb()
}

// GetPostgresDb is func for create one global connection
func GetPostgresDb() *gorm.DB {
	return postgresDbCon
}
