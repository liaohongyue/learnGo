package utils

import (
	"math/rand"
	"time"
)

func RandString(n int) string {
	rand.Seed(time.Now().Unix())
	letters := []byte("asdfghjklqwertyuiopzxcvbnm")
	result := make([]byte, n)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
