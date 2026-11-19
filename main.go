package main

import (
	"CARAPP-api/internal/routes"
	"log"

	"CARAPP-api/internal/storage"

	"github.com/gin-gonic/gin"
)

func main() {
	death_comments, err := storage.LoadDeathComments("data/death_comments.json")
	assistant_comments, err := storage.LoadAssistantComments("data/driving_assistant.json")

	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	routes.SetupRoutes(router, death_comments, assistant_comments)

	if err := router.Run("localhost:8080"); err != nil {
		log.Fatal(err)
	}
}
