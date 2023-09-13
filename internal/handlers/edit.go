package handlers

import (
	"context"
	"fmt"
	"net/http"
)

const (
	pinQueryParam    = "pin"
	authContextParam = "auth"
	pinValue         = "1111"
	cookieToken      = "123makdjo8y12jn"
)

func Edit(w http.ResponseWriter, r *http.Request) {
	ii := r.Context().Value(authContextParam)
	var auth string
	var ok bool
	if auth, ok = ii.(string); !ok {
		fmt.Println("Edit handler not getting auth")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write([]byte(auth))
}

func Authorization(next http.Handler) http.Handler {
	// #TODO
	// Как-то нужно для мидлвари придуать отдельно место,
	// Но пока тут все завязано на общих константах я не хочу это делать
	// ещё бы как-то разделить вебхендлеры, и апи хендлеры
	// возможно веб лежат в сервере, а апи в интернал
	// но тогда странно что они в разных местах
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pin := r.URL.Query().Get(pinQueryParam)
		fmt.Println("auth: pin is ", pin)

		if pin != pinValue {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		rc := r.Context()
		ctx := context.WithValue(rc, authContextParam, "by pin")
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
