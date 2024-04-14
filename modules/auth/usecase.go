package auth

import "github.com/MyWhySaputra/go-gin-sqlc-clean-simple/internal/database"

type AuthUsecase struct {
	AuthRepository AuthRepository
}

func (u AuthUsecase) Create(req database.CreateUserParams) (database.User, error) {
	return u.AuthRepository.Create(req)
}

func (u AuthUsecase) ReadByEmail(email string) (database.User, error) {
	return u.AuthRepository.ReadByEmail(email)
}