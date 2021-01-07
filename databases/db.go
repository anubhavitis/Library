package database

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

//Member Structure
type Member struct {
	UID      string `json:"uid"`
	UserName string `json:"username"`
	Fname    string `json:"fname"`
	Lname    string `json:"lname"`
	Time     string `json:"time"`
	Email    string `json:"email"`
	College  string `json:"college"`
	Password string `json:"password"`
}

//GenerateUUID ..
func GenerateUUID() string {
	v, _ := uuid.NewUUID()
	return v.String()
}

//Mydb function
var Mydb *sql.DB

//InitDb funtion
func InitDb() (*sql.DB, error) {

	dab, err := sql.Open("mysql", "anubhavitis:anubhh@v123@tcp(127.0.0.1:3306)/library")
	if err != nil {
		fmt.Println("Error at opening database")
		return nil, err
	}
	if err := dab.Ping(); err != nil {
		fmt.Println("Error at ping.")
		return nil, err
	}
	if e := CreateBooksTable(dab); e != nil {
		fmt.Println("Error at creating books:", e)
	}
	if e := CreateMemberTable(dab); e != nil {
		fmt.Println("Error at creating users:", e)
	}

	return dab, nil
}

//CreateMemberTable function
func CreateMemberTable(db *sql.DB) error {
	queryUsers := `
	CREATE TABLE IF NOT EXISTS
	users (
		id INT PRIMARY KEY AUTO_INCREMENT,
		username varchar(255),
		fname varchar(255),
		lname varchar(255),
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		email varchar(255),
		college varchar(255),
		password varchar(255)
	);
	`

	if _, err := db.Exec(queryUsers); err != nil {
		return err
	}
	fmt.Println("Users Created!")
	return nil
}

//FindEmail finds user with particular email, and returns it
func FindEmail(email string) (Member, error) {
	var User Member

	query := `Select * from users where email=?`
	res, e := Mydb.Query(query, email)

	if e != nil {
		return User, e
	}

	defer res.Close()

	for res.Next() {
		if err := res.Scan(&User.UID, &User.UserName, &User.Fname, &User.Lname, &User.Time, &User.Email, &User.College, &User.Password); err != nil {
			return User, err
		}
	}

	return User, nil
}

//FindUser finds user with particular email, and returns it
func FindUser(uname string) (Member, error) {
	var User Member

	query := `Select * from users where username=?`
	res, e := Mydb.Query(query, uname)

	if e != nil {
		return User, e
	}

	defer res.Close()

	for res.Next() {
		if err := res.Scan(&User.UID, &User.UserName, &User.Fname, &User.Lname, &User.Time, &User.Email, &User.College, &User.Password); err != nil {
			return User, err
		}
	}

	return User, nil
}

//AddMember to add member to database
func AddMember(mem Member) error {
	q := `
	INSERT INTO users
	(username, fname, lname, email, college, password)
	Values (?,?,?,?,?,?)
	`

	if _, e := Mydb.Exec(q, mem.UserName, mem.Fname, mem.Lname, mem.Email, mem.College, mem.Password); e != nil {
		return e
	}

	return nil
}
