package utils

import (
	"crypto/rand"
	"fmt"
)

func GenerateID() string {
	id := make([]byte, 16)
	if _, err := rand.Read(id); err != nil {
		panic(err)
	}
	return fmt.Sprintf("%x", id)
}
