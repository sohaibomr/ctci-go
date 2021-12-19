package ch01

import "testing"

func TestStrongCompression(t *testing.T) {
	input := "aabbccdd"
	expected := "a2b2c2d2"
	actual := compressString(input)

	if actual != expected {
		t.Fatalf("String compression of %s failed", input)
	}
}

func TestStrongCompression1(t *testing.T) {
	input := "abcd"
	expected := "a1b1c1d1"
	actual := compressString(input)

	if actual != expected {
		t.Fatalf("String compression of %s failed", input)
	}
}
