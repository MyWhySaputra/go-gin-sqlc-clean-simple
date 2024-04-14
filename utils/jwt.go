package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(id int64) (string, error) {

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "id": id,
        "exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
    })

    tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
    if err != nil {
        return "", err
    }

    return tokenString, nil
}
