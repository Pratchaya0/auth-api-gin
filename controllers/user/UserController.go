package controllers

import (
	"net/http"

	"github.com/Pratchaya0/auth-api-gin/helpers"
	"github.com/Pratchaya0/auth-api-gin/models"
	"github.com/Pratchaya0/auth-api-gin/usecases"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userUseCase *usecases.UserUseCase
}

func NewUserController(useCase *usecases.UserUseCase) *UserController {
	return &UserController{userUseCase: useCase}
}

func (ctrl *UserController) List(c *gin.Context) {
	users, err := ctrl.userUseCase.GetListUser()
	if err != nil {
		helpers.WebResponseWithJSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	helpers.WebResponseWithJSON(c, http.StatusOK, "OK", users)
}

func (ctrl *UserController) GetUserByUserName(c *gin.Context) {
	username := c.Param("username")

	user, err := ctrl.userUseCase.GetUserByUserName(username)
	if err != nil {
		helpers.WebResponseWithJSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	helpers.WebResponseWithJSON(c, http.StatusOK, "OK", user)
}

func (ctrl *UserController) GetUserByProviderId(c *gin.Context) {
	providerID := c.Param("id")

	user, err := ctrl.userUseCase.GetUserByProviderId(providerID)
	if err != nil {
		helpers.WebResponseWithJSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	helpers.WebResponseWithJSON(c, http.StatusOK, "OK", user)
}

func (ctrl *UserController) Update(c *gin.Context) {
	var request models.User

	if err := c.ShouldBind(&request); err != nil {
		helpers.WebResponseWithJSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if err := ctrl.userUseCase.UpdateUser(request); err != nil {
		helpers.WebResponseWithJSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	helpers.WebResponseWithJSON(c, http.StatusOK, "OK", request)
}

func (ctrl *UserController) Delete(c *gin.Context) {
	id := c.Param("id")

	if err := ctrl.userUseCase.DeleteUser(id); err != nil {
		helpers.WebResponseWithJSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	helpers.WebResponseWithJSON(c, http.StatusOK, "OK", nil)
}
