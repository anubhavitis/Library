package v1

import (
	"encoding/json"
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
		res.Error = fmt.Sprintf("%s", e)
		utility.SendResponse(w, http.StatusBadRequest, res)
		return
	}

	user, ok := DB.FindUser(cred.Username)
	fmt.Println("Checked username")

	if (ok != nil || user == DB.Member{} || user.Password != cred.Password) {
		if ok != nil {
			fmt.Println(ok)
			res.Error = fmt.Sprintf("%s", ok)
			utility.SendResponse(w, http.StatusUnauthorized, res)
			return
		}
		res.Error = "username/password combination not found"
		utility.SendResponse(w, http.StatusUnauthorized, res)
		return
	}
	fmt.Println("Checked Password")

	tokenstr, err := jwtauth.GenerateToken(user)
	if err != nil {
		res.Error = fmt.Sprintf("%s", err)
		utility.SendResponse(w, http.StatusInternalServerError, res)
		return
	}

	res.Success = true
	res.Body = map[string]interface{}{
		"token": tokenstr,
	}
	utility.SendResponse(w, http.StatusAccepted, res)
}

//SignUp handler
func SignUp(w http.ResponseWriter, r *http.Request) {
	var NewUser DB.Member
	var res utility.Result

	if err := json.NewDecoder(r.Body).Decode(&NewUser); err != nil {
		res.Error = fmt.Sprintf("%s", err)
		utility.SendResponse(w, 400, res)
		return
	}
	if EmailCheck, e := DB.FindEmail(NewUser.Email); (e != nil || EmailCheck != DB.Member{}) {
		if e != nil {
			res.Error = fmt.Sprintf("%s", e)
			utility.SendResponse(w, http.StatusInternalServerError, res)
			return
		}
		res.Error = "email unavailable"
		utility.SendResponse(w, 307, res)
		return
	}
	if UserCheck, e := DB.FindUser(NewUser.UserName); (e != nil || UserCheck != DB.Member{}) {
		if e != nil {
			res.Error = fmt.Sprintf("%s", e)
			utility.SendResponse(w, http.StatusInternalServerError, res)
			return
		}
		res.Error = "username unavailable"
		utility.SendResponse(w, 307, res)
		return
	}
	token, err := jwtauth.GenerateToken(NewUser)
	if err != nil {
		res.Error = fmt.Sprintf("%s", err)
		utility.SendResponse(w, 400, res)
		return
	}

	str := "https://libraryz.herokuapp.com/verify?token=" + token
	ok := email.SendWelcomeEmail(NewUser.Email, NewUser.Fname+NewUser.Lname, str)
	if !ok {
		res.Error = "error at contacting given email"
		utility.SendResponse(w, http.StatusBadRequest, res)
	}

	res.Success = true
	res.Body = map[string]interface{}{
		"token": token,
	}
	utility.SendResponse(w, 200, res)
}

//VerifyEmail func
func VerifyEmail(w http.ResponseWriter, r *http.Request) {
	tokenStr := r.FormValue("token")
	var res utility.Result

	if len(tokenStr) < 1 {
		res.Error = "token not found"
		utility.SendResponse(w, http.StatusBadRequest, res)
		return
	}

	NewUser, e := jwtauth.ExtractClaims(tokenStr)
	if !e {
		res.Error = "error while extracting values form token"
		utility.SendResponse(w, http.StatusBadRequest, res)
		return
	}

	if e := database.AddMember(NewUser); e != nil {
		res.Error = fmt.Sprintf("%s", e)
		utility.SendResponse(w, http.StatusExpectationFailed, res)
		return
	}

	Newurl := "https://anubhavitis.github.io/Library"

	html := `<html> <script> 
	window.location.href ="` + Newurl + `";
	localStorage.setItem("token", ` + tokenStr + `);
	 </script> </html>`
	fmt.Fprintln(w, html)
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

//Welcome handler
func Welcome(w http.ResponseWriter, r *http.Request) {
	utility.SendResponse(w, http.StatusAccepted, &utility.Result{Success: true})
	return
}
