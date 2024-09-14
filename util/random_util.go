package util

import (
	"math/rand"
	"time"
)

func RandomInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func RandomTime(min int, max int) time.Duration {
	return time.Duration(min+rand.Intn(max-min)) * time.Second
}
