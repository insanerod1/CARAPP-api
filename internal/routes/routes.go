package routes

import (
	"CARAPP-api/internal/handlers"
	"CARAPP-api/internal/models"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, deathcomments []models.DeathComments, assistantcomments []models.AssistantComments) {
	router.GET("/comments/death", handlers.NewDeathCommentHandler(deathcomments))
	router.GET("/comments/assistant", handlers.NewDrivingAssistantHandler(assistantcomments))
}
