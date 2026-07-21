package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		fmt.Println(" entrou aq")
		fmt.Printf("O vet vai pegar isso: %d", "uma string")

		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	router.Run(":8080")
}
