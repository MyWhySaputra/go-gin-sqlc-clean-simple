package user

import (
	"github.com/MyWhySaputra/go-gin-sqlc-clean-simple/internal/database"
)

type Usercase struct {
	UserRepository UserRepository
}

func (u Usercase) Create(req database.CreateUserParams) (database.User, error) {
	return u.UserRepository.Create(req)
}
