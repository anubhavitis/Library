package auth

import (
	"os"
	"time"

	database "github.com/anubhavitis/Library/databases"
	"github.com/dgrijalva/jwt-go"
)

//JwtKey as a secret Key
var JwtKey = []byte(os.Getenv("Secret-Key"))

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken(user database.Member) (string, error) {
	expTime := time.Now().Add(30 * time.Minute)

	claims := &Claims{
		Username: user.UserName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims.StandardClaims)
	tokenstr, err := token.SignedString(JwtKey)
	return tokenstr, err
}
