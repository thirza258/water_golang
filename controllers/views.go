package controllers

import (
	"github.com/gin-gonic/gin"
	"water/models"
	"water/services"
)

func GetQuality(c *gin.Context) {
	var quality []models.Quality
	err := services.GetAllQuality(&quality)
	if err != nil {
		c.JSON(404, gin.H{"status": 404, "message": "Not Found", "response": ""})
		return
	}
	c.JSON(200, quality)
}
