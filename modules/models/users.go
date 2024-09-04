package models

// Entities หรือ Models
// มีหน้าที่ กำหนดโครงสร้างของข้อมูล แลพ Table Diagrams
// เช่น ผมต้องการเขียน Service API สำหรับสมาชิก Entities ก็ควรจะประกอบไปด้วย Struct ของ Request และ Response สำหรับการสมัครสมาชิก

type UserRegisterReq struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

type UserRegisterRes struct {
	Id       uint64 `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
}
