package middleware

import (
	"fmt"
	"net/http"

	jwtauth "github.com/anubhavitis/Library/pkg/auth"
	jwt "github.com/dgrijalva/jwt-go"
)

//UserCred as expected User Credential while login
type UserCred struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

//Auth function
func Auth(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		c, err := r.Cookie("Token")
		if err != nil {
			if err == http.ErrNoCookie {
				fmt.Fprintf(w, "No Cookie found!")
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			fmt.Fprintf(w, "Request Invalid!")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		tokenStr := c.Value
		claims := &jwtauth.Claims{}

		tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtauth.JwtKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				fmt.Fprintf(w, "UnAuthorized!")
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			fmt.Fprintf(w, "Request Invalid!"+err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if !tkn.Valid {
			fmt.Fprintf(w, "Request Invalid!")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		f.ServeHTTP(w, r)
	}
}
