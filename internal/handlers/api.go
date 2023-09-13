package handlers

import "github.com/thefrol/go-vue-recipe-blog/internal/data"

type RecipesResponse struct {
	Recipes []data.Recipe `json:"recipes"`
}
