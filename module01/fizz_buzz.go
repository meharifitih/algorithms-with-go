package module01

import (
	"fmt"
)

// FizzBuzz will print out all of the numbers
// from 1 to N replacing any divisible by 3
// with "Fizz", and divisible by 5 with "Buzz",
// and any divisible by both with "Fizz Buzz".
//
// Note: The test for this is a little
// complicated so that you can just use the
// `fmt` package and print to standard out.
// I wouldn't normally recommend this, but did
// it here to make life easier for beginners.
func FizzBuzz(n int) {
	for i := 1; i <= n; i++ {

		switch {
		// check number id divisible by 5 and 3 with out any remainder
		case i%5 == 0 && i%3 == 0:
			fmt.Print("Fizz Buzz")
		// check num is divisible by 3
		case i%3 == 0:
			fmt.Print("Fizz")
		// check num is divisible by 3
		case i%5 == 0:
			fmt.Print("Buzz")
		default:
			fmt.Print(i)
		}
		// add  comma and space after each number
		if i != n {
			fmt.Print(", ")
		}
	}

	// add new line for string
	fmt.Println()
}

// using separate function to print
// func printFizzBuzz(i int) {
// 	switch {
// 	case i%5 == 0 && i%3 == 0:
// 		fmt.Print("Fizz Buzz")
// 	case i%3 == 0:
// 		fmt.Print("Fizz")
// 	case i%5 == 0:
// 		fmt.Print("Buzz")
// 	default:
// 		fmt.Print(i)
// 	}
// }
