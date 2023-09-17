package credentials

import (
	"encoding/json"

	"github.com/thefrol/go-vue-recipe-blog/internal/filestorage"
)

type Keyer interface {
	Get(string) ([]byte, error)
	Set(string, []byte) error
}

// Tokens хранит токены в указанной папке в виде джейсон файлов,
// проверяет валидность токенов
// Токены хранятся хешами на диске, чтобы токены невозможно было угнать
// Если токен найден, то возвращает структуру в нем обраруженную в формате json хранящуюся
type Tokens struct {
	storage Keyer
}

func LocalTokens(folder string) *Tokens {
	s := filestorage.New(folder).WithPostfix(".token")
	return &Tokens{storage: s}
}

// проверяет валидность токена на текущую дату и пути
// acl можно задать по системе путей /admin/moderator/recipe.editor, -
// такой человек получит права редактора рецептов
func (t Tokens) Get(token string, tokendata interface{}) (bool, error) {
	data, err := t.storage.Get(string(Hash(token)))
	if err != nil {
		return false, err
	}
	err = json.Unmarshal(data, &tokendata)
	if err != nil {
		return false, err
	}
	return true, nil

}
