package v1

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/anubhavitis/Library/apis/utility"
	database "github.com/anubhavitis/Library/databases"
	jwtauth "github.com/anubhavitis/Library/pkg/auth"
	"github.com/dgrijalva/jwt-go"
)

//Auth handler
func Auth(next func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var res utility.Result

		var Token struct {
			Tokenstr string `json:"token"`
		}

		err := json.NewDecoder(r.Body).Decode(&Token)

		if err != nil {
			res.Error = fmt.Sprintf("%s", err)
			utility.SendResponse(w, http.StatusUnauthorized, res)
			return
		}

		claims := &jwtauth.Claims{}

		tkn, err := jwt.ParseWithClaims(Token.Tokenstr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtauth.JwtKey, nil
		})

		if err != nil {
			res.Error = fmt.Sprintf("%s", err)
			utility.SendResponse(w, http.StatusUnauthorized, res)
			return
		}

		UserCheck, e := database.FindUser(claims.Username)

		if (!tkn.Valid || e != nil || UserCheck == database.Member{}) {
			if e != nil {
				res.Error = fmt.Sprintf("%s", e)
				utility.SendResponse(w, http.StatusUnauthorized, res)
				return
			}
			res.Error = "not authorised, please sign in"
			utility.SendResponse(w, http.StatusUnauthorized, res)
			return
		}

		next(w, r)
	})
}
