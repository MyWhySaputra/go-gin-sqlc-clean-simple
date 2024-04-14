package main

import (
	"fmt"
	"log"
	"os"

	"github.com/MyWhySaputra/go-gin-sqlc-clean-simple/middlewares"
	"github.com/MyWhySaputra/go-gin-sqlc-clean-simple/modules/auth"
	"github.com/MyWhySaputra/go-gin-sqlc-clean-simple/modules/user"
	"github.com/MyWhySaputra/go-gin-sqlc-clean-simple/modules/profile"

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

	AuthRepo := auth.AuthRepository{DB: db}
	AuthUsecase := auth.AuthUsecase{AuthRepository: AuthRepo}
	AuthHandler := auth.AuthHandler{AuthUsecase: AuthUsecase}

	r.POST("/signup", AuthHandler.Signup)
	r.POST("/login", AuthHandler.Login)
	r.GET("/validate", middlewares.RequireAuth, AuthHandler.Validate)
	r.POST("/logout", AuthHandler.Logout)

	userRepo := user.UserRepository{DB: db}
	userUsecase := user.UserUsecase{UserRepository: userRepo}
	userHandler := user.UserHandler{UserUsecase: userUsecase}

	r.GET("/users", userHandler.ReadAll)
	r.GET("/users/:id", userHandler.ReadById)
	r.POST("/users:id", userHandler.Update)
	r.DELETE("/users/:id", userHandler.Delete)

	profileRepo := profile.ProfileRepository{DB: db}
	profileUsecase := profile.ProfileUsecase{ProfileRepository: profileRepo}
	profileHandler := profile.ProfileHandler{ProfileUsecase: profileUsecase}

	r.POST("/profiles", middlewares.RequireAuth, profileHandler.Create)
	r.GET("/profiles", middlewares.RequireAuth, profileHandler.ReadById)
	r.GET("/profiles", middlewares.RequireAuth, profileHandler.ReadAll)
	r.PUT("/profiles", middlewares.RequireAuth, profileHandler.Update)
	r.DELETE("/profiles/:id", middlewares.RequireAuth, profileHandler.Delete)

	fmt.Println("starting web server at localhost:", PORT)
	err := r.Run(":" + PORT)
	if err != nil {
		log.Fatal(err)
	}
}