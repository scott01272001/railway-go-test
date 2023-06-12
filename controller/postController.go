package controller

import (
	"gitgub.com/go-gin-gorm/initializers"
	"gitgub.com/go-gin-gorm/model"
	"github.com/gin-gonic/gin"
)

func PostCreate(c *gin.Context) {
	var content model.Post
	err := c.ShouldBindJSON(&content)
	if err != nil {
		c.Status(400)
		c.JSON(400, gin.H{
			"error": err,
		})
	}

	post := model.Post{Name: content.Name, Title: content.Title, Body: content.Body}
	result := initializers.DB.Create(&post)
	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}
