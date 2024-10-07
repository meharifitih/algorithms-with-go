package module01

// Factor takes in a list of primes and a number and factors that number with
// the provided primes.
//
// The returned numbers can be in any order; tests will sort them in increasing
// order to make checking easier.
//
// Bonus: Any remainder (aside from 1) that can't be factored will be treated as
// a prime in the results.
//
// Examples:
//
//	Factor([]int{2,3,5}, 30) // []int{2,3,5}
//	Factor([]int{2,3,5}, 28) // []int{2,2,7}
//	Factor([]int{2,3,5}, 720) // []int{2,2,2,2,3,3,5}
//
// Examples with remainders:
//
//	Factor([2,5], 30) // []int{2,5,3}
//	Factor([3,5], 720) // []int{3,3,5,16}
//	Factor([], 4) // []int{4}
func Factor(primes []int, number int) []int {
	var result []int

	// iterate over the primes and check if it is divisible by the prime
	for _, prime := range primes {
		// also divide by the same prime if it is possible
		for number%prime == 0 {
			// append the prime to result
			result = append(result, prime)
			// the number value changes to number divide by prime
			number = number / prime
		}

	}
	// for _, prime1 := range primes {
	// 	if number%prime1 == 0 {
	// 		result = append(result, prime1)
	// 		number = number / prime1
	// 		for _, prime2 := range primes {
	// 			if number%prime2 == 0 {
	// 				result = append(result, prime2)
	// 				number = number / prime2
	// 			}
	// 		}
	// 	}
	// }

	if number > 1 {
		result = append(result, number)
	}

	return result
}
