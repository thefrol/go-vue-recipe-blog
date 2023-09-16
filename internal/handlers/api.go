package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/thefrol/go-vue-recipe-blog/internal/data"
	"github.com/thefrol/go-vue-recipe-blog/internal/localstorage"
)

type RecipesResponse struct {
	Recipes []data.Recipe `json:"recipes"`
}

const (
	storageFolder = "../web/.storage/"
)

var store = localstorage.New(storageFolder)

// my token "123lasudhjnqwoealskndlajwjelijqwe" my pass "mypass"

func Recipes(w http.ResponseWriter, r *http.Request) {
	// TODO
	// нейросеть считает каллорийность этих блюд по рецепту))

	recipes, err := store.Recipes()
	if err != nil {
		// TODO
		// хелпер такого вида
		// Respond(w, Code, msg)
		http.Error(w, "Не могу получить рецепты из хранилища;"+err.Error(), http.StatusInternalServerError)
		return
	}

	response := RecipesResponse{Recipes: recipes}
	bb, err := json.Marshal(response)
	if err != nil {
		fmt.Printf("Cant marshal a json with recipes: %+v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "appliation/json")
	w.Write([]byte(bb))
}

func GetRecipe(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	recipe, err := store.Recipe(id)
	if err != nil {
		// TODO
		// хелпер такого вида
		// Respond(w, Code, msg)
		http.Error(w, "Не могу получить рецепт из хранилища;"+err.Error(), http.StatusInternalServerError)
		return
	}

	//TODO это тоже можно выделить в функцию!
	response := recipe
	bb, err := json.Marshal(response)
	if err != nil {
		fmt.Printf("Cant marshal a json with recipes: %+v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "appliation/json")
	w.Write([]byte(bb))
}

func PostRecipe(w http.ResponseWriter, r *http.Request) { // todo, а что если есть какая-то структура с функциями или мапа! Типа поторая сразу содержит get, post, delete. уже похоже на rpc
	w.Header().Add("Content-Type", "appliation/json")
}
