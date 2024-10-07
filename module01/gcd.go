package module01

func GCD(a, b int) int {
	// // Step 1: if b == 0, then return a
	// if b == 0 {
	// 	return a
	// }

	// // Step 2: A becomes b and B becomes the remainder of A divided by B
	// // a, b = b, a % b
	// a, b = b, a%b

	// // Step 3: Go to step 1
	// return GCD(a, b)

	// using Euclidean algorithm
	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}

	return a

}
