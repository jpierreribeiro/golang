package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "POST opening",
	})
}

func ShowOpeningHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "GET opening",
	})
}

func UpdateOpeningHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "PUT opening",
	})
}

func DeleteOpeningHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "DELETE opening",
	})
}

func ListOpeningHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "GET openings",
	})
}
