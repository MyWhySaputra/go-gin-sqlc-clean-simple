package user

import (
	"context"

	"github.com/MyWhySaputra/go-gin-sqlc-clean-simple/internal/database"
	"github.com/jackc/pgx/v5"
)

type UserRepository struct {
	DB *pgx.Conn
}

func (r UserRepository) ReadAll() ([]database.User, error) {
	queries := database.New(r.DB)
	user, err := queries.ListUser(context.Background())
	return user, err
}

func (r UserRepository) ReadById(id int64) (database.User, error) {
	queries := database.New(r.DB)
	user, err := queries.GetUser(context.Background(), id)
	return user, err
}

func (r UserRepository) Update(id int64, req *database.CreateUserParams) (database.User, error) {
	queries := database.New(r.DB)
	user, err := queries.UpdateUser(context.Background(), database.UpdateUserParams{
		Email:    req.Email,
		Password: req.Password,
		ID:       id,
	})
	return user, err
}

func (r UserRepository) Delete(id int64) error {
	queries := database.New(r.DB)
	err := queries.DeleteUser(context.Background(), id)
	return err
}