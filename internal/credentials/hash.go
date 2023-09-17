package credentials

import "crypto/sha256"

func Hash(value string) []byte {
	h := sha256.Sum256([]byte(value))
	return h[:]
}
