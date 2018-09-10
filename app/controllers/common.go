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

// Docs controller
func Docs(c *gin.Context) {
	c.Redirect(
		http.StatusMovedPermanently,
		"https://hackmd.io/_j-QIksAS46DBwtwCZ3vuw?view",
	)
}
