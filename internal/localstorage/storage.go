package localstorage

import (
	"bytes"
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
	Recipes() ([]data.Recipe, error)

	Tokens() ([]string, error)
	AddToken(hash []byte) error

	Password(login string) (checksum []byte, err error)
	SetPassword(login string, hash []byte) error
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

func (s FileStorage) Tokens() (hashes []string, err error) {
	//TODO получается хранилище не только сохраняет и получает токены но и как бы их проверяет, это по дурацки как-то
	bb, err := os.ReadFile(s.tokensFile())

	if err != nil {
		fmt.Println("Не создан файл токенов /web/.storage/tokens")
		return nil, err
	}

	for _, hash := range bytes.Split(bb, []byte("\n")) {
		hash = bytes.TrimSpace(hash)
		hashes = append(hashes, string(hash))
	}
	return hashes, nil
}

func (s FileStorage) AddToken(hash []byte) error {
	f, err := os.OpenFile(s.tokensFile(), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	f.Write([]byte("\n"))
	f.Write(hash)
	return nil
}

func (s FileStorage) Password(login string) (checksum []byte, err error) {
	bb, err := os.ReadFile(s.passhashFile(login))
	if err != nil {
		return nil, err
	}
	return bb, nil
}

func (s FileStorage) SetPassword(login string, hash []byte) error {
	err := os.WriteFile(s.passhashFile(login), hash, os.FileMode(os.O_WRONLY|os.O_CREATE))
	if err != nil {
		return err
	}
	return nil
}

const tokenFile = "tokens"
const recipeFolderName = "recipe"
const passHashExt = ".passhash"

func (s FileStorage) recipeFolder() string {
	return path.Join(s.folder, recipeFolderName)
}

func (s FileStorage) tokensFile() string {
	return path.Join(s.folder, tokenFile)
}

func (s FileStorage) passhashFile(login string) string {
	return path.Join(s.folder, login+passHashExt)
}

var _ Storager = new(FileStorage)
