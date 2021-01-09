package v1

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

func sendResponse(w http.ResponseWriter, code int, data interface{}) bool {
	w.WriteHeader(code)
	res, err := json.Marshal(data)
	if err == nil {
		w.Write(res)
		return true
	}
	return false
}

//GenerateUUID func
func GenerateUUID() string {

	sd := uuid.New()
	return (sd.String())

}
