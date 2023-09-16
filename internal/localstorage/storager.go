package localstorage

import "github.com/thefrol/go-vue-recipe-blog/internal/data"

// Storager дает доступ до основых сущеностей так,
// чтобы мы могли не париться над тем, как они там лежат
// мы можем получить нужные рецепты и установить токены
// проверить логин и пароль
type Storager interface {
	Recipe(id string) (*data.Recipe, error) // странно что тут мы обрабатываем данные, а в токенах нет, думаю, это должны быть какие-то другие все же слои этим заниматься #todo
	SetRecipe(id string, r data.Recipe)
	Recipes() ([]data.Recipe, error)

	Token(hash []byte) (found bool, err error)
	AddToken(hash []byte) error

	Password(login string) (checksum []byte, err error)
	SetPassword(login string, hash []byte) error
}
