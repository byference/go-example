package utils

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GetRandomIntn(i int) int {
	return rand.Intn(i)
}
