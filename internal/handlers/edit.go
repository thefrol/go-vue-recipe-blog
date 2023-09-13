package handlers

import (
	"context"
	"fmt"
	"net/http"
	"time"
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

// TODO
// Далее хранить где-то токен	....

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

			cookie := http.Cookie{}
			cookie.Name = cookieName
			cookie.Value = cookieToken
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

		c, err := r.Cookie(cookieName)
		if err != nil {
			fmt.Printf("Wrong cookie. Request cookie: %v\n", c)
		} else if cookieToken == c.Value {

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

// IDEA
// ещё одну мидлварь - куки сеттер, если не установлено и ты авторизован

func BlockUnauthorized(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ii := r.Context().Value(authContextParam)
		var auth string
		var ok bool
		if auth, ok = ii.(string); !ok {
			fmt.Println("Edit handler not getting auth=ok")
			// w.WriteHeader(http.StatusInternalServerError)
			// return
		}
		if auth != "ok" {
			http.Error(w, "Not authorized", http.StatusUnauthorized)
			return
		}
		fmt.Print("AUTH OK")
		next.ServeHTTP(w, r)
	})
}
