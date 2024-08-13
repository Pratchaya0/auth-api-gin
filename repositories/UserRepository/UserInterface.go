package repositories

import "github.com/Pratchaya0/auth-api-gin/models"

type UserInterface interface {
	GetListUser() ([]models.User, error)
	GetUserByUserName(username string) (models.User, error)
	GetUserByProviderId(providerID string) (models.User, error)
	CreateUser(user models.User) error
	UpdateUser(user models.User) error
	DeleteUser(id string) error
}
