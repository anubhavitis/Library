package v1

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/anubhavitis/Library/apis/middleware"
	"github.com/anubhavitis/Library/apis/utility"
	Db "github.com/anubhavitis/Library/databases"
)

//AddBook func
func AddBook(w http.ResponseWriter, r *http.Request) {
	var resp utility.Result
	var book Db.Book

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		resp.Error = fmt.Sprintf("%s", err)
		utility.SendResponse(w, http.StatusBadRequest, resp)
	}

	// book.UID = GenerateUUID()
	book.Likes = 0

	if err = Db.AddBook(book); err != nil {
		resp.Error = "error while adding book to database"
		utility.SendResponse(w, 400, resp)
		return
	}
	resp.Success = true
	utility.SendResponse(w, 200, resp)
}

//DeleteBook func
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	var book Db.Book
	var res utility.Result
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		res.Error = fmt.Sprintf("%s", err)
		utility.SendResponse(w, http.StatusBadRequest, res)
		return
	}

	if _, e := Db.DeleteBook(book.UID); e != nil {
		res.Error = fmt.Sprintf("%s", e)
		utility.SendResponse(w, http.StatusConflict, res)
		return
	}

	res.Success = true
	utility.SendResponse(w, http.StatusAccepted, res)
}

//UpdateBookInfo func
func UpdateBookInfo(w http.ResponseWriter, r *http.Request) {
	var book Db.Book
	var res utility.Result
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		res.Error = fmt.Sprintf("%s", err)
		utility.SendResponse(w, http.StatusBadRequest, res)
		return
	}

	if _, e := Db.UpdateBook(book); e != nil {
		res.Error = fmt.Sprintf("%s", e)
		utility.SendResponse(w, http.StatusConflict, res)
		return
	}

	res.Success = true
	utility.SendResponse(w, http.StatusAccepted, res)
}

//GetBook func
func GetBook(w http.ResponseWriter, r *http.Request) {
	var res utility.Result
	var user middleware.UserCred

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		res.Error = fmt.Sprintf("%s", err)
		utility.SendResponse(w, http.StatusBadRequest, res)
		return
	}

	books, err := Db.ListUserBooks(user.Username)
	if err != nil {
		res.Error = fmt.Sprintf("%s", err)
		utility.SendResponse(w, http.StatusConflict, res)
		return
	}
	res.Body = map[string]interface{}{
		"books": books,
	}
	res.Success = true
	utility.SendResponse(w, http.StatusAccepted, res)
}
