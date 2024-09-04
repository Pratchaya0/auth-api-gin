package repositories

// Repositories
// .
// มีหน้าที่ ในการรับส่ง Entities เข้าออกจาก Database หรือพูดง่ายๆ ก็คือมีหน้าที่ Query ข้อมูลจาก Database

import (
	"github.com/Pratchaya0/auth-api-gin/modules/models"
	"github.com/Pratchaya0/auth-api-gin/modules/users/interfaces"
	"gorm.io/gorm"
)

type usersRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) interfaces.UsersRepository {
	return &usersRepo{db: db}
}

func (r *usersRepo) Register(req *models.UserRegisterReq) (*models.UserRegisterRes, error) {
	var res models.UserRegisterRes

	err := r.db.Find(&res).Error

	return &res, err
}
