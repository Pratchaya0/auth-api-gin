package main

import (
	"github.com/Pratchaya0/auth-api-gin/database"
	"github.com/Pratchaya0/auth-api-gin/middlewares"
	"github.com/Pratchaya0/auth-api-gin/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	database.SetupDatabase()

	middlewares.CORSMiddleware(router)

	routes.PublicRoute(router)
	routes.CredentialRouteSetup(router)
	routes.Auth2RouteSetup(router)

	router.Run(":8080")
}
