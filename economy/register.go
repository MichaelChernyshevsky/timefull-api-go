package economy

import "net/http"

func RegisterApi() {
	http.HandleFunc("/economy", Handler)

}
