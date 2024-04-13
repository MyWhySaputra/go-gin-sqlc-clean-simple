package main

import (
	"fmt"
	"log"
	"os"

	"github.com/MyWhySaputra/go-gin-sqlc-clean-simple/modules/user"

	"github.com/MyWhySaputra/go-gin-sqlc-clean-simple/config"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnvVariables()
}

func main() {
	PORT := os.Getenv("PORT")

	r := gin.Default()

	db := config.ConnectToDB()
	if db == nil {
		log.Fatal("Database connection failed")
	}

	userRepo := user.UserRepository{DB: db}

	usercase := user.Usercase{UserRepository: userRepo}

	userHandler := user.UserHandler{Usercase: usercase}

	r.POST("/users", userHandler.Create)

	fmt.Println("starting web server at localhost:", PORT)
	err := r.Run(":" + PORT)
	if err != nil {
		log.Fatal(err)
	}
}