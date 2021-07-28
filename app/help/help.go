package help

import (
	"encoding/json"
	"net/http"
)

func ResponseWithError(w http.ResponseWriter, code int, msg string) {
	ResponseWithJson(w, code, map[string]string{"error": msg})
}

func ResponseWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
