// Хранит данные в файлах на диске, где один ключ это файл
package filestorage

import (
	"errors"
	"os"
)

// FileStorage позволяет хранить строки в файловой системе по ключам,
// Postfix, Prefix позволяеют настроить постфиксы и префиксы там в одноай
// папке можно будет хранить разные сущности, например пароли и токены
// для удобства создания префиксы и постфиксы можно назначить
// специальными функциями билдерами, чтобы получить вот так
// tokenStore := storage.New("my_folder").WithPrefix("token-")
type FileStorage struct {
	folder  string
	Postfix string
	Prefix  string
}

// New связывается с указанным хранилищем в папке storageFolder.
// Если такой папки нет, то создает
func New(storageFolder string) FileStorage {
	if _, err := os.Stat(storageFolder); errors.Is(err, os.ErrNotExist) {
		// такого пути нет
		err := os.Mkdir(storageFolder, os.FileMode(os.O_RDONLY|os.O_RDWR))
		if err != nil {
			panic("не удалется создать хранилище в папке " + storageFolder)
		}
	} else if err != nil {
		panic("Не удалось открыть хранилище в " + storageFolder)
	}
	return FileStorage{folder: storageFolder}
}

func (s FileStorage) Get(key string) ([]byte, error) {
	bb, err := os.ReadFile(s.filename(key))
	if err != nil {
		return nil, err
	}
	return bb, nil
}

func (s FileStorage) Set(key string, value []byte) error {
	err := os.WriteFile(s.filename(key), value, os.FileMode(os.O_WRONLY|os.O_CREATE))
	if err != nil {
		return err
	}
	return nil
}

func (s FileStorage) WithPrefix(prefix string) {
	s.Prefix = prefix
}

func (s FileStorage) WithPostfix(postfix string) {
	s.Postfix = postfix
}

// filename возвращает путь до файла с учетом срех настроен постфиксом и префиксов
func (s FileStorage) filename(key string) string {
	return s.Prefix + key + s.Postfix
}
