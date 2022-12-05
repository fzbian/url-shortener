package utils

import (
	"math/rand"
	"time"
)

var (
	ABC = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func GenString(lenght int) string {
	rand.Seed(time.Now().UnixNano())
	var letters = []rune(ABC)
	b := make([]rune, lenght)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
