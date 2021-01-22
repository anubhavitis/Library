package main

import (
	"fmt"
	"net/http"
	"os"

	v1 "github.com/anubhavitis/Library/apis/v1"
	database "github.com/anubhavitis/Library/databases"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	Mydb, err := database.InitDb()
	if err != nil {
		fmt.Println("DB: ", Mydb)
	}
	database.Mydb = Mydb

	r := mux.NewRouter().StrictSlash(true)
	headers := handlers.AllowedHeaders([]string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"})
	methods := handlers.AllowedMethods([]string{"POST", "GET", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})
	r.HandleFunc("/", v1.Welcome)

	r.HandleFunc("/signup", v1.SignUp)
	r.HandleFunc("/googlesignup", v1.GoogleSignupHandler)
	r.HandleFunc("/google_callback", v1.GoogleCallbackHandler)
	r.HandleFunc("/verify", v1.VerifyEmail)
	r.HandleFunc("/google/signup", v1.GoogleSignupHandler)
	r.HandleFunc("/google/callback", v1.GoogleCallbackHandler)
	r.HandleFunc("/twitter/signup", v1.TwitterSignupHandler)
	r.HandleFunc("/twitter/callback", v1.TwitterCallbackHandler)

	r.HandleFunc("/signin", v1.SignIn)
	r.HandleFunc("/refresh", v1.Refresh)
	r.HandleFunc("/welcome", v1.Welcome)
	r.HandleFunc("/addbook", v1.AddBook)
	r.HandleFunc("/deletebook", v1.DeleteBook)
	r.HandleFunc("/updatebook", v1.UpdateBookInfo)
	r.HandleFunc("/getallbook", v1.GetBook)
	http.Handle("/", handlers.CORS(headers, methods, origins)(r))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	http.ListenAndServe(":"+port, nil)
}
