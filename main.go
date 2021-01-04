package main

import (
	"database/sql"
	"fmt"
	"net/http"

	// DB "Library/APIs/databases/database"

	"github.com/anubhavitis/Library/apis"
	DB "github.com/anubhavitis/Library/databases"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

//Mydb pointer to database
var Mydb *sql.DB

func main() {

	Mydb, err := DB.InitDb()
	if err != nil {
		fmt.Println("DB: ", Mydb)
	}
	if e := DB.CreateBooksTable(Mydb); e != nil {
		fmt.Println("Error at creating books:", e)
	}
	if e := DB.CreateMemberTable(Mydb); e != nil {
		fmt.Println("Error at creating users:", e)
	}

	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", apis.Homepage)
	r.HandleFunc("/signin", apis.SignIn)
	r.HandleFunc("/signup", apis.SignUp)
	r.HandleFunc("/welcome", apis.Welcome)
	r.HandleFunc("/refresh", apis.Refresh)

	http.ListenAndServe(":8080", r)
}
