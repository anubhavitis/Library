package main

import (
	"database/sql"
	"fmt"
	"net/http"

	// DB "Library/APIs/databases/database"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB

//InitDb function
func InitDb() (*sql.DB, error) {
	dab, err := sql.Open("mysql", "anubhavitis:anubhh@v123@tcp(127.0.0.1:3306)/library")
	if err != nil {
		fmt.Println("Error at opening database: ", err)
		return nil, err
	}

	defer dab.Close()
	if err := dab.Ping(); err != nil {
		fmt.Println("Error at ping: ", err)
		return nil, err
	}
	return dab, nil
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

//SignUp handler
func SignUp(w http.ResponseWriter, r *http.Request) {

}

func main() {
	db, err := InitDb()
	if err != nil {
		fmt.Println("DB: ", db)
	} else {
		fmt.Println("DB okay, okay")
	}
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", Homepage)
	r.HandleFunc("/signup", SignUp)

	http.ListenAndServe(":8080", r)
}
