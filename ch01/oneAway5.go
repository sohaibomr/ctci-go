package ch01

//One Away: There are three types of edits that can be performed on strings: insert a character, remove a character, or replace a character. Given two strings, write a function to check if they are one edit (or zero edits) away.
//EXAMPLE
//pale, ple -> true pales, pale -> true pale, bale -> true pale, bake -> false
import "fmt"

func isOneAway(text1 string, text2 string) bool {
	editCount := 0
	text1S := []rune(text1)
	text2S := []rune(text2)

	lenText2 := len(text2S)
	// if idx is equal to len of text2 and edit count is 0, increase edit count and break
	for idx, elem := range text1S {
		if editCount > 1 {
			return false
		}
		if idx >= lenText2 {
			editCount += 1
			continue
		}
		if elem != text2S[idx] {
			editCount += 1
			text2S = append(text2S, 0)
			copy(text2S[idx+1:], text2S[idx:])
			text2S[idx] = elem
			fmt.Println(string(text2S))
		}

	}
	if editCount > 1 {
		return false
	}
	return true
}
