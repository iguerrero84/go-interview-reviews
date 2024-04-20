package exercises

import (
	"errors"
	"strings"
)

func IsPalindrome(word string) (bool, error) {
	if word == "" {
		return false, errors.New("empty input string")
	}

	runes := []rune(strings.ToLower(word))
	for i := 0; i < len(runes)/2; i++ {
		if runes[i] != runes[len(runes)-1-i] {
			return false, nil
		}
	}
	return true, nil
}
