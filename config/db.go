package config

import (
    "context"
    "os"

    "github.com/jackc/pgx/v5"
)

func ConnectToDB() *pgx.Conn {
    DB, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
    if err != nil {
      panic("Failed to connect to database")
    }
    return DB
}
