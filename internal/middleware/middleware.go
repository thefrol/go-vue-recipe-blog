package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/thefrol/go-vue-recipe-blog/internal/localstorage"
	"github.com/thefrol/go-vue-recipe-blog/internal/utils"
)

const (
	cookieName     = "accessToken"
	cookieToken    = "ro8BS6Hiivgzy8Xuu09JDjlNLnSLldY5"
	cookieLifeDays = 7

	authContextParam = "auth"
	// TODO
	// где же эта константа должна храниться,
	// может передаваться в функции???
	// и хранилище тоже например
)

const (
	storageFolder = "../web/.storage/"
)

// сделать дефолтное хранилище TODO
// вообще этот прикол, что мы не может хранить storageFolder в каком-то одном пакете намекает, что для рецептов нужно отдальное хранилще
// Даже гитигнор как-то не получается адекватно написать под эти задачи
// рецепты лежат среди очень чувствительного хранилища, тут как- бы в будущем не запутаться
var store = localstorage.New(storageFolder)

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
			fmt.Printf("Autorization required at %v not getting auth\n", r.URL.Path)
			http.Redirect(w, r, "/login", http.StatusFound)

			// return
		}
		if auth != "ok" {
			http.Redirect(w, r, "/login", http.StatusFound)
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
