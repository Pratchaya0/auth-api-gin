package main

import (
	"github.com/Pratchaya0/auth-api-gin/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Serve static files from the "public" directory
	router.Static("/static", "./public")

	// Serve index.html when visiting the root URL
	router.GET("/", func(c *gin.Context) {
		c.File("./public/index.html")
	})

	routes.CredentialRouteSetup(router)
	routes.Auth2RouteSetup(router)

	router.Run(":8080")
}
