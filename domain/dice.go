package domain

import (
	"math/rand"
	"time"
)

type Dice struct{}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func (d Dice) Roll() int8 {
	value := int8(rand.Intn(6) + 1)
	return value
}
