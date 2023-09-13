package localstorage

import (
	"bytes"
	"crypto/sha256"
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

	Token(id string) (bool, error)
	SetToken(id string) error

	Password(login string) (checksum string, err error)
	SetCredentials(login, pass string) error
}

// Хранилище данных, реализовано в файлах
// в папке .storage
//
//   - Токены хранятся захешированно в файле tokens.
//   - Пароли хранятся в файлах вида <login>.passhash,
//     где внутри файла хеш пароля
//   - Рецепты лежат в папке recipe
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

const tokenFile = "tokens"

func (s FileStorage) Token(token string) (bool, error) {
	println(s.tokensFile())
	bb, err := os.ReadFile(s.tokensFile())

	if err != nil {
		fmt.Println("Не создан файл токенов /web/.storage/tokens")
		return false, err
	}

	for _, tokenHash := range bytes.Split(bb, []byte("\n")) {
		tokenHash = bytes.TrimSpace(tokenHash)
		if bytes.Compare(hash(token), tokenHash) == 0 {
			return true, nil
		}
	}
	return false, nil
}

func (s FileStorage) SetToken(token string) error {
	f, err := os.OpenFile(s.tokensFile(), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	f.Write([]byte("\n"))
	f.Write(hash(token))
	return nil
}

func (s FileStorage) Password(login string) (checksum string, err error) {
	panic("not implemented") // TODO: Implement
}

func (s FileStorage) SetCredentials(login string, pass string) error {
	panic("not implemented") // TODO: Implement
}

const recipeFolderName = "recipe"

func (s FileStorage) recipeFolder() string {
	return path.Join(s.folder, recipeFolderName)
}

func (s FileStorage) tokensFile() string {
	return path.Join(s.folder, tokenFile)
}

func hash(value string) []byte {
	h := sha256.Sum256([]byte(value))
	return h[:]
}
