package module01

// Reverse will return the provided word in reverse
// order. Eg:
//
//	Reverse("cat") => "tac"
//	Reverse("alphabet") => "tebahpla"
func Reverse(word string) string {
	var res string
	// // iterate in normal order
	// for i := 0; i < len(word); i++ {
	// 	// attach the character to res
	// 	res = string(word[i]) + res

	// }

	for _, r := range word {
		res = string(r) + res
	}

	return res
}

// func Reverse(word string) string {
// 	var res string
// 	// var sb strings.Builder
// 	// iterate in reverse order
// 	for i := len(word) - 1; i >= 0; i-- {
// 		// attach the character to res
// 		res += string(word[i])
// 		// sb.WriteByte(word[i])
// 	}

// 	return res
// 	// return sb.String()
// }

// starting from the first letter and using while loops
// func Reverse(word string) string {
// 	// to hold all the characters in the word and split them into a single string
// 	charmap := strings.Split(word, "")

// 	// to hold the chars in reverse order
// 	var reverse []string

// 	// iterate through the charmap until it is zero or more characters
// 	for len(charmap) > 0 {

// 		// iterate through the charmaps and check for their orders
// 		for i, ch := range charmap {
// 			// if it is the last character append it to the reverse
// 			if i == len(charmap)-1 {
// 				reverse = append(reverse, ch)
// 				// remove the last character from the charmap
// 				charmap = charmap[:i]
// 			}
// 		}

// 	}

// 	return strings.Join(reverse, "")
// }
