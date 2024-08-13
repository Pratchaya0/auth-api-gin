package routes

import (
	controllers "github.com/Pratchaya0/auth-api-gin/controllers/auth"
	"github.com/Pratchaya0/auth-api-gin/database"
	repositories "github.com/Pratchaya0/auth-api-gin/repositories/UserRepository"
	"github.com/Pratchaya0/auth-api-gin/usecases"
	"github.com/gin-gonic/gin"
)

func Auth2RouteSetup(engin *gin.Engine) {
	userRepo := repositories.NewGormUserRepository(database.DB())
	userUseCase := usecases.NewUserUseCase(*userRepo)
	auth2Controller := controllers.NewAuth2Controller(userUseCase)

	auth2Route := engin.Group("/auth")
	{
		auth2Route.GET("/", auth2Controller.OAuthIndex)
		auth2Route.GET("/:provider", auth2Controller.OAuthStart)
		auth2Route.GET("/logout/:provider", auth2Controller.OAuthLogout)
		auth2Route.GET("/:provider/callback", auth2Controller.OAuthCallback)
	}
}
