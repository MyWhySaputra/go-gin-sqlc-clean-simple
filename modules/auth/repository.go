package auth

import (
	"context"

	"github.com/MyWhySaputra/go-gin-sqlc-clean-simple/internal/database"
	"github.com/jackc/pgx/v5"
)

type AuthRepository struct {
	DB *pgx.Conn
}

func (r AuthRepository) Create(req database.CreateUserParams) (database.User, error) {
	queries := database.New(r.DB)
	user, err := queries.CreateUser(context.Background(), req)
	return user, err
}

func (r AuthRepository) ReadByEmail(email string) (database.User, error) {
	queries := database.New(r.DB)
	user, err := queries.GetUserByEmail(context.Background(), email)
	return user, err
}