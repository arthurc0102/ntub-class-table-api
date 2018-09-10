package main

import (
	"github.com/arthurc0102/ntub-class-table-api/app/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Init server
	server := gin.Default()

	// Middleware
	server.Use(cors.Default())

	// Routing
	server.GET("", controllers.Root)
	server.GET("/docs", controllers.Docs)
	server.GET("/personal/:id", controllers.PersonalClassTable)

	// Start server
	server.Run()
}
