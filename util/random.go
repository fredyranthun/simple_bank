package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var rng *rand.Rand

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	source := rand.NewSource(time.Now().UnixNano())
	rng = rand.New(source)
}

// Generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rng.Int63n(max - min + 1)
}

// RandomString Generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rng.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	currencies := []string{EUR, USD, CAD}
	n := len(currencies)
	return currencies[rng.Intn(n)]
}

func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}