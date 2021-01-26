package database

import (
	"database/sql"
	"fmt"
)

//Book strucuture
type Book struct {
	UID    string `json:"uid"`
	Owner  string `json:"owner"`
	Name   string `json:"name"`
	Author string `json:"author"`
	Genre  string `json:"genre"`
	About  string `json:"about"`
	Likes  int    `json:"likes"`
	Image  string `json:"image"`
}

//CreateBooksTable function
func CreateBooksTable(db *sql.DB) error {
	queryBooks := `
	CREATE TABLE IF NOT EXISTS
	books (
		id INT PRIMARY KEY AUTO_INCREMENT,
		name varchar(255),
		owner varchar(255),
		author varchar(255),
		genre varchar(255),
		about varchar(500),
		likes INT DEFAULT 0,
		image varchar(500) DEFAULT "https://cutt.ly/AjJ7pCN"
	);
	`

	if _, err := db.Exec(queryBooks); err != nil {
		fmt.Println("Books can not be created!")
		return err
	}

	// ShowBooks(db)
	fmt.Println("Books Created!")
	return nil
}

//AddBook function
func AddBook(book Book) error {

	q := `
	INSERT INTO books
	(name, owner, author, genre, about, likes, image)
	Values (?,?,?,?,?,?,?)
	`
	if book.Image == "" {
		book.Image = "https://cutt.ly/AjJ7pCN"
	}
	if _, e := Mydb.Exec(q, book.Name, book.Owner, book.Author, book.Genre, book.About, book.Likes, book.Image); e != nil {
		return e
	}

	return nil
}

//ListUserBooks finds user with particular email, and returns it
func ListUserBooks(uname string) ([]Book, error) {
	var books []Book

	query := `Select * from books where owner=?`
	res, e := Mydb.Query(query, uname)

	if e != nil {
		return books, e
	}

	defer res.Close()

	for res.Next() {
		var t Book
		if err := res.Scan(&t.UID, &t.Name, &t.Owner, &t.Author, &t.Genre, &t.About, &t.Likes, &t.Image); err != nil {
			return books, nil
		}
		books = append(books, t)
	}

	return books, nil
}

//GetAllBook func
func GetAllBook() ([]Book, error) {
	var books []Book

	query := `Select * from books`
	res, e := Mydb.Query(query)

	if e != nil {
		return books, e
	}

	defer res.Close()

	for res.Next() {
		var t Book
		if err := res.Scan(&t.UID, &t.Name, &t.Owner, &t.Author, &t.Genre, &t.About, &t.Likes, &t.Image); err != nil {
			return books, nil
		}
		books = append(books, t)
	}

	return books, nil
}

//DeleteBook func
func DeleteBook(id string) (bool, error) {
	q := `DELETE from books WHERE id=?`

	if _, e := Mydb.Exec(q, id); e != nil {
		return false, e
	}
	return true, nil
}

//UpdateBook func
func UpdateBook(book Book) (bool, error) {
	q := ` UPDATE books
		SET name=?, owner=?, genre=?, about=?, likes=?, image=?
		where uid=?`

	if _, e := Mydb.Exec(q, book.Name, book.Owner, book.Genre, book.About, book.Likes, book.Image, book.UID); e != nil {
		return false, e
	}
	return true, nil
}

//ShowBooks temp func
func ShowBooks(db *sql.DB) {
	var uname, owner string
	fmt.Println("Available books are: ")
	query := `Select name, owner from books`
	res, e := db.Query(query)

	if e != nil {
		return
	}

	defer res.Close()

	for res.Next() {
		if err := res.Scan(&uname, &owner); err != nil {
			return
		}
		fmt.Println(uname + " " + owner)
	}
	fmt.Println("#####################")
}
