package admin

import (
	"encoding/json"
	"net/http"
)

type RequestInfo struct {
	UserModelname string `json:"email"`
	Password      string `json:"password"`
}

func info(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var request RequestInfo
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
}
