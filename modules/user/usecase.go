package user

import (
	"github.com/MyWhySaputra/go-gin-sqlc-clean-simple/internal/database"
)

type UserUsecase struct {
	UserRepository UserRepository
}

func (u UserUsecase) ReadAll() ([]database.User, error) {
	return u.UserRepository.ReadAll()
}

func (u UserUsecase) ReadById(id int64) (database.User, error) {
	return u.UserRepository.ReadById(id)
}

func (u UserUsecase) Update(id int64, req *database.CreateUserParams) (database.User, error) {
	return u.UserRepository.Update(id, req)
}

func (u UserUsecase) Delete(id int64) error {
	return u.UserRepository.Delete(id)
}