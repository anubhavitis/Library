package auth

import (
	"fmt"
	"time"

	database "github.com/anubhavitis/Library/databases"
	"github.com/dgrijalva/jwt-go"
)

//JwtKey as a secret Key
// var JwtKey = []byte(os.Getenv("Secret-Key"))
var JwtKey = []byte("Surprise MotherFucker")

//Claims for JWT token
type Claims struct {
	Username string `json:"username"`
	Fname    string `json:"fname"`
	Lname    string `json:"lname"`
	Email    string `json:"email"`
	College  string `json:"college"`
	Password string `json:"password"`

	jwt.StandardClaims
}

//GenerateToken func
func GenerateToken(user database.Member) (string, error) {
	expTime := time.Now().Add(30 * time.Minute)

	claims := Claims{
		user.UserName,
		user.Fname,
		user.Lname,
		user.Email,
		user.College,
		user.Password,
		jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenstr, err := token.SignedString(JwtKey)
	return tokenstr, err
}

//ExtractClaims func
func ExtractClaims(tokenStr string) (database.Member, bool) {
	var usr database.Member
	hmacSecret := JwtKey
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return hmacSecret, nil
	})

	if err != nil {
		return usr, false
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		usr.UserName = fmt.Sprintf("%v", claims["username"])
		usr.Fname = fmt.Sprintf("%v", claims["fname"])
		usr.Lname = fmt.Sprintf("%v", claims["lname"])
		usr.College = fmt.Sprintf("%v", claims["college"])
		usr.Password = fmt.Sprintf("%v", claims["password"])
		usr.Email = fmt.Sprintf("%v", claims["email"])

		return usr, true
	}

	return usr, false
}
