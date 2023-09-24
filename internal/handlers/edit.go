package handlers

import (
	"net/http"
)

func Edit(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../web/edit.html")
}
