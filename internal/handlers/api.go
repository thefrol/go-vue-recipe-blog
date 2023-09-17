package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/thefrol/go-vue-recipe-blog/internal/data"
	"github.com/thefrol/go-vue-recipe-blog/internal/recipes"
)

type RecipesResponse struct {
	Recipes []data.Recipe `json:"recipes"`
}

const (
	storageFolder = "../assets/recipes/" //todo: отправить в server
)

var store = recipes.New(storageFolder)

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
		fmt.Printf("Cant marshal a json with %v recipe: %+v", id, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "appliation/json")
	w.Write([]byte(bb))
}

func PostRecipe(w http.ResponseWriter, r *http.Request) { // todo, а что если есть какая-то структура с функциями или мапа! Типа поторая сразу содержит get, post, delete. уже похоже на rpc
	id := chi.URLParam(r, "id")

	if r.Body == nil {
		http.Error(w, "пришло пустое тело", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body) // todo посмотреть как это делали в практикуме
	if err != nil {
		// TODO
		// хелпер такого вида
		// Respond(w, Code, msg)
		http.Error(w, "Не могу получить рецепт из хранилища;"+err.Error(), http.StatusInternalServerError)
		return
	}

	recipe := new(data.Recipe)
	err = json.Unmarshal(body, recipe)
	if err != nil {

		//TODO это тоже можно выделить в функцию!
		fmt.Printf("Cant unmarshal a json with %v recipe: %+v", id, err)
		http.Error(w, "Cant unmarshal a json", http.StatusBadRequest)
		return
	}

	err = store.SetRecipe(id, *recipe)
	if err != nil {
		fmt.Printf("Cant save %v recipe: %+v", id, err)
		http.Error(w, "Cant save a recipe", http.StatusBadRequest)
	}

	w.Header().Add("Content-Type", "appliation/json")
}
