package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path"

	"github.com/thefrol/go-vue-recipe-blog/internal/data"
)

type RecipesResponse struct {
	Recipes []data.Recipe `json:"recipes"`
}

const (
	recipesFolder = "../web/recipe/"
)

func Recipes(w http.ResponseWriter, r *http.Request) {
	files, err := os.ReadDir(recipesFolder)
	if err != nil {
		fmt.Printf("Cant read recipes folder: %+v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var recipes []data.Recipe

	for _, f := range files {
		if f.IsDir() {
			continue
		}

		bb, err := os.ReadFile(path.Join(recipesFolder, f.Name()))
		if err != nil {
			fmt.Printf("Cant read recipes file %v: %+v", f.Name(), err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		recipe := data.Recipe{}
		json.Unmarshal(bb, &recipe)
		recipes = append(recipes, recipe)
	}

	// TODO
	// нейросеть считает каллорийность этих блюд по рецепту))

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
