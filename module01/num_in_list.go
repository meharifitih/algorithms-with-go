package module01

// NumInList will return true if num is in the
// list slice, and false otherwise.
func NumInList(list []int, num int) bool {

	// iterate through the list and check if the number exists
	for _, value := range list {
		if value == num {
			return true
		}
	}
	return false
}
