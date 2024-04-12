package utils

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Respon struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func HandleSuccess(c *gin.Context, data interface{}) {
	responseData := Respon{
		Status:  "200",
		Message: "Success",
		Data:    data,
	}
	c.JSON(http.StatusOK, responseData)
}

func HandleError(c *gin.Context, status int, message string) {
	responseData := Respon{
		Status:  strconv.Itoa(status),
		Message: message,
	}
	c.JSON(status, responseData)
}
