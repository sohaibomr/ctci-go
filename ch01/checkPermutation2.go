package ch01

import (
	"sort"
)

//Check Permutation: Given two strings,write a method to decide if one is a permutation of the
//other.
func checkPermutation(text1 string, text2 string) bool {
	//fail if len is different
	text1S := []rune(text1)
	text2S := []rune(text2)
	sort.Slice(text1S, func(i, j int) bool {
		return text1S[i] < text1S[j]

	})
	sort.Slice(text2S, func(i, j int) bool {
		return text2S[i] < text2S[j]

	})
	if string(text1S) == string(text2S) {
		return true
	}
	return false
}
