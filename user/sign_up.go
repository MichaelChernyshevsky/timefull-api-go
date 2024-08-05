package user

import (
	"encoding/json"
	"net/http"
	"timefull-api-go/db"
)

type RequestSignUp struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SignUpHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var request RequestSignUp
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	newUserModel := UserModel{
		Email:    request.Email,
		Password: request.Password,
		Phone:    "",
		// Packages:   1,
		Admin:      false,
		Creator:    false,
		Subscribed: true,
		Sex:        "",
		Name:       "",
		Name2:      "",
		Age:        30,
	}

	result := db.DB.Create(newUserModel)
	if result.Error != nil {
		// return nil, result.Error
	}
	// return &newUserModel, nil
}
