package helpers

import (
	dtos "github.com/Pratchaya0/auth-api-gin/dtos/responses"
	"github.com/gin-gonic/gin"
)

func WebResponseWithJSON(c *gin.Context, status int, message string, data interface{}) {
	webResponse := dtos.WebResponse{
		Status:  status,
		Message: message,
		Data:    data,
	}

	c.JSON(status, webResponse)
}
