package localstorage

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/thefrol/go-vue-recipe-blog/internal/data"
)

func (s FileStorage) Recipe(id string) (*data.Recipe, error) {
	bb, err := os.ReadFile(path.Join(s.recipeFolder(), id))
	if err != nil {
		return nil, fmt.Errorf("Cant read recipe with id %v: %+v", id, err)
	}

	recipe := new(data.Recipe)
	json.Unmarshal(bb, recipe)
	return recipe, nil
}

func (s FileStorage) SetRecipe(id string, r data.Recipe) {
	panic("not implemented") // TODO: Implement
}

func (s FileStorage) Recipes() (recipes []data.Recipe, err error) {
	files, err := os.ReadDir(s.recipeFolder())
	if err != nil {
		return nil, fmt.Errorf("Cant read recipes folder: %+v", err)
	}

	for _, id := range files {
		r, err := s.Recipe(id.Name())
		if err != nil {
			return nil, err
		}
		recipes = append(recipes, *r)
	}
	return
}

func (s FileStorage) recipeFolder() string {
	return path.Join(s.folder, recipeFolderName)
}
