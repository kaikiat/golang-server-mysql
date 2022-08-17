package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Ping the server
// @Produce  json
// @Router /ping [get]
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ping server",
	})
}
