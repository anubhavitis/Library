package middleware

import (
	jwt "github.com/dgrijalva/jwt-go"
)

//JwtKey as a secret Key
var JwtKey = []byte("my_secret_key")

//UserCred as expected User Credential while login
type UserCred struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

//Claims struct for jwt
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
