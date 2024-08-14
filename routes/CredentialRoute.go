package routes

import (
	controllers "github.com/Pratchaya0/auth-api-gin/controllers/auth"
	"github.com/Pratchaya0/auth-api-gin/database"
	repositories "github.com/Pratchaya0/auth-api-gin/repositories/UserRepository"
	"github.com/Pratchaya0/auth-api-gin/usecases"
	"github.com/gin-gonic/gin"
)

func CredentialRouteSetup(engin *gin.Engine) {
	userRepo := repositories.NewGormUserRepository(database.DB())
	userUseCase := usecases.NewUserUseCase(*userRepo)
	credentialController := controllers.NewCredentialController(userUseCase)

	credentialRoute := engin.Group("/credential")
	{
		credentialRoute.GET("/currentUser", credentialController.CurrentUser)
		credentialRoute.POST("/register", credentialController.Register)
		credentialRoute.POST("/login", credentialController.Login)
		credentialRoute.POST("/logout", credentialController.LogOut)
	}
}
