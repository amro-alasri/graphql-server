package main

import (
	"os"

	"github.com/amro-alasri/graphQL-server/controllers"
	"github.com/amro-alasri/graphQL-server/middleware"
	"github.com/gin-gonic/gin"
)

const defaultPort = ":8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	server := gin.Default()

	server.Use(middleware.BasicAuth())
	server.GET("/", controllers.Playground())
	server.POST("/query", controllers.GraphQLHandler())
	server.Run(defaultPort)

}
