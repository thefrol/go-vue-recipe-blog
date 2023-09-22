package recipes

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/thefrol/go-vue-recipe-blog/internal/data"
)

var RecipeNotExist = errors.New("recipe not exist")

const recipeFileExtension = ".json"

// Recipe возвращает рецепт с идентификатором id, такой рецепт лежит в файл <id>.json
//
//	okroshka -> okroshka.json
func (s FileStorage) Recipe(id string) (*data.Recipe, error) {
	bb, err := os.ReadFile(path.Join(s.folder, id+recipeFileExtension))
	if os.IsNotExist(err) {
		return nil, RecipeNotExist
	} else if err != nil {

		return nil, fmt.Errorf("сant read recipe with id %v: %+v", id, err)
	}

	recipe := new(data.Recipe)
	json.Unmarshal(bb, recipe)
	return recipe, nil
}

func (s FileStorage) SetRecipe(id string, r data.Recipe) error {
	json, err := json.Marshal(&r)
	if err != nil {
		return err
	}
	os.WriteFile(path.Join(s.folder, id+recipeFileExtension), json, os.FileMode(os.O_WRONLY|os.O_CREATE))
	return nil
}

// Recipes возвращет все рецепты из папки, это должны быть файлы с расширением .json
func (s FileStorage) Recipes() (recipes []data.Recipe, err error) {
	files, err := os.ReadDir(s.folder)
	if err != nil {
		return nil, fmt.Errorf("сant read recipes folder: %+v", err)
	}

	for _, f := range files {
		ext := path.Ext(f.Name())
		if f.IsDir() || ext != recipeFileExtension {
			fmt.Println("some annoying file in recipe folder")
			continue
		}
		id := strings.TrimRight(f.Name(), recipeFileExtension)
		r, err := s.Recipe(id)
		if err != nil {
			return nil, err
		}
		recipes = append(recipes, *r)
	}
	return
}
