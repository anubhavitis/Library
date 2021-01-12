package v1

import (
	"encoding/json"
	"errors"
	"net/http"

	auth "github.com/anubhavitis/Library/apis/middleware"
	Db "github.com/anubhavitis/Library/databases"
)

//AddBook func
func AddBook(w http.ResponseWriter, r *http.Request) {
	var resp Result
	var book Db.Book

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		resp.Error = err
		SendResponse(w, http.StatusBadRequest, resp)
	}

	// book.UID = GenerateUUID()
	book.Likes = 0

	if err = Db.AddBook(book); err != nil {
		resp.Error = errors.New("error while adding book to database")
		SendResponse(w, 400, resp)
		return
	}
	resp.Success = true
	SendResponse(w, 200, resp)
}

//DeleteBook func
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	var book Db.Book
	var res Result
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		res.Error = err
		SendResponse(w, http.StatusBadRequest, res)
		return
	}

	if _, e := Db.DeleteBook(book.UID); e != nil {
		res.Error = e
		SendResponse(w, http.StatusConflict, res)
		return
	}

	res.Success = true
	SendResponse(w, http.StatusAccepted, res)
}

//UpdateBookInfo func
func UpdateBookInfo(w http.ResponseWriter, r *http.Request) {
	var book Db.Book
	var res Result
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		res.Error = err
		SendResponse(w, http.StatusBadRequest, res)
		return
	}

	if _, e := Db.UpdateBook(book); e != nil {
		res.Error = e
		SendResponse(w, http.StatusConflict, res)
		return
	}

	res.Success = true
	SendResponse(w, http.StatusAccepted, res)
}

//GetBook func
func GetBook(w http.ResponseWriter, r *http.Request) {
	var res Result
	var user auth.UserCred

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		res.Error = err
		SendResponse(w, http.StatusBadRequest, res)
		return
	}

	books, err := Db.ListUserBooks(user.Username)
	if err != nil {
		res.Error = err
		SendResponse(w, http.StatusConflict, res)
		return
	}
	res.Body = map[string]interface{}{
		"books": books,
	}
	res.Success = true
	SendResponse(w, http.StatusAccepted, res)
}
