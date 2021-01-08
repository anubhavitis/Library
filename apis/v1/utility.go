package v1

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

func sendResponse(w http.ResponseWriter, code int, data interface{}) bool {
	w.WriteHeader(code)
	res, err := json.Marshal(data)
	if err == nil {
		w.Write(res)
		return true
	}
	return false
}

func GenerateUUID() string {

	sd := uuid.New()
	return (sd.String())

}
