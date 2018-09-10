package controllers

import (
	"net/http"
	"strconv"

	"github.com/arthurc0102/ntub-class-table-api/app/services"
	"github.com/gin-gonic/gin"
)

// PersonalClassTable controller
func PersonalClassTable(c *gin.Context) {
	id := c.Params.ByName("id")
	if _, err := strconv.Atoi(id); err != nil || id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Student ID not valid",
		})
		return
	}

	classTable, errorList := services.PersonalClassTable(id)
	if len(errorList) != 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Can't get classtable for this student ID, please try again later",
			"errors":  errorList,
		})
		return
	}

	c.JSON(http.StatusOK, classTable)
}
