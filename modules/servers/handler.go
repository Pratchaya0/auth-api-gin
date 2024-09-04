package servers

import (
	"net/http"

	_usersHttpService "github.com/Pratchaya0/auth-api-gin/modules/users/controllers"
	_userRepositories "github.com/Pratchaya0/auth-api-gin/modules/users/repositories"
	_useUseCases "github.com/Pratchaya0/auth-api-gin/modules/users/usecases"
	"github.com/gin-gonic/gin"
)

func (s *Server) MapHandlers() error {
	v1 := s.App.Group("/v1")

	userGroup := v1.Group("/users")
	userRepository := _userRepositories.NewUserRepository(s.Db)
	usersUseCase := _useUseCases.NewUsersUseCase(userRepository)
	_usersHttpService.NewUsersController(userGroup, usersUseCase)

	s.App.Use(func(c *gin.Context) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":      "InternalServerError",
			"status_code": http.StatusInternalServerError,
			"message":     "error, end point not found",
			"result":      nil,
		})
		return
	})

	return nil
}
