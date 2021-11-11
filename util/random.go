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

func RandomInt(min, max int64) int {
	return int(min + rand.Int63n(max-min+1))
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomName() string {
	return RandomString(8)
}

func RandomSumGoal() int {
	return RandomInt(0, 999999)
}

func RandomStartSum() int {
	return RandomInt(0, 99999)
}

func RandomChatId() int {
	return RandomInt(100, 999999)
}

func RandomSum() int {
	return RandomInt(0, 99999)
}
