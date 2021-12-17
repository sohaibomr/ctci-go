package ch01

import "testing"

func TestCheckPermutation(t *testing.T) {
	text1, text2 := "AABBCCDD", "DDBBAACC"
	result := checkPermutation(text1, text2)
	if !result {
		t.Fatalf("%s and %s Strings are permutation", text1, text2)
	}

}

func TestCheckPermutation2(t *testing.T) {
	text1, text2 := "AABBCCD", "DDBBAACC"
	result := checkPermutation(text1, text2)
	if result {
		t.Fatalf("%s and %s Strings are not permutation", text1, text2)
	}
	if !result {
		t.Logf("%s and %s Strings are not permutation", text1, text2)
	}

}
