package localstorage

import (
	"bytes"
	"fmt"
	"os"
	"path"
)

func (s FileStorage) Tokens() (hashes []string, err error) {
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

func (s FileStorage) tokensFile() string {
	return path.Join(s.folder, tokenFile)
}
