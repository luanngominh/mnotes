package util

import (
	"crypto/rand"
	"math/big"
	"strconv"
)

//GenerateToken generate 6 digits
func GenerateToken() string {
	token := ""
	for i := 0; i < 6; i++ {
		n, _ := rand.Int(rand.Reader, big.NewInt(10))
		token = token + strconv.Itoa(int(n.Int64()))
	}
	return token
}
