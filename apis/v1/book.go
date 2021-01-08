package v1

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	Db "github.com/anubhavitis/Library/databases"
	"github.com/anubhavitis/Library/pkg/models"
)

func AddBook(w http.ResponseWriter, r *http.Request) {
	var resp models.Response
	var book Db.Book
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		resp.Error = "error while marshalling body"
		sendResponse(w, 400, resp)
		return
	}
	err = json.Unmarshal(data, &book)
	if err != nil {
		resp.Error = "error while marshalling book"
		sendResponse(w, 400, resp)
		return
	}
	book.UID = GenerateUUID()
	book.Likes = 0
	err = Db.AddBook(book)
	if err != nil {
		resp.Error = "error while adding book to database"
		sendResponse(w, 400, resp)
		return
	}
	resp.Success = "Book successfully added"
	sendResponse(w, 200, resp)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	return
}

func UpdateBookInfo(w http.ResponseWriter, r *http.Request) {
	return
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	return
}
