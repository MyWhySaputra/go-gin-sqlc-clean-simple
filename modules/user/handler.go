package user

import (
	"net/http"
	"strconv"

	"github.com/MyWhySaputra/go-gin-sqlc-clean-simple/internal/database"
	"github.com/MyWhySaputra/go-gin-sqlc-clean-simple/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	UserUsecase UserUsecase
}

func (h UserHandler) ReadAll(c *gin.Context) {
	users, err := h.UserUsecase.ReadAll()
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}

	var viewUsers []UserResponse

	for _, user := range users {
		viewUser := UserResponse{
			ID:    user.ID,
			Email: user.Email,
		}
		viewUsers = append(viewUsers, viewUser)
	}

	utils.HandleSuccess(c, viewUsers)
}

func (h UserHandler) ReadById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, "id has be number")
		return
	}

	if id <= 0 {
		utils.HandleError(c, http.StatusBadRequest, "id must be greater than 0")
		return
	}

	id64 := int64(id)

	user, err := h.UserUsecase.ReadById(id64)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}

	viewUser := UserResponse{
		ID:    user.ID,
		Email: user.Email,
	}

	utils.HandleSuccess(c, viewUser)
}

func (h UserHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, "id has be number")
		return
	}

	id64 := int64(id)

	_, err = h.UserUsecase.ReadById(id64)
	if err != nil {
		utils.HandleError(c, http.StatusNotFound, err.Error())
		return
	}

	var request = database.CreateUserParams{}
	err = c.Bind(&request)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, "Oopss server someting wrong")
		return
	}
	if request.Email == "" || request.Password == "" {
		utils.HandleError(c, http.StatusBadRequest, "column cannot be empty")
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	request.Password = string(hash)

	user, err := h.UserUsecase.Update(id64, &request)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}

	viewUser := UserResponse{
		ID:    user.ID,
		Email: user.Email,
	}

	utils.HandleSuccess(c, viewUser)
}

func (h UserHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, "id has be number")
		return
	}

	if id <= 0 {
		utils.HandleError(c, http.StatusBadRequest, "id must be greater than 0")
		return
	}

	id64 := int64(id)

	_, err = h.UserUsecase.ReadById(id64)
	if err != nil {
		utils.HandleError(c, http.StatusNotFound, err.Error())
		return
	}

	err = h.UserUsecase.Delete(id64)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.HandleSuccess(c, "Success delete user")
}