package controllers

import (
	"net/http"

	"github.com/Pratchaya0/auth-api-gin/modules/models"
	"github.com/Pratchaya0/auth-api-gin/modules/users/interfaces"
	"github.com/gin-gonic/gin"
)

// Controllers
// .
// มีหน้าที่ในการรับส่ง Context จาก HTTP Request หรือพูดง่ายๆก็คือ รับส่งข้อมูลหรือบริบทต่างๆที่ถูกยิงมากจาก API ที่ client ทำการยิงมา
// .
// โดยในการสร้าง Controller สำหรับการสมัครสมาชิก ก็ควรจะมี Endpoint และ ใช้เป็น Method POST เพื่อรับ Body ที่ Request มาจาก Client
// เพื่อทำการสมัครสมาชิก และเราก็อาจจะมีการตรวจสอบข้อมูลที่ส่งเข้ามาเล็กน้อย ก่อนส่งต่อเข้า Usecase ให้ไปทำงานในลำดับถัดไป

type usersController struct {
	usersUseCase interfaces.UsersUseCase
}

func NewUsersController(r *gin.RouterGroup, usersUseCase interfaces.UsersUseCase) {
	controllers := &usersController{usersUseCase: usersUseCase}

	// Router
	r.POST("/register", controllers.Register)
}

func (h *usersController) Register(c *gin.Context) {
	req := new(models.UserRegisterReq)
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":      "BadRequest",
			"status_code": http.StatusBadRequest,
			"message":     err.Error(),
			"result":      nil,
		})
		return
	}

	res, err := h.usersUseCase.Register(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":      "InternalServerError",
			"status_code": http.StatusInternalServerError,
			"message":     err.Error(),
			"result":      nil,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status":      "OK",
		"status_code": http.StatusOK,
		"message":     nil,
		"result":      res,
	})
}
