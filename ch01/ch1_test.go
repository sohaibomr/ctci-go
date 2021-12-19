package ch01

import "testing"

func TestZeroMatrix(t *testing.T) {
	totalRows := 9
	totalColumns := 2
	matrix := [][]int{{1, 2}, {3, 0}, {5, 6}, {1, 2}, {3, 0}, {5, 6}, {1, 2}, {3, 0}, {5, 6}}
	zeroMatrix(totalRows, totalColumns, matrix)
}

func TestStringRotation(t *testing.T) {
	s1 := "waterbottle"
	s2 := "erbottlewat"
	expected := true
	actual := stringRotation(s1, s2)
	if actual != expected {
		t.Fatalf("%s is rotation of %s", s2, s1)
	}
}

func TestStringRotation2(t *testing.T) {
	s1 := "abc"
	s2 := "bca"
	expected := true
	actual := stringRotation(s1, s2)
	if actual != expected {
		t.Fatalf("%s is rotation of %s", s2, s1)
	}
}
