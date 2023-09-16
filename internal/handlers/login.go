package handlers

import (
	"bytes"
	"net/http"
	"time"

	"github.com/thefrol/go-vue-recipe-blog/internal/utils"
)

const (
	cookieName     = "accessToken"
	cookieToken    = "ro8BS6Hiivgzy8Xuu09JDjlNLnSLldY5"
	cookieLifeDays = 7
)

func Authorize(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // Это очень важная часть, без этого даже не прочитается
	login := r.FormValue("username")
	pass := r.FormValue("password")
	hash, err := store.Password(login)
	if err != nil {
		http.Error(w, "Cant authorize", http.StatusInternalServerError) // не пали контору
		return
	}

	if bytes.Compare(utils.Hash(pass), hash) != 0 {
		http.Error(w, "Wrong login/pass", http.StatusUnauthorized)
		return
	}

	http.SetCookie(w, makeCookie())

	// TODO
	// Надо чтобы он ещё запоминал, куда шел пользователь и выбрасывал на нужную страницу
	http.Redirect(w, r, "/edit/22", http.StatusFound) // todo что делает StatusFound - переделывает в GET запрос?
}

func makeCookie() *http.Cookie {
	//store cookie
	t := utils.UUID()
	store.AddToken(utils.Hash(t))

	//set cookie
	cookie := http.Cookie{}
	cookie.Name = cookieName
	cookie.Value = t
	cookie.Expires = time.Now().Add(cookieLifeDays * 24 * time.Hour)
	cookie.Secure = false
	cookie.HttpOnly = true
	cookie.Path = "/"

	return &cookie
}
