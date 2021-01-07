package v1

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	auth "github.com/anubhavitis/Library/apis/middleware"
	DB "github.com/anubhavitis/Library/databases"
	jwtauth "github.com/anubhavitis/Library/pkg/auth"
	"github.com/dgrijalva/jwt-go"
)

//SignIn handler
func SignIn(w http.ResponseWriter, r *http.Request) {
	var cred auth.UserCred

	if e := json.NewDecoder(r.Body).Decode(&cred); e != nil {
		fmt.Println("Error at decoding request", e)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user, ok := DB.FindUser(cred.Username)

	if (ok != nil || user == DB.Member{} || user.Password != cred.Password) {
		fmt.Println("Error at matching password", ok)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expTime := time.Now().Add(30 * time.Minute)
	var claims jwtauth.Claims
	claims.Username = cred.Username
	claims.ExpiresAt = expTime.Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenstr, err := token.SignedString(jwtauth.JwtKey)

	if err != nil {
		fmt.Println("Error at making token string", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "Token",
		Value:   tokenstr,
		Expires: expTime,
	})

	fmt.Fprintf(w, " Cookie thing worked")
}

type result struct {
	Success bool   `json:"done"`
	Token   string `json:"token"`
	Error   error  `json:"error"`
}

//SignUp handler
func SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applicaton/json")
	var NewUser DB.Member
	res := &result{
		Success: false,
		Token:   "",
		Error:   nil,
	}

	if err := json.NewDecoder(r.Body).Decode(&NewUser); err != nil {
		res.Error = err
		sendResponse(w, 400, res)
		return
	}

	if EmailCheck, e := DB.FindEmail(NewUser.Email); (e != nil || EmailCheck != DB.Member{}) {
		if e != nil {
			fmt.Println("Error at finding user with email at signup", e)
			res.Error = e
			sendResponse(w, http.StatusInternalServerError, res)
			return
		}
		res.Error = errors.New("email unavailable")
		sendResponse(w, 307, res)
		return
	}

	if UserCheck, e := DB.FindUser(NewUser.UserName); (e != nil || UserCheck != DB.Member{}) {
		if e != nil {
			res.Error = e
			sendResponse(w, http.StatusInternalServerError, res)
			return
		}
		res.Error = errors.New("username unavailable")
		sendResponse(w, 307, res)
		return
	}

	if e := DB.AddMember(NewUser); e != nil {
		res.Error = e
		sendResponse(w, http.StatusExpectationFailed, res)
		return
	}

	token, err := jwtauth.GenerateToken(NewUser)
	if err != nil {
		res.Error = err
		sendResponse(w, 400, res)
		return
	}
	res.Success = true
	res.Token = token
	sendResponse(w, 200, res)
	return
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