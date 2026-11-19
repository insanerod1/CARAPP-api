package handlers

import (
	"math/rand/v2"
	"net/http"

	"CARAPP-api/internal/models"

	"github.com/gin-gonic/gin"
)

func NewDeathCommentHandler(comments []models.DeathComments) gin.HandlerFunc {
	return func(c *gin.Context) {
		obstacle := c.Query("obstacle")
		if obstacle == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "An obstacle is required",
			})
			return
		}

		for i := 0; i < len(comments); i++ {
			if comments[i].Obstacle == obstacle {
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
			"error": "Unknown obstacle",
		})
	}
}
