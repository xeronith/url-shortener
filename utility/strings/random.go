package strings

import (
	"math/rand"
	"strings"
	"time"
)

const AlphanumericCharacters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GenerateRandomString(length int) string {
	alphanumericCharactersCount := len(AlphanumericCharacters)

	builder := strings.Builder{}
	for i := 0; i < length; i++ {
		randomAlphanumericCharacter := AlphanumericCharacters[rand.Intn(alphanumericCharactersCount)]
		builder.WriteByte(randomAlphanumericCharacter)
	}

	return builder.String()
}
