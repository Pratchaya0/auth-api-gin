package controllers

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/Pratchaya0/auth-api-gin/helpers"
	"github.com/Pratchaya0/auth-api-gin/models"
	"github.com/Pratchaya0/auth-api-gin/usecases"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type CredentialController struct {
	userUseCase *usecases.UserUseCase
}

func NewCredentialController(userUseCase *usecases.UserUseCase) *CredentialController {
	return &CredentialController{userUseCase: userUseCase}
}

func (ctrl *CredentialController) CurrentUser(c *gin.Context) {
	cookie, _ := c.Cookie("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SESSION_SECRET")), nil
	})

	if err != nil {
		helpers.WebResponseWithJSON(c, http.StatusUnauthorized, "", nil)
		return
	}

	claims := token.Claims.(*jwt.StandardClaims)

	user, err := ctrl.userUseCase.GetUserById(claims.Issuer)
	if err != nil {
		helpers.WebResponseWithJSON(c, http.StatusNotFound, "unauthenticated", nil)
		return
	}

	helpers.WebResponseWithJSON(c, http.StatusOK, "OK", user)
}

func (ctrl *CredentialController) Register(c *gin.Context) {
	var data map[string]string

	if err := c.ShouldBind(&data); err != nil {
		helpers.WebResponseWithJSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.User{
		FirstName: data["firstname"],
		LastName:  data["lastname"],
		Email:     data["email"],
		Password:  password,
	}

	err := ctrl.userUseCase.CreateUser(user)
	if err != nil {
		helpers.WebResponseWithJSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	helpers.WebResponseWithJSON(c, http.StatusOK, "OK", nil)
}

func (ctrl *CredentialController) Login(c *gin.Context) {
	var data map[string]string

	if err := c.ShouldBind(&data); err != nil {
		helpers.WebResponseWithJSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	user, err := ctrl.userUseCase.GetUserByEmail(data["email"])
	if err != nil {
		helpers.WebResponseWithJSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if user.ID == 0 {
		helpers.WebResponseWithJSON(c, http.StatusNotFound, "user not found", nil)
		return
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		helpers.WebResponseWithJSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(os.Getenv("SESSION_SECRET")))
	if err != nil {
		helpers.WebResponseWithJSON(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	c.SetCookie("jwt", token, 3600*24, "/", "", false, true)

	helpers.WebResponseWithJSON(c, http.StatusOK, "OK", gin.H{
		"message": "login success",
		"token":   "Bearer " + token,
	})
}

func (ctrl *CredentialController) LogOut(c *gin.Context) {
	c.SetCookie("jwt", "", -1, "/", "", false, true)

	helpers.WebResponseWithJSON(c, http.StatusOK, "OK", gin.H{"message": "logout success"})
}
