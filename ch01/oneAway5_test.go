package ch01

import "testing"
import "fmt"

func TestOnAway(t *testing.T) {

	text1 := "paless"
	text2 := "pale"
	result := isOneAway(text1, text2)
	fmt.Println(result)
}
