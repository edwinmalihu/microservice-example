package utils

import (
	"crypto/sha256"
	"fmt"
)

func HashPassword(pass *string) {
	bytePass := []byte(*pass)
	hPass := sha256.Sum256(bytePass)
	datapass := fmt.Sprintf("%x", hPass)
	*pass = string(datapass)
}

func ComparePassword(dbPass, pass string) bool {
	verify := false

	if dbPass == pass {
		verify = true
	} else {
		verify = false
	}

	return verify

}
