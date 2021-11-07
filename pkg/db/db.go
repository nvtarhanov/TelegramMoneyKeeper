package db

import (
	"fmt"

	"github.com/nvtarhanov/TelegramMoneyKeeper/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func Init(config config.DatabaseConfig) error {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Shanghai", config.Host, config.Username, config.Password, config.Dbname, config.Port, "disable")

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	DB.AutoMigrate(&Product{})

	// Create
	DB.Create(&Product{Code: "D42", Price: 100})

	return nil
}
