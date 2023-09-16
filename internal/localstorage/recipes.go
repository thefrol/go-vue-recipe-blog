package localstorage

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/thefrol/go-vue-recipe-blog/internal/data"
)

func (s FileStorage) Recipe(id string) (*data.Recipe, error) {
	// TODO
	// на данный момент он выдает рецепт с постфиксом .json, c этим надо разобраться, и не в роутинге
	bb, err := os.ReadFile(path.Join(s.recipeFolder(), id))
	if err != nil {
		return nil, fmt.Errorf("Cant read recipe with id %v: %+v", id, err)
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
	os.WriteFile(path.Join(s.recipeFolder(), id), json, os.FileMode(os.O_WRONLY|os.O_CREATE))
	return nil

	// TODO
	// 1. os.FileMode(os.O_WRONLY|os.O_CREATE) в константу и посмотреть где ещё использовать можно
	//
	// 2. recipeFolder я думаю нам не нужна в таком виде, она скорее должна быть как recipePath(recipeId string)
	//
	// 3. У меня даже в файлах кредентиалс выделен в отдельный блок, как бы я уже подсознательно хочу это все разделить

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
