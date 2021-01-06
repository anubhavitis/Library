package main

import (
	"database/sql"
	"net/http"

	// DB "Library/APIs/databases/database"

	v1 "github.com/anubhavitis/Library/apis/v1"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

//Mydb pointer to database
var Mydb *sql.DB
var port = "8000"

func init() {
	// Mydb, err := DB.InitDb()
	// if err != nil {
	// 	fmt.Println("DB: ", Mydb)
	// }
	// if e := DB.CreateBooksTable(Mydb); e != nil {
	// 	fmt.Println("Error at creating books:", e)
	// }
	// if e := DB.CreateMemberTable(Mydb); e != nil {
	// 	fmt.Println("Error at creating users:", e)
	// }
}

func main() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", v1.Homepage)

	//Signup routes
	r.HandleFunc("/signup", v1.SignUp)
	r.HandleFunc("/google/signup", v1.GoogleSignupHandler)
	r.HandleFunc("/google/callback", v1.GoogleCallbackHandler)
	r.HandleFunc("/twitter/signup", v1.TwitterSignupHandler)
	r.HandleFunc("/twitter/callback", v1.TwitterCallbackHandler)

	r.HandleFunc("/signin", v1.SignIn)
	r.HandleFunc("/welcome", v1.Welcome)
	r.HandleFunc("/refresh", v1.Refresh)
	http.Handle("/", r)
	http.ListenAndServe(":"+port, nil)
}
