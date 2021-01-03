package database

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

//Member Structure
type Member struct {
	UID     string
	Fname   string
	Lname   string
	Time    string
	Email   string
	College string
}

//Book strucuture
type Book struct {
	UID         string
	Name        string
	Author      string
	Genre       string
	Description string
}

//GenerateUUID ..
func GenerateUUID() string {
	v, _ := uuid.NewUUID()
	return v.String()
}

//InitDb funtion
func InitDb() (*sql.DB, error) {

	dab, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println("Error at opening database")
		return nil, err
	}
	if err := dab.Ping(); err != nil {
		fmt.Println("Error at ping.")
		return nil, err
	}
	return dab, nil
}

//CreateMemberTable function
func CreateMemberTable(db *sql.DB) error {

	if _, err := db.Exec("DROP TABLE users"); err != nil {
		return err
	}

	queryUsers := `
	CREATE TABLE users (
		id INT PRIMARY KEY AUTO_INCREMENT,
		fname varchar(255),
		lname varchar(255),
		created_at timestamp,
		email email,
		college varchar(255),
	);
	CREATE INDEX users_index_0 ON users (email);
	CREATE INDEX users_index_1 ON users (college);
	`

	if _, err := db.Exec(queryUsers); err != nil {
		return err
	}
	return nil
}

//CreateBooksTable function
func CreateBooksTable(db *sql.DB) error {

	if _, err := db.Exec("DROP TABLE books"); err != nil {
		return err
	}

	queryBooks := `CREATE TABLE books (
		id int PRIMARY KEY AUTO_INCREMENT,
		name varchar(255),
		author varchar(255),
		genre varchar(255),
		description varchar(255),
	  );
	  
	  	ALTER TABLE books ADD FOREIGN KEY (author) REFERENCES users (id);
		CREATE INDEX books_index_2 ON books (author);
		CREATE INDEX books_index_3 ON books (genre);`

	if _, err := db.Exec(queryBooks); err != nil {
		return err
	}
	return nil

}
