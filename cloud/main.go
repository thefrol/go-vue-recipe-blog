// Based on https://github.com/zooraze/chi-vue-spa
// and https://github.com/thefrol/go-chi-yandex-cloud-template

package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/thefrol/go-vue-recipe-blog/internal/handlers"
	"github.com/thefrol/go-vue-recipe-blog/internal/middleware"
)

var Router = chi.NewRouter()

func init() {
	Router.Use(middleware.PinAuthorization, middleware.CookieAuthorization)

	Router.Route("/", files)

	// admining
	Router.Get("/login", Login)
	Router.Post("/login", handlers.Authorize)
	Router.Route("/edit", func(r chi.Router) {
		// TODO
		// мне видятся разные пакеты для апи и веба, а так же для мидлвари тоже
		// возможно и отдельные переменные, как минимум роутер в отдельный файл
		// admin/web/api

		//Требуется авторизация
		r.Use(middleware.RequireAuthorization)
		r.Get("/*", handlers.Edit)
	})

	// api
	Router.Route("/api/v1/", func(r chi.Router) {
		r.Get("/recipes", handlers.Recipes)
	})
}

func main() {
	http.ListenAndServe(":8080", Router)
}

func files(r chi.Router) {
	fileServer := http.FileServer(http.Dir("../web"))
	r.Handle("/", fileServer) // только index.html доступен извне в корневом каталоге
	r.Handle("/css/{file}", fileServer)
	r.Handle("/script/{file}", fileServer)
}

func Login(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../web/login.html")
}
