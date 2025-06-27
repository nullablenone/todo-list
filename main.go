package main

import (
	"todo-list/config"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadENV()
	config.ConnectDB()

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "success",
		})
	})
	router.Run()
}
