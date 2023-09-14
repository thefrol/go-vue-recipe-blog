package localstorage

import (
	"errors"
	"os"
)

const tokenFile = "tokens"
const recipeFolderName = "recipe"
const passHashExt = ".passhash"

// FileStorage Хранилище данных, реализовано в файлах
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

var _ Storager = new(FileStorage)
