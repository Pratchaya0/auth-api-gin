package interfaces

// Interface
// .
// มีหน้าที่ กำหนดโครงสร้างของข้อมูล เพื่อใช้สื่อสารระหว่าง Layer
// .
// เช่น ผมต้องการเขียน Service API สำหรับสมาชิก Interface ก็ควรจะประกอบไปด้วย interface ของ UseCase และ Repository สำหรับการสมัครสมาชิก

import "github.com/Pratchaya0/auth-api-gin/modules/models"

// ตัวกลางระหว่าง Layer Controller กับ UseCase
type UsersUseCase interface {
	Register(req *models.UserRegisterReq) (*models.UserRegisterRes, error)
}

// ตัวกลางระหว่าง Layer UseCase กับ Repository
type UsersRepository interface {
	Register(req *models.UserRegisterReq) (*models.UserRegisterRes, error)
}
