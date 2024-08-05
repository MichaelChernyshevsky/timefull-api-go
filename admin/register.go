package admin

import "net/http"

func RegisterApi() {
	http.HandleFunc("/admin/info", info)
	
}
