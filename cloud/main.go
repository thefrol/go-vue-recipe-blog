// Based on https://github.com/zooraze/chi-vue-spa
// and https://github.com/thefrol/go-chi-yandex-cloud-template

package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/thefrol/go-vue-recipe-blog/internal/handlers"
)

var Router = chi.NewRouter()

func init() {
	fileServer := http.FileServer(http.Dir("../web"))
	Router.Handle("/*", fileServer)
	// admining
	Router.Route("/edit", func(r chi.Router) {
		r.Use(handlers.Authorization)
		r.Get("/*", http.NotFound)
	})

	Router.Get("/login", http.NotFound)

	// api
	Router.Route("/api/v1/", func(r chi.Router) {
		// поупражняться в gRPC
		r.Get("/recipes", handlers.RecipesHandler)
	})
}

func main() {
	http.ListenAndServe(":8080", Router)
}
