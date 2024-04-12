package main

import (
	"fmt"
	"log"
	"os"

	"github.com/MyWhySaputra/go-gin-sqlc-clean-simple/config"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectToDB()
}

func main() {
	PORT := os.Getenv("PORT")

	r := gin.Default()

	

	fmt.Println("starting web server at localhost:", PORT)
	err := r.Run(":" + PORT)
	if err != nil {
		log.Fatal(err)
	}
}