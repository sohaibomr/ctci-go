package ch01

import (
	"strings"
)

//String Rotation:Assumeyou have a method isSubstringwhich checks if one word is a substring of another. Given two strings, sl and s2, write code to check if s2 is a rotation of sl using only one call to isSubstring (e.g.,"waterbottle" is a rotation of"erbottlewat").

func stringRotation(s1 string, s2 string) bool {
	s1s1 := s1 + s1
	isSubstring := strings.Contains(s1s1, s2)
	return isSubstring
}
