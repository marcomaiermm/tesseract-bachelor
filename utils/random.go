package utils

import (
	"math"
	"math/rand"
	"strconv"
)

// function to generate a random string number between given min and max digits
func RandomStringNumber(min int, max int) string {
	// generates a random number between min and max digits
	// if min and max are the same, it will generate a number with that many digits
	var randomDigits = rand.Intn(max-min+1) + min
	// generates a random string number
	var randomStringNumber = strconv.Itoa(rand.Intn(int(math.Pow10(randomDigits) - 1)) + int(math.Pow10(randomDigits - 1)))
	return randomStringNumber
}
