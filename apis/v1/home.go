package v1

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	auth "github.com/anubhavitis/Library/apis/middleware"
	DB "github.com/anubhavitis/Library/databases"
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

type result struct {
	Done  bool   `json:"done"`
	Token string `json:"token"`
	Error error  `json:"error"`
}

//SignUp handler
func SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applicaton/json")
	var NewUser DB.Member
	res := &result{
		Done:  false,
		Token: "",
		Error: nil,
	}

	if err := json.NewDecoder(r.Body).Decode(&NewUser); err != nil {
		res.Error = err
		resJ, e := json.Marshal(res)
		if e != nil {
			fmt.Println(e)
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		w.Write(resJ)
		return
	}

	if EmailCheck, e := DB.FindEmail(NewUser.Email); (e != nil || EmailCheck != DB.Member{}) {
		if e != nil {
			fmt.Println("Error at finding user with email at signup", e)
			w.WriteHeader(http.StatusConflict)
			res.Error = e
			resJ, err := json.Marshal(res)
			if err != nil {
				fmt.Println(err)
				return
			}

			w.WriteHeader(http.StatusBadRequest)
			w.Write(resJ)
			return
		}
		res.Error = errors.New("email unavailable")
		resJ, e := json.Marshal(res)
		if e != nil {
			fmt.Println(e)
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		w.Write(resJ)
		return
	}

	if UserCheck, e := DB.FindUser(NewUser.UserName); (e != nil || UserCheck != DB.Member{}) {
		if e != nil {
			w.WriteHeader(http.StatusConflict)
			res.Error = e
			resJ, e := json.Marshal(res)
			if e != nil {
				fmt.Println(e)
			}
			w.Write(resJ)
			return
		}
		res.Error = errors.New("username unavailable")
		resJ, e := json.Marshal(res)
		if e != nil {
			fmt.Println(e)
		}

		w.WriteHeader(http.StatusUnauthorized)
		w.Write(resJ)
		return
	}

	if e := DB.AddMember(NewUser); e != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		res.Error = e
		resJ, e := json.Marshal(res)
		if e != nil {
			fmt.Println(e)
		}
		w.Write(resJ)
		return
	}

	expTime := time.Now().Add(30 * time.Minute)

	claims := &auth.Claims{
		Username: NewUser.UserName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenstr, err := token.SignedString(auth.JwtKey)

	if err != nil {
		res.Error = err
		resJ, e := json.Marshal(res)
		if e != nil {
			fmt.Println(e)
		}

		w.WriteHeader(http.StatusUnauthorized)
		w.Write(resJ)
		return
	}

	res.Done = true
	res.Token = tokenstr
	resJ, e := json.Marshal(res)
	if e != nil {
		fmt.Println(e)
	}

	w.WriteHeader(http.StatusUnauthorized)
	w.Write(resJ)
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
