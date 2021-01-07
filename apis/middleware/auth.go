package middleware

import (
	"fmt"
	"net/http"

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

//Auth function
func Auth(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("MiddleWare Initiated")
		defer fmt.Println("MiddleWare ended")

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
		claims := &Claims{}

		tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return JwtKey, nil
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

//Homepage handler
func Homepage(w http.ResponseWriter, r *http.Request) {
	html := `
	<html> <body> 
		<h1> Welcome to TestAPIs </h1>
		<a href="\signup"> SignUp</a>
	</body> </html>`

	fmt.Fprintln(w, html)
}
