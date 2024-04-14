package profile

import (
	"context"

	"github.com/MyWhySaputra/go-gin-sqlc-clean-simple/internal/database"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type ProfileRepository struct {
	DB *pgx.Conn
}

func (r ProfileRepository) Create(req *database.CreateProfileParams) (database.Profile, error) {
	queries := database.New(r.DB)
	profile, err := queries.CreateProfile(context.Background(), *req)
	return profile, err
}

func (r ProfileRepository) ReadById(user_id pgtype.Int8) (database.Profile, error) {
	queries := database.New(r.DB)
	profile, err := queries.GetProfileByUserId(context.Background(), user_id)
	return profile, err
}

func (r ProfileRepository) ReadAll() ([]database.Profile, error) {
	queries := database.New(r.DB)
	profile, err := queries.ListProfiles(context.Background())
	return profile, err
}

func (r ProfileRepository) Update(req *database.UpdateProfileByUserIdParams) (database.Profile, error) {
	queries := database.New(r.DB)
	profile, err := queries.UpdateProfileByUserId(context.Background(), *req)
	return profile, err
}

func (r ProfileRepository) Delete(user_id pgtype.Int8) error {
	queries := database.New(r.DB)
	err := queries.DeleteProfileByUserId(context.Background(), user_id)
	return err
}