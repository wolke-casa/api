package utils

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"io"
)

func IoReaderToByteSlice(body io.Reader) []byte {
	buf := new(bytes.Buffer)
	buf.ReadFrom(body)

	return buf.Bytes()
}

func GenerateSecureToken(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}

func CheckIfElementExists(slice []string, key string) bool {
	for _, v := range slice {
		if v == key {
			return true
		}
	}
	return false
}
