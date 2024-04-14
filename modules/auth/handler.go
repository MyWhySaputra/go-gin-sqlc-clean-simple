package auth

import (
	"net/http"

	"github.com/MyWhySaputra/go-gin-sqlc-clean-simple/internal/database"
	"github.com/MyWhySaputra/go-gin-sqlc-clean-simple/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	AuthUsecase AuthUsecase
}

func (h AuthHandler) Signup(c *gin.Context) {
	var request database.CreateUserParams
	if err := c.Bind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	req := database.CreateUserParams{
		Email:	request.Email,
		Password:  string(hash),
	}

	user, err := h.AuthUsecase.Create(req)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	viewUser := authResponse{
		ID:        user.ID,
		Email:     user.Email,
	}

	utils.HandleSuccess(c, viewUser)
}

func (h AuthHandler) Login(c *gin.Context) {
	var request database.CreateUserParams
	if err := c.Bind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.AuthUsecase.ReadByEmail(request.Email)
	if err != nil {
		utils.HandleError(c, http.StatusNotFound, err.Error())
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		utils.HandleError(c, http.StatusUnauthorized, err.Error())
		return
	}

	tokenString, err := utils.GenerateToken(user.ID)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600 * 24 * 30, "", "", false, true)

	utils.HandleSuccess(c, "Success login")
}

func (h AuthHandler) Validate(c *gin.Context) {
	user, _ := c.Get("user")

	userModel := user.(authResponse)

	c.JSON(http.StatusOK, gin.H{
		"message": "You are logged in",
		"user": userModel.ID,
	})
}

func (h AuthHandler) Logout(c *gin.Context) {
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", "", -1, "", "", false, true)
	utils.HandleSuccess(c, "Success logout")
}