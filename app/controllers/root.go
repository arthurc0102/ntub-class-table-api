package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Root controller
func Root(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "NTUB Class Table API",
	})
}
