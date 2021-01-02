package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	// DB "Library/APIs/databases/database"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type person struct {
	Fname string `json:"Fname"`
	Lname string `json:"Lname"`
}

type user []person

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

func allUsers(w http.ResponseWriter, r *http.Request) {
	users := user{
		person{Fname: "Anubhav", Lname: "Singhal"},
		person{Fname: "Prachi", Lname: "Singhal"},
	}

	html := `<html> <body> 
				<div><a href="\"> Home</a></div>
				<div><a href="\page1"> Page 1</a></div>
				<div> <h5>`
	html2 := `</h5></div> </body> </html>`

	fmt.Fprintln(w, html)
	json.NewEncoder(w).Encode(users)
	fmt.Fprintln(w, html2)
}

//Homepage handler
func Homepage(w http.ResponseWriter, r *http.Request) {
	html := `<html> <body> 
				<a href="\page1"> Page 1</a>
			</body> </html>`

	fmt.Fprintln(w, html)
}

//GotoP1 handler
func GotoP1(w http.ResponseWriter, r *http.Request) {
	html := `<html> <body> 
				<div> <h1> This is Page1</h1></div>
				<div><a href="\"> Home</a></div>
				<div><a href="\users"> Show users</a></div>
			</body> </html>`

	fmt.Fprintln(w, html)
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
	r.HandleFunc("/page1", GotoP1)
	r.HandleFunc("/users", allUsers)

	http.ListenAndServe(":8080", r)
}
