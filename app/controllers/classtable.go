package controllers

import (
	"net/http"

	"github.com/arthurc0102/ntub-class-table-api/app/services"
	"github.com/gin-gonic/gin"
)

// PersonalClassTable controller
func PersonalClassTable(c *gin.Context) {
	id := c.Params.ByName("id")

	classTable, classTime, errorList := services.PersonalClassTable(id)
	if len(errorList) != 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Can't get classtable for this student ID, please try again later",
			"errors":  errorList,
		})
		return
	}

	total := 0
	for _, classList := range classTable {
		for _, class := range classList {
			total += len(class)
		}
	}

	if total == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "No data found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"time":  classTime,
		"class": classTable,
	})
}
