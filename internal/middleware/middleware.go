package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/thefrol/go-vue-recipe-blog/internal/credentials"
)

const (
	cookieName     = "accessToken"
	cookieToken    = "ro8BS6Hiivgzy8Xuu09JDjlNLnSLldY5" // todo удалить
	cookieLifeDays = 7

	authContextParam = "auth"
	// TODO
	// где же эта константа должна храниться,
	// может передаваться в функции???
	// и хранилище тоже например
)

func CookieAuthorization(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if validateCookie(r) {
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

// validateCookie проверяет есть ли нужный куки в запросе и сверяет его хеш,
// с уже хранимыми хешами, если есть совпадение вернет true первым параметром
func validateCookie(r *http.Request) bool {
	c, err := r.Cookie(cookieName)
	if err != nil {
		fmt.Printf("Cookie %v not found in cookies \n", c)
		return false
	}
	return credentials.CheckToken(c.Value)
}
