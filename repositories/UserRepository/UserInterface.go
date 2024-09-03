package repositories

import "github.com/Pratchaya0/auth-api-gin/models"

type UserInterface interface {
	GetListUser() ([]models.User, error)
	GetUserById(id string) (models.User, error)
	GetUserByUserName(username string) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
	GetUserByProviderId(providerID string) ([]models.User, error)
	CreateUser(user models.User) error
	UpdateUser(user models.User) error
	DeleteUser(user models.User) error
}
