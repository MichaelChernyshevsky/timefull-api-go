package task

import "net/http"

func RegisterApi() {
	http.HandleFunc("/task", Handler)

}
