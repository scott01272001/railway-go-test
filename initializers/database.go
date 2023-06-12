package initializers

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectToDB() {
	value, err := GetConfigValue("DB_URL")
	if err != nil {
		panic(err)
	}
	DB, err = gorm.Open(postgres.Open(value), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to DB")
	}
}
