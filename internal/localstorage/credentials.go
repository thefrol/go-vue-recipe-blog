package localstorage

import (
	"os"
	"path"
)

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

func (s FileStorage) passhashFile(login string) string {
	return path.Join(s.folder, login+passHashExt)
}
