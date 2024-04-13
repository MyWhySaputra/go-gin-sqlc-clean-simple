package user

import (
	"net/http"

	"github.com/MyWhySaputra/go-gin-sqlc-clean-simple/internal/database"
	"github.com/MyWhySaputra/go-gin-sqlc-clean-simple/utils"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Usercase Usercase
}

func (h UserHandler) Create(c *gin.Context) {
	var request database.User
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	params := database.CreateUserParams{
		Email:	request.Email,
		Password:  request.Password,
	}

	user, err := h.Usercase.Create(params)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.HandleSuccess(c, user)
}

func (h UserHandler) ReadById(c *gin.Context) {
	// TODO
}

func (h UserHandler) ReadAll(c *gin.Context) {
	// TODO
}

func (h UserHandler) Update(c *gin.Context) {
	// TODO
}

func (h UserHandler) Delete(c *gin.Context) {
	// TODO
}