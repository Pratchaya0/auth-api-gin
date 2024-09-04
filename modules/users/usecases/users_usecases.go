package usecases

import (
	"fmt"

	"github.com/Pratchaya0/auth-api-gin/modules/models"
	"github.com/Pratchaya0/auth-api-gin/modules/users/interfaces"
	"golang.org/x/crypto/bcrypt"
)

// Usecases
// .
// มีหน้าที่ รับมือกับ Logic ต่างๆ ก่อนที่จะส่งข้อมูลเข้าออก Database เช่น Search, Sort, Hash, …
// .
// อย่างถ้าเราต้องการจะเขียน Usecase สำหรับการสมัครสมาชิกของ User ก็ควรจะมี Logic ในการ Hash Password ให้กับ user สักหน่อยเพื่อความปลอดภัยก่อนที่จะส่งเข้า Database

type usersUseCase struct {
	usersRepo interfaces.UsersRepository
}

func NewUsersUseCase(usersRepo interfaces.UsersRepository) interfaces.UsersUseCase {
	return &usersUseCase{usersRepo: usersRepo}
}

func (u *usersUseCase) Register(req *models.UserRegisterReq) (*models.UserRegisterRes, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	req.Password = string(hashed)

	user, err := u.usersRepo.Register(req)
	if err != nil {
		return nil, err
	}

	return user, nil
}
