package exercises

import "strings"

func IsPalindrome(word string) bool {
	runes := []rune(strings.ToLower(word))
	for i := 0; i < len(runes)/2; i++ {
		if runes[i] != runes[len(runes)-1-i] {
			return false
		}
	}
	return true
}
