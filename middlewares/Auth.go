package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/MyWhySaputra/go-gin-sqlc-clean-simple/config"
	"github.com/MyWhySaputra/go-gin-sqlc-clean-simple/internal/database"
	"github.com/MyWhySaputra/go-gin-sqlc-clean-simple/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		utils.HandleError(c, http.StatusUnauthorized, "Unauthorized: Token not found")
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		utils.HandleError(c, http.StatusUnauthorized, fmt.Sprintf("Unauthorized: %v", err))
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
    if float64(time.Now().Unix()) > claims["exp"].(float64) {
        utils.HandleError(c, http.StatusUnauthorized, "Unauthorized: Token expired")
        return
    }
    
    id64 := int64(claims["id"].(float64))
    
    auth, err := database.New(config.ConnectToDB()).GetUser(context.Background(), id64)
    if err != nil {
        utils.HandleError(c, http.StatusInternalServerError, fmt.Sprintf("Error fetching user: %v", err))
        return
    }
    
    if auth.ID == 0 {
        utils.HandleError(c, http.StatusUnauthorized, "Unauthorized: User not found")
        return
    }
    
    c.Set("user", auth)
    c.Next()
} else {
    utils.HandleError(c, http.StatusUnauthorized, "Unauthorized: Invalid token")
    return
}
}