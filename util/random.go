package util

import (
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const alphabet = "abcdefghigklmnopqrstuvwxyz"

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(b int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < b; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomOwner() string {
	return RandomString(1000)
}

func RandomMoney() int64 {
	return RandomInt(1000, 1200)
}

func RandomCurrency() string {
	currencies := []string{"USD", "CAD", "KZT"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
