package main

import (
	"gitgub.com/go-gin-gorm/controller"
	"gitgub.com/go-gin-gorm/initializers"
	"github.com/gin-gonic/gin"
	"log"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	router := gin.Default()

	router.POST("/posts", controller.PostCreate)

	port, err := initializers.GetConfigValue("PORT")
	if err != nil {
		log.Fatal(err)
	}
	router.Run(":" + port)
}
