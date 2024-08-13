package usecases

import (
	"github.com/Pratchaya0/auth-api-gin/models"
	repositories "github.com/Pratchaya0/auth-api-gin/repositories/UserRepository"
)

type UserUseCase struct {
	userInterface repositories.UserRepository
}

func NewUserUseCase(userInterface repositories.UserRepository) *UserUseCase {
	return &UserUseCase{userInterface: userInterface}
}

func (u *UserUseCase) GetListUser() ([]models.User, error) {
	return u.userInterface.GetListUser()
}

func (u *UserUseCase) GetUserByUserName(username string) (models.User, error) {
	return u.userInterface.GetUserByUserName(username)
}

// func (u *UserUseCase) GetUserByProviderId(providerID string) (models.User, error) {
// 	return u.userInterface.GetUserByProviderId(providerID)
// }

func (u *UserUseCase) CreateUser(user models.User) error {
	return u.userInterface.CreateUser(user)
}

func (u *UserUseCase) UpdateUser(user models.User) error {
	return u.userInterface.UpdateUser(user)
}

// func (u *UserUseCase) DeleteUser(id string) error {
// 	return u.userInterface.DeleteUser(id)
// }
