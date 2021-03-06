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
	Error   string                 `json:"error"`
}

//SendResponse func
func SendResponse(w http.ResponseWriter, code int, data interface{}) bool {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	res, err := json.Marshal(data)
	if err == nil {
		w.Write(res)
		return true
	}
	return false
}

//UserCred as expected User Credential while login
type UserCred struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

//GenerateUUID func
func GenerateUUID() string {

	sd := uuid.New()
	return (sd.String())

}
