package main

import (
	"github.com/ahmedkhaeld/graphql-server/http"
	"github.com/ahmedkhaeld/graphql-server/middleware"
	"github.com/gin-gonic/gin"
	"os"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	server := gin.Default()
	server.Use(middleware.BasicAuth())
	server.POST("/query", http.GraphqlHandler())
	server.GET("/", http.PlaygroundHandler())
	server.Run(port)

}
