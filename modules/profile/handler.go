package profile

import (
	"net/http"
	"strconv"

	"github.com/MyWhySaputra/go-gin-sqlc-clean-simple/internal/database"
	"github.com/MyWhySaputra/go-gin-sqlc-clean-simple/utils"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

type ProfileHandler struct {
	ProfileUsecase ProfileUsecase
}

func (h ProfileHandler) Create(c *gin.Context) {
	token, _ := c.Get("user")
	if token == nil {
		return
	}

	var request ProfileResponse
	if err := c.Bind(&request); err != nil {
		utils.HandleError(c, http.StatusBadRequest, err.Error())
		return
	}

	user, _ := c.Get("user")
	userModel := user.(database.User)

	intUser := pgtype.Int8{Int64: userModel.ID,Valid: true}

	req := database.CreateProfileParams{
		UserID:  intUser,
		Name:    request.Name,
		Bio:     request.Bio,
	}

	profile, err := h.ProfileUsecase.Create(&req)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.HandleSuccess(c, profile)
}

func (h ProfileHandler) ReadById(c *gin.Context) {
	token, _ := c.Get("user")
	if token == nil {
		return
	}

	user, _ := c.Get("user")
	userModel := user.(database.User)

	UserID := pgtype.Int8{Int64: userModel.ID, Valid: true}

	profile, err := h.ProfileUsecase.ReadById(UserID)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.HandleSuccess(c, profile)
}

func (h ProfileHandler) ReadAll(c *gin.Context) {
	token, _ := c.Get("user")
	if token == nil {
		return
	}

	profiles, err := h.ProfileUsecase.ReadAll()
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}

	var viewProfiles []ProfileResponse

	for _, profile := range profiles {
		viewUser := ProfileResponse{
			Name:  profile.Name,
			Bio:   profile.Bio,
		}
		viewProfiles = append(viewProfiles, viewUser)
	}

	utils.HandleSuccess(c, viewProfiles)
}

func (h ProfileHandler) Update(c *gin.Context) {
	token, _ := c.Get("user")
	if token == nil {
		return
	}

	var request ProfileResponse
	if err := c.Bind(&request); err != nil {
		utils.HandleError(c, http.StatusBadRequest, err.Error())
		return
	}
	user, _ := c.Get("user")
	userModel := user.(database.User)

	req := database.UpdateProfileByUserIdParams{
		UserID:  pgtype.Int8{Int64: userModel.ID, Valid: true},
		Name:    request.Name,
		Bio:     request.Bio,
	}

	profile, err := h.ProfileUsecase.Update(&req)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.HandleSuccess(c, profile)
}

func (h ProfileHandler) Delete(c *gin.Context) {
	token, _ := c.Get("user")
	if token == nil {
		return
	}
	
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

	id8 := pgtype.Int8{Int64: int64(id), Valid: true}

	_, err = h.ProfileUsecase.ReadById(id8)
	if err != nil {
		utils.HandleError(c, http.StatusNotFound, err.Error())
		return
	}

	err = h.ProfileUsecase.Delete(id8)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.HandleSuccess(c, "Success delete profile")
}