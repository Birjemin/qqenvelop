package utils

import (
	"math/rand"
	"time"
)

// RandNum rand num
func RandNum(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}
