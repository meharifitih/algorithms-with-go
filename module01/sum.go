package module01

// Sum will sum up all of the numbers passed
// in and return the result
func Sum(numbers []int) int {
	// value to hold the sum of the numbers
	var sum int = 0

	// iterate over the numbers and add the value to the sum of previous
	for _, n := range numbers {
		sum += n
	}

	return sum
}

// with recursion
// func Sum(numbers []int) int {
// 	if len(numbers) == 0 {
// 		return 0
// 	}

// 	return numbers[0] + Sum(numbers[1:])
// }
