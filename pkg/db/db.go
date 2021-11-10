package db

import (
	"fmt"

	"github.com/nvtarhanov/TelegramMoneyKeeper/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//Need to refactor this
//https://techinscribed.com/different-approaches-to-pass-database-connection-into-controllers-in-golang/

var DB *gorm.DB

func Init(config config.DatabaseConfig) error {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Shanghai", config.Host, config.Username, config.Password, config.Dbname, config.Port, "disable")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	DB = db

	return nil
}

func GetDB() *gorm.DB {
	return DB
}

func SetDB(db *gorm.DB) {
	DB = db
}

//Dont forget to close db connection after server stopped (defer)
func CloseDbConnection() error {

	return nil
}