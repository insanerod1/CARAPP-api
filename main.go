package main

import (
	"CARAPP-api/internal/routes"
	"log"

	"CARAPP-api/internal/storage"

	"github.com/gin-gonic/gin"
)

func main() {
	comments, err := storage.LoadDeathComments("data/death_comments.json")
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	routes.SetupRoutes(router, comments)

	if err := router.Run("localhost:8080"); err != nil {
		log.Fatal(err)
	}
}
