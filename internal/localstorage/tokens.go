package localstorage

import (
	"bytes"
	"fmt"
	"os"
	"path"
)

// TODO
// По хорошему хранить бы не только токены, но ещё и acl токена - права, и где он действует и когда перестанет работать

func (s FileStorage) Token(hash []byte) (found bool, err error) {
	bb, err := os.ReadFile(s.tokensFile())

	if err != nil {
		fmt.Println("Не создан файл токенов /web/.storage/tokens")
		return false, err
	}

	for _, h := range bytes.Split(bb, []byte("\n")) {
		h = bytes.TrimSpace(h)
		if bytes.Equal(hash, h) {
			return true, nil
		}
	}
	return false, nil
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
