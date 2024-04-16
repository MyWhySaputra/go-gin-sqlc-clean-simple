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
	r.POST("/logout", AuthHandler.Logout)
	
	authRoutes := r.Group("/")
  authRoutes.Use(middlewares.RequireAuth)
	
	authRoutes.GET("/validate", AuthHandler.Validate)

	userRepo := user.UserRepository{DB: db}
	userUsecase := user.UserUsecase{UserRepository: userRepo}
	userHandler := user.UserHandler{UserUsecase: userUsecase}

	authRoutes.GET("/users",  userHandler.ReadAll)
	authRoutes.GET("/users/:id", userHandler.ReadById)
	authRoutes.PUT("/users/:id", userHandler.Update)
	authRoutes.DELETE("/users/:id", userHandler.Delete)

	profileRepo := profile.ProfileRepository{DB: db}
	profileUsecase := profile.ProfileUsecase{ProfileRepository: profileRepo}
	profileHandler := profile.ProfileHandler{ProfileUsecase: profileUsecase}

	authRoutes.POST("/profiles", profileHandler.Create)
	authRoutes.GET("/profiles", profileHandler.ReadById)
	authRoutes.GET("/profiles/all", profileHandler.ReadAll)
	authRoutes.PUT("/profiles", profileHandler.Update)
	authRoutes.DELETE("/profiles/:id", profileHandler.Delete)

	fmt.Println("starting web server at localhost:", PORT)
	err := r.Run(":" + PORT)
	if err != nil {
		log.Fatal(err)
	}
}