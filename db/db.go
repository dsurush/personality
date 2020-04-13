package db

import (
	"MF/settings"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	log "github.com/sirupsen/logrus"
)

var postgresDbCon *gorm.DB

// InitDb postgresMegafonDB init
func init() {
	fmt.Println("DB INIT")
	var err error
	settings.AppSettings = settings.ReadSettings()

	postgresMegafondbParams := settings.AppSettings.PostgresMegafonDbParams
//	fmt.Println(postgresMegafondbParams)
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s database=%s sslmode=disable",
		postgresMegafondbParams.Server, postgresMegafondbParams.Port,
		postgresMegafondbParams.User, postgresMegafondbParams.Password,
		postgresMegafondbParams.Database)

	// Opening connection
	postgresDbCon, err = gorm.Open("postgres", connString)

	postgresDbCon.LogMode(true)
	// Error
	if err != nil {
		//fmt.Println("error ошибка - ", err)
		log.Warn("Database init error", err.Error())
		panic(err)
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
