package profile

import (
	"github.com/MyWhySaputra/go-gin-sqlc-clean-simple/internal/database"
	"github.com/jackc/pgx/v5/pgtype"
)

type ProfileUsecase struct {
	ProfileRepository ProfileRepository
}

func (u ProfileUsecase) Create(req *database.CreateProfileParams) (database.Profile, error) {
	return u.ProfileRepository.Create(req)
}

func (u ProfileUsecase) ReadById(user_id pgtype.Int8) (database.Profile, error) {
	return u.ProfileRepository.ReadById(user_id)
}

func (u ProfileUsecase) ReadAll() ([]database.Profile, error) {
	return u.ProfileRepository.ReadAll()
}

func (u ProfileUsecase) Update(req *database.UpdateProfileByUserIdParams) (database.Profile, error) {
	return u.ProfileRepository.Update(req)
}

func (u ProfileUsecase) Delete(user_id pgtype.Int8) error {
	return u.ProfileRepository.Delete(user_id)
}