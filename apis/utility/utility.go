package utility

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

//Result for every api request
type Result struct {
	Success bool                   `json:"success"`
	Body    map[string]interface{} `json:"body"`
	Error   error                  `json:"error"`
}

//SendResponse func
func SendResponse(w http.ResponseWriter, code int, data interface{}) bool {
	w.Header().Set("Content-Type", "applicaton/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.WriteHeader(code)
	res, err := json.Marshal(data)
	if err == nil {
		w.Write(res)
		return true
	}
	return false
}

//GetTokenFromCookie func
func GetTokenFromCookie(r *http.Request) (string, error) {
	c, err := r.Cookie("Token")
	if err != nil {
		return "", err
	}
	return c.Value, err

}

//GenerateUUID func
func GenerateUUID() string {

	sd := uuid.New()
	return (sd.String())

}
