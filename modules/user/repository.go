package user

import (
	"context"

	"github.com/MyWhySaputra/go-gin-sqlc-clean-simple/internal/database"
	"github.com/jackc/pgx/v5"
)

type UserRepository struct {
	DB *pgx.Conn
}

func (r UserRepository) Create(req database.CreateUserParams) (database.User, error) {
	queries := database.New(r.DB)
	user, err := queries.CreateUser(context.Background(), req)
	return user, err
}