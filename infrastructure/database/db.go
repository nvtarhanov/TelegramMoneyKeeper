package database

import (
	"fmt"

	"github.com/nvtarhanov/TelegramMoneyKeeper/infrastructure/config"
	"github.com/nvtarhanov/TelegramMoneyKeeper/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//Need to refactor this
//https://techinscribed.com/different-approaches-to-pass-database-connection-into-controllers-in-golang/

func Init(config config.DatabaseConfig) (*gorm.DB, error) {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Shanghai", config.Host, config.Username, config.Password, config.Dbname, config.Port, "disable")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&model.Account{}, &model.Entrie{}, &model.Transaction{}, &model.State{})

	return db, nil
}

//Dont forget to close db connection after server stopped (defer)
func CloseDbConnection() error {

	return nil
}
