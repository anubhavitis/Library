package v1

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/anubhavitis/Library/apis/middleware"
	"github.com/anubhavitis/Library/apis/utility"
	DB "github.com/anubhavitis/Library/databases"
	database "github.com/anubhavitis/Library/databases"
	jwtauth "github.com/anubhavitis/Library/pkg/auth"
	"github.com/anubhavitis/Library/pkg/email"
	"github.com/dgrijalva/jwt-go"
)

//SignIn handler
func SignIn(w http.ResponseWriter, r *http.Request) {
	var cred middleware.UserCred
	var res utility.Result
	if e := json.NewDecoder(r.Body).Decode(&cred); e != nil {
		res.Error = e
		utility.SendResponse(w, http.StatusBadRequest, res)
		return
	}

	user, ok := DB.FindUser(cred.Username)
	fmt.Println("Checked username")

	if (ok != nil || user == DB.Member{} || user.Password != cred.Password) {
		if ok != nil {
			fmt.Println(ok)
			res.Error = ok
			utility.SendResponse(w, http.StatusUnauthorized, res)
			return
		}
		res.Error = errors.New("username/password combination not found")
		utility.SendResponse(w, http.StatusUnauthorized, res)
		return
	}
	fmt.Println("Checked Password")

	tokenstr, err := jwtauth.GenerateToken(user)
	if err != nil {
		res.Error = err
		utility.SendResponse(w, http.StatusInternalServerError, res)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "Token",
		Value:   tokenstr,
		Expires: time.Now().Add(30 * time.Minute),
	})

	res.Success = true
	utility.SendResponse(w, http.StatusAccepted, res)
}

//SignUp handler
func SignUp(w http.ResponseWriter, r *http.Request) {
	var NewUser DB.Member
	var res utility.Result

	if err := json.NewDecoder(r.Body).Decode(&NewUser); err != nil {
		res.Error = err
		utility.SendResponse(w, 400, res)
		return
	}
	if EmailCheck, e := DB.FindEmail(NewUser.Email); (e != nil || EmailCheck != DB.Member{}) {
		if e != nil {
			res.Error = e
			utility.SendResponse(w, http.StatusInternalServerError, res)
			return
		}
		res.Error = errors.New("email unavailable")
		utility.SendResponse(w, 307, res)
		return
	}
	if UserCheck, e := DB.FindUser(NewUser.UserName); (e != nil || UserCheck != DB.Member{}) {
		if e != nil {
			res.Error = e
			utility.SendResponse(w, http.StatusInternalServerError, res)
			return
		}
		res.Error = errors.New("username unavailable")
		utility.SendResponse(w, 307, res)
		return
	}
	token, err := jwtauth.GenerateToken(NewUser)
	if err != nil {
		res.Error = err
		utility.SendResponse(w, 400, res)
		return
	}

	str := "localhost:8080/verify?token=" + token
	ok := email.SendWelcomeEmail(NewUser.Email, NewUser.Fname+NewUser.Lname, str)
	if !ok {
		res.Error = errors.New("error at contacting given email")
		utility.SendResponse(w, http.StatusBadRequest, res)
	}

	res.Success = true
	res.Body = map[string]interface{}{
		"Action": "Check given email for confirmation",
	}
	utility.SendResponse(w, 200, res)
}

//VerifyEmail func
func VerifyEmail(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenStr, ok := r.URL.Query()["token"]
		var res utility.Result

		if !ok || len(tokenStr[0]) < 1 {
			res.Error = errors.New("token not found")
			utility.SendResponse(w, http.StatusBadRequest, res)
			return
		}

		NewUser, e := jwtauth.ExtractClaims(tokenStr[0])
		if !e {
			res.Error = errors.New("error while extracting values form token")
			utility.SendResponse(w, http.StatusBadRequest, res)
			return
		}
		if e := database.AddMember(NewUser); e != nil {
			res.Error = e
			utility.SendResponse(w, http.StatusExpectationFailed, res)
			return
		}

		f.ServeHTTP(w, r)
	}
}

//Refresh handler
func Refresh(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("Token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
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
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	expTime := time.Now().Add(30 * time.Minute)
	claims.ExpiresAt = expTime.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtauth.JwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expTime,
	})
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
