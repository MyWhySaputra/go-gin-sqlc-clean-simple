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
        c.Abort()
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
        c.Abort()
        return
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        if exp, ok := claims["exp"].(float64); ok && float64(time.Now().Unix()) > exp {
            utils.HandleError(c, http.StatusUnauthorized, "Unauthorized: Token expired")
            c.Abort()
            return
        }

        id64, ok := claims["id"].(float64)
        if !ok {
            utils.HandleError(c, http.StatusInternalServerError, "Error parsing user ID")
            c.Abort()
            return
        }

        db := config.ConnectToDB()
        user, err := database.New(db).GetUser(context.Background(), int64(id64))
        if err != nil {
            utils.HandleError(c, http.StatusInternalServerError, fmt.Sprintf("Error fetching user: %v", err))
            c.Abort()
            return
        }

        if user.ID == 0 {
            utils.HandleError(c, http.StatusUnauthorized, "Unauthorized: User not found")
            c.Abort()
            return
        }

        c.Set("user", user)
        c.Next()
    } else {
        utils.HandleError(c, http.StatusUnauthorized, "Unauthorized: Invalid token")
        c.Abort()
        return
    }
}
