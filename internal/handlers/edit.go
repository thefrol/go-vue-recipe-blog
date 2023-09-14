package handlers

import (
	"net/http"
)

func Edit(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Вы авторизированы. Привет! Могу сказать сколько ещё будет жить токен"))
}
