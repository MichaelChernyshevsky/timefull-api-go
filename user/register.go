package user

import "net/http"

func RegisterApi() {
	http.HandleFunc("/user", Handler)
	http.HandleFunc("/signin", SignInHandler)
	http.HandleFunc("/signup", SignUpHandler)
	http.HandleFunc("/signout", SignOutHandler)
	http.HandleFunc("/delete", DeleteHandler)

}
