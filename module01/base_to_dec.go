package module01

import (
	"math"
	"strings"
)

// BaseToDec takes in a number and the base it is currently
// in and returns the decimal equivalent as an integer.
//
// Eg:
//
//	BaseToDec("E", 16) => 14
//	BaseToDec("1110", 2) => 14
func BaseToDec(value string, base int) int {
	charset := "0123456789ABCDEF"

	var result int = 0

	power := 0
	// iterate through value and find the index character in charset
	for i := len(value) - 1; i >= 0; i-- {
		// find index from charset for the current value of index
		index := strings.Index(charset, string(value[i]))

		// power of base with power and sum with previous value
		result = result + int(math.Pow(float64(base), float64(power)))*index

		// increment power by one
		power++
	}

	// multiplier := 1
	// for i := len(value) - 1; i >= 0; i-- {
	// 	index := -1
	// 	for j, char := range charset {
	// 		if rune(value[i]) == char {
	// 			index = j
	// 			break
	// 		}

	// 	}
	// 	result += index * multiplier
	// 	multiplier *= base
	// }

	return result
}
