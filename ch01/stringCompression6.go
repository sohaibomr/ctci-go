package ch01

import (
	"fmt"
	"strings"
)

//TODO: Optimize me
//String Compression: Implement a method to perform basic string compression using the counts of repeated characters. For example, the string aabcccccaaa would become a2blc5a3. If the "compressed" string would not become smaller than the original string, your method should return
//the original string. You can assume the string has only uppercase and lowercase letters (a - z).
func compressString(text string) string {
	var wordCount = make(map[string]int)
	textS := []rune(text)
	var textUniq []string
	var sb strings.Builder
	for _, elem := range textS {
		elemS := string(elem)
		if _, ok := wordCount[elemS]; !ok {
			wordCount[elemS] = 1
			textUniq = append(textUniq, elemS)
		} else {
			wordCount[elemS] += 1
		}
	}

	for _, elem := range textUniq {
		elemS := string(elem)
		sb.WriteString(fmt.Sprintf("%s%d", elemS, wordCount[elemS]))
	}
	return sb.String()
}
