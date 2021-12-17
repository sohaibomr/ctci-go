package ch01

import "testing"

func TestURLIfy(t *testing.T) {
	input := "Mr John Smith "
	expected := "Mr%20John%20Smith"

	actual := uRLIfyString(input)
	if actual != expected {
		t.Fatalf("Failed to URLIfy the string %s", input)
	}
}
