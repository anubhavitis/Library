package db

import (
	"database/sql"
	"fmt"
)

//InitDb funtion
func InitDb() (*sql.DB, error) {

	dab, err := sql.Open("mysql", "sql12349917:VEDK9mPCkq@(sql12.freemysqlhosting.net)/sql12349917?parseTime=true")
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

//CreateTable function
func CreateTable() {
	queryUsers := `CREATE TABLE users (
		id INT PRIMARY KEY AUTO_INCREMENT,
		fname varchar(255),
		lname varchar(255),
		created_at timestamp,
		email email,
		college varchar(255),
	  );
		CREATE INDEX users_index_0 ON users (email);
		CREATE INDEX users_index_1 ON users (college);`

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

	fmt.Println(queryUsers, queryBooks)
}
