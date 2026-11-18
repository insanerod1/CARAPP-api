package routes

import (
	"CARAPP-api/internal/handlers"
	"CARAPP-api/internal/models"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, comments []models.DeathComments) {
	router.GET("/comments/death", handlers.NewDeathCommentHandler(comments))
}
