package ch01

import "testing"

func TestIsUniquePart1(t *testing.T) {
	input1 := "#777"

	if isUniquePart1(input1) {
		t.Fatalf("Input %s does not have unique elements", input1)
	}

	input2 := "#1234"
	if !isUniquePart1(input2) {
		t.Fatalf("Input %s have unique elements", input2)
	}
}

func TestIsUniquePart2(t *testing.T) {
	input1 := "#777"

	if iSUniquePart2(input1) {
		t.Fatalf("Input %s does not have unique elements", input1)
	}

	input2 := "#1234"
	if !iSUniquePart2(input2) {
		t.Fatalf("Input %s have unique elements", input2)
	}

	input3 := "#7 7 7"
	if iSUniquePart2(input3) {
		t.Fatalf("Input %s does not have unique elements", input3)
	}
}
