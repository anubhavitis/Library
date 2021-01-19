package middleware

import (
	"fmt"
	"net/http"

	"github.com/anubhavitis/Library/apis/utility"
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
		var res utility.Result
		tokenStr, err := utility.GetTokenFromCookie(r)

		if err != nil {
			res.Error = fmt.Sprintf("%s", err)
			utility.SendResponse(w, http.StatusBadRequest, res)
			return
		}
		claims := &jwtauth.Claims{}

		tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtauth.JwtKey, nil
		})

		if err != nil {
			res.Error = fmt.Sprintf("%s", err)
			utility.SendResponse(w, http.StatusBadRequest, res)
			return
		}

		if !tkn.Valid {
			res.Error = "not authorised, please sign in"
			utility.SendResponse(w, http.StatusUnauthorized, res)
			return
		}
		f.ServeHTTP(w, r)
	}
}
