package random

import (
	"math/rand"
	"time"
)

//NewRandomNumber returns a random number with max parameter
func NewRandomNumber(max int) int {
	seed := rand.NewSource(time.Now().UnixNano())
	randomNumber := rand.New(seed)

	return randomNumber.Intn(max)

}
