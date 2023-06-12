package main

import (
	"gitgub.com/go-gin-gorm/initializers"
	"gitgub.com/go-gin-gorm/model"
	"log"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	err := initializers.DB.AutoMigrate(&model.Post{})
	if err != nil {
		log.Fatal(err)
	}
}
