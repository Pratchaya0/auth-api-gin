package routes

import "github.com/gin-gonic/gin"

func PublicRoute(engin *gin.Engine) {
	// Serve static files from the "public" directory
	engin.Static("/static", "./public")

	// Serve index.html when visiting the root URL
	engin.GET("/", func(c *gin.Context) {
		c.File("./public/index.html")
	})

	engin.GET("/login", func(c *gin.Context) {
		c.File("./public/login.html")
	})

	engin.GET("/register", func(c *gin.Context) {
		c.File("./public/register.html")
	})
}
