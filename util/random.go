package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1) // random in range (min; max)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		char := alphabet[rand.Intn(k)]
		sb.WriteByte(char)
	}
	return sb.String()
}

// func to random Owner name for account creation
func RandomOwner() string {
	return RandomString(6) // random name has length from 7 to 1
}

// func to random Money for account creation
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// func to random Currency for account creation
func RandomCurrency() string {
	enum := []string{"USD", "VND", "CAD"}
	k := len(enum)
	return enum[rand.Intn(k)]
}
