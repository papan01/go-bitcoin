package modules

import (
	"net/http"
)

//主頁面:"/"
func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	http.ServeFile(w, r, "./static/index.html")
}
