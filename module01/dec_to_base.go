package module01

import (
	"fmt"
	"math"
)

// DecToBase will return a string representing
// the provided decimal number in the provided base.
// This is limited to bases 2-16 for simplicity.
//
// Eg:
//
//	DecToBase(14, 16) => "E"
//	DecToBase(14, 2) => "1110"
func DecToBase(dec, base int) string {
	var result string
	for dec > 0 {

		rem := dec % base

		// using switch statement to set the value to the remaining
		// if dec%base >= 10 && base >= 10 {
		// 	switch dec % base {
		// 	case 10:
		// 		newDigit = "A"
		// 	case 11:
		// 		newDigit = "B"
		// 	case 12:
		// 		newDigit = "C"
		// 	case 13:
		// 		newDigit = "D"
		// 	case 14:
		// 		newDigit = "E"
		// 	case 15:
		// 		newDigit = "F"
		// 	}
		// } else {
		// 	newDigit = fmt.Sprintf("%d", dec%base)
		// }

		// using fmt.Sprintf to set values for hex and binary
		result = fmt.Sprintf("%X%s", rem, result)

		// result = newDigit + result
		dec = int(math.Floor(float64(dec) / float64(base)))

	}

	return result
}
