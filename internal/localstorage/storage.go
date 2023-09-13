package localstorage

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/thefrol/go-vue-recipe-blog/internal/data"
)

// Storager дает доступ до основых сущеностей так,
// чтобы мы могли не париться над тем, как они там лежат
// мы можем получить нужные рецепты и установить токены
// проверить логин и пароль

type Storager interface {
	Recipe(id string) (*data.Recipe, error)
	SetRecipe(id string, r data.Recipe)
	Recipes() ([]*data.Recipe, error)

	Token(id string) bool
	SetToken(id string)

	Password(login string) (checksum string, err error)
	SetCredentials(login, pass string)
}

// Хранилище данных, реализовано в файлах
// в папке .storage
//
// - Токены хранятся захешированно в файле tokens.
// - Пароли хранятся в файлах вида <login>.passhash,
// 		где внутри файла хеш пароля
// - Рецепты лежат в папке recipe

type FileStorage struct {
	folder string
}

func New(storageFolder string) FileStorage {
	if _, err := os.Stat(storageFolder); errors.Is(err, os.ErrNotExist) {
		// такого пути нет
		panic("Неправильно создано хранилище, такого пути нет")
	}
	return FileStorage{folder: storageFolder}
}

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

func (s FileStorage) Token(id string) bool {
	panic("not implemented") // TODO: Implement
}

func (s FileStorage) SetToken(id string) {
	panic("not implemented") // TODO: Implement
}

func (s FileStorage) Password(login string) (checksum string, err error) {
	panic("not implemented") // TODO: Implement
}

func (s FileStorage) SetCredentials(login string, pass string) {
	panic("not implemented") // TODO: Implement
}

const recipeFolderName = "recipe"

func (s FileStorage) recipeFolder() string {
	return path.Join(s.folder, recipeFolderName)
}

func hash(value string) []byte {
	h := sha256.Sum256([]byte(value))
	return h[:]
}
