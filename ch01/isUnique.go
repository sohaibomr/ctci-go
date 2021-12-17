package ch01

// Is Unique: Implement an algorithm to determine if a string has all unique characters. What if you cannot use additional data structures?
// Hints: #44, #7 7 7, #732

func isUniquePart1(text string) bool {
	var elemsExists = make(map[string]bool)
	for _, elem := range text {
		elemS := string(elem)
		if exists, _ := elemsExists[elemS]; exists {
			return false
		}
		elemsExists[elemS] = true
	}
	return true
}

func iSUniquePart2(text string) bool {
	//without additional data structure

	//compare each element with every other element
	for idx, elem := range text {
		for idx2, elem2 := range text {
			if elem2 == elem && idx != idx2 {
				return false
			}
		}
	}
	return true
}