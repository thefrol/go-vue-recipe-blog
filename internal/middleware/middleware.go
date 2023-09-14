package middleware

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/thefrol/go-vue-recipe-blog/internal/localstorage"
	"github.com/thefrol/go-vue-recipe-blog/internal/utils"
)

const (
	pinQueryParam    = "pin"
	authContextParam = "auth"
	pinValue         = "1111"
)

const (
	cookieName     = "accessToken"
	cookieToken    = "ro8BS6Hiivgzy8Xuu09JDjlNLnSLldY5"
	cookieLifeDays = 7
)

const (
	storageFolder = "../web/.storage/"
)

// сделать дефолтное хранилище TODO
// вообще этот прикол, что мы не может хранить storageFolder в каком-то одном пакете намекает, что для рецептов нужно отдальное хранилще
var store = localstorage.New(storageFolder)

func PinAuthorization(next http.Handler) http.Handler {
	// #TODO
	// Как-то нужно для мидлвари придуать отдельно место,
	// Но пока тут все завязано на общих константах я не хочу это делать
	// ещё бы как-то разделить вебхендлеры, и апи хендлеры
	// возможно веб лежат в сервере, а апи в интернал
	// но тогда странно что они в разных местах
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pin := r.URL.Query().Get(pinQueryParam)
		fmt.Println("auth: pin is ", pin)

		// Только устанавливаем контекст

		if pin == pinValue {

			rc := r.Context()
			ctx := context.WithValue(rc, authContextParam, "ok")
			r = r.WithContext(ctx)

			// TODO
			// Выделить это в отдельную функцию

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
			http.SetCookie(w, &cookie)

		}

		//так же говорим, как прошла авторизация, можно без этого обойтись в общем-то, хотя нет, следующие мидлвари должны знать что авторизация успешна

		next.ServeHTTP(w, r)
	})
}

func CookieAuthorization(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if findCookie(r) {
			// TODO
			// тоже выделить во что-то отдельное
			rc := r.Context()
			ctx := context.WithValue(rc, authContextParam, "ok")
			r = r.WithContext(ctx)
			fmt.Println("Authorized by cookie")
		}
		next.ServeHTTP(w, r)
	})
}

func RequireAuthorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ii := r.Context().Value(authContextParam)
		var auth string
		var ok bool
		if auth, ok = ii.(string); !ok {
			fmt.Printf("Autorization required at %v not getting auth=ok\n", r.URL.Path)
			// w.WriteHeader(http.StatusInternalServerError)
			// return
		}
		if auth != "ok" {
			http.Error(w, "Not authorized", http.StatusUnauthorized)
			return
		}
		fmt.Println("AUTH OK")
		next.ServeHTTP(w, r)
	})
}

// findCookie проверяет есть ли нужный куки в запросе и сверяет его хеш,
// с уже хранимыми хешами, если есть совпадение вернет true первым параметром
func findCookie(r *http.Request) bool {
	c, err := r.Cookie(cookieName)
	if err != nil {
		fmt.Printf("Cookie %v not found in cookies \n", c)
		return false
	}
	found, err := store.Token(utils.Hash(c.Value))
	if err != nil {
		fmt.Printf("Token not found: %v\n", c)
		return false
	}
	return found
}
