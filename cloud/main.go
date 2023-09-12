// Based on https://github.com/zooraze/chi-vue-spa
// and https://github.com/thefrol/go-chi-yandex-cloud-template

package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

var Router = chi.NewRouter()

func init() {
	fileServer := http.FileServer(http.Dir("../web"))
	Router.Handle("/*", fileServer)

	Router.Route("/api/v1/", func(r chi.Router) {
		// поупражняться в gRPC
		r.Get("/recipes", recipesHandler)
	})
}

func recipesHandler(w http.ResponseWriter, r *http.Request) {
	json := `
		{
			"recipes": [{
					"name": "Быстро сырники",
					"text": "баночка обезжиренного йогурта 120г, 1 яйцо, 2 столовые ложки муки.",
					"tags": ["Пароварка"]
				},
				{
					"name": "Банан и яйцо",
					"text": "Один банан и одно яйцо взбить блендером и пожарить на сковородке. Охрененно ещё сверху намазать творогом",
					"tags": ["Сковорода", "Блины"]
				},
				{
					"name": "Быстро сырники3",
					"text": "баночка обезжиренного йогурта 120г, 1 яйцо, 2 столовые ложки муки по нажаттию раскрываются прям тут.",
					"tags": ["Пароварка"]
				},
				{
					"name": "Быстро сырники4",
					"text": "баночка обезжиренного йогурта 120г, 1 яйцо, 2 столовые ложки муки по нажаттию раскрываются прям тут.",
					"tags": ["Пароварка"]
				},
				{
					"name": "Быстро232 сырники5",
					"text": "баночка обезжиренного йогурта 120г, 1 яйцо, 2 столовые ложки муки по нажаттию раскрываются прям тут.",
					"tags": ["Пароварка"]
				}
			]
		}`
	w.Header().Add("Content-Type", "appliation/json")
	w.Write([]byte(json))
}

func main() {
	http.ListenAndServe(":8080", Router)
}
