package main

import (
	"fmt"
	"net/http"

	"github.com/Pratchaya0/auth-api-gin/routes"
	"github.com/gin-gonic/gin"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	router := gin.Default()

	routes.Auth2RouteSetup(router)

	router.Run(":8080")
}
