package user

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
}

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "User endpoint")
}

func SignOutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "User endpoint")
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "User endpoint")
}
