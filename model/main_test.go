package model

import (
	"log"
	"os"
	"testing"

	"github.com/nvtarhanov/TelegramMoneyKeeper/pkg/db"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func TestMain(m *testing.M) {

	dsn := "host=localhost user=root password=password dbname=telegramdb port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Cannot connect to database")
	}

	DB = database

	db.SetDB(DB)

	os.Exit(m.Run())

}
