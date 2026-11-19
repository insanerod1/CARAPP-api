package handlers

import (
	"math/rand/v2"
	"net/http"

	"CARAPP-api/internal/models"

	"github.com/gin-gonic/gin"
)

func NewDrivingAssistantHandler(comments []models.AssistantComments) gin.HandlerFunc {
	return func(c *gin.Context) {
		event := c.Query("event")
		if event == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Please specify an event",
			})
			return
		}

		for i := 0; i < len(comments); i++ {
			if comments[i].Event == event {
				options := comments[i].Comments

				if len(options) == 0 {
					c.JSON(http.StatusInternalServerError, gin.H{
						"error": "No comments available for this obstacle",
					})
					return
				}

				index := rand.IntN(len(options))

				c.JSON(http.StatusOK, gin.H{
					"comment": options[index],
				})
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{
			"error": "Unknown event",
		})
	}
}
