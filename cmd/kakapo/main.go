package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		fmt.Println(" entrou aq")

		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	router.Run(":8080")
}
