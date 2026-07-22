package main

import (
	"log"

	"kakapo/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(gin.Logger())

	v1 := r.Group("/v1")
	router.SetupRouter(v1)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
