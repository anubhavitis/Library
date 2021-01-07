package main

import (
	"fmt"
	"net/http"

	"github.com/anubhavitis/Library/apis/middleware"
	v1 "github.com/anubhavitis/Library/apis/v1"
	database "github.com/anubhavitis/Library/databases"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var port = "8080"

func main() {

	Mydb, err := database.InitDb()
	if err != nil {
		fmt.Println("DB: ", Mydb)
	}
	database.Mydb = Mydb

	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", middleware.Homepage)

	//Signup routes
	r.HandleFunc("/signup", v1.SignUp)
	r.HandleFunc("/google/signup", v1.GoogleSignupHandler)
	r.HandleFunc("/google/callback", v1.GoogleCallbackHandler)
	r.HandleFunc("/twitter/signup", v1.TwitterSignupHandler)
	r.HandleFunc("/twitter/callback", v1.TwitterCallbackHandler)

	r.HandleFunc("/signin", v1.SignIn)
	r.HandleFunc("/welcome", middleware.Auth(middleware.Homepage))
	// r.HandleFunc("/test", v1.Welcome)
	r.HandleFunc("/refresh", v1.Refresh)
	http.Handle("/", r)
	http.ListenAndServe(":"+port, nil)
}
