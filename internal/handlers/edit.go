package handlers

import (
	"fmt"
	"net/http"
)

// TODO
// Далее хранить где-то токен	....
const authContextParam = "auth"

func Edit(w http.ResponseWriter, r *http.Request) {
	ii := r.Context().Value(authContextParam)
	var auth string
	var ok bool
	if auth, ok = ii.(string); !ok || auth != "ok" {
		fmt.Println("in edit handler not found auth, big error in logic")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Вы авторизированы. Привет! Могу сказать сколько ещё будет жить токен"))
}
