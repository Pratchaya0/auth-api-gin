package repositories

import (
	"github.com/Pratchaya0/auth-api-gin/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetListUser() ([]models.User, error) {
	var users []models.User

	err := r.db.Preload("Provider").Preload("Authorities").Find(&users).Error

	return users, err
}

func (r *UserRepository) GetUserById(id string) (models.User, error) {
	var user models.User

	err := r.db.Where("id = ?", user).First(&user).Error

	return user, err
}

func (r *UserRepository) GetUserByUserName(username string) (models.User, error) {
	var user models.User

	err := r.db.Where(&models.User{UserName: username}).First(&user).Error

	return user, err
}

func (r *UserRepository) GetUserByProviderId(providerID string) ([]models.User, error) {
	var user []models.User

	err := r.db.Joins("Provider").Where(&models.Provider{ProviderID: providerID}).Find(&user).Error

	return user, err
}

func (r *UserRepository) CreateUser(user models.User) error {
	err := r.db.Create(&user).Error

	return err
}

func (r *UserRepository) UpdateUser(user models.User) error {
	err := r.db.Save(&user).Error

	return err
}

func (r *UserRepository) DeleteUser(user models.User) error {
	err := r.db.Delete(&user).Error

	return err
}
