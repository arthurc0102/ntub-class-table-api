package main

import (
	"github.com/arthurc0102/ntub-class-table-api/app/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("", controllers.Root)
	server.GET("/docs", controllers.Docs)
	server.GET("/personal/:id", controllers.PersonalClassTable)

	server.Run()
}
