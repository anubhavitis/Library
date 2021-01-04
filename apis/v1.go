package apis

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/anubhavitis/Library/auth"
	DB "github.com/anubhavitis/Library/databases"
	"github.com/dgrijalva/jwt-go"
)

//Homepage handler
func Homepage(w http.ResponseWriter, r *http.Request) {
	html := `
	<html> <body> 
		<h1> Welcome to TestAPIs </h1>
		<a href="\signup"> SignUp</a>
	</body> </html>`

	fmt.Fprintln(w, html)
}

var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

//SignIn handler
func SignIn(w http.ResponseWriter, r *http.Request) {
	var cred auth.UserCred

	if e := json.NewDecoder(r.Body).Decode(&cred); e != nil {
		fmt.Println("Error at decoding request", e)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	expectedPassword, ok := users[cred.Username]

	if !ok || expectedPassword != cred.Password {
		fmt.Println("Error at matching password", ok)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expTime := time.Now().Add(30 * time.Minute)

	claims := &auth.Claims{
		Username: cred.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenstr, err := token.SignedString(auth.JwtKey)

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

//SignUp handler
func SignUp(w http.ResponseWriter, r *http.Request) {
	var NewUser DB.Member
	Mydb, err := DB.InitDb()
	if err != nil {
		fmt.Println("DB: ", Mydb)
	}

	if err := json.NewDecoder(r.Body).Decode(&NewUser); err != nil {
		fmt.Println("Error at parsing signin request: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if EmailCheck, e := DB.FindEmail(Mydb, NewUser.Email); (e != nil || EmailCheck != DB.Member{}) {
		if e != nil {
			fmt.Println("Error at finding user with email at signup", e)
			w.WriteHeader(http.StatusConflict)
			return
		}

		fmt.Fprintf(w, "Email already exists!")
		return
	}

	if UserCheck, e := DB.FindUser(Mydb, NewUser.UserName); (e != nil || UserCheck != DB.Member{}) {
		if e != nil {
			fmt.Println("Error at finding user with username at signup", e)
			w.WriteHeader(http.StatusConflict)
			return
		}

		fmt.Fprintf(w, "Username already exists!")
		return
	}

	if e := DB.AddMember(Mydb, NewUser); e != nil {
		fmt.Fprintf(w, e.Error())
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	fmt.Fprintf(w, "User Added Successfully!!")
}

//Welcome handler
func Welcome(w http.ResponseWriter, r *http.Request) {
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
	claims := &auth.Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return auth.JwtKey, nil
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

	Homepage(w, r)
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
	claims := &auth.Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return auth.JwtKey, nil
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
	tokenString, err := token.SignedString(auth.JwtKey)
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
