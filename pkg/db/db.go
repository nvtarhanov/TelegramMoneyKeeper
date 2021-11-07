package db

import (
	"fmt"

	"github.com/nvtarhanov/TelegramMoneyKeeper/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init(config config.DatabaseConfig) error {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Shanghai", config.Host, config.Username, config.Password, config.Dbname, config.Port, "disable")

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	DB.AutoMigrate()

	return nil
}

func GetDB() *gorm.DB {
	return DB
}
