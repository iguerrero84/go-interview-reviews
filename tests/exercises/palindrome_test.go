package exercises

import (
	exercise "go/interview/reviews/pkg/exercises"
	"testing"
)

func TestPalindrome(t *testing.T) {
	word := "level"
	result, err := exercise.IsPalindrome(word)

	if !result || err != nil {
		t.Fatalf(`input word is not palindrome`)
	}
}

func TestPalindromeWithNotPalindromeWord(t *testing.T) {
	word := "parking"
	result, err := exercise.IsPalindrome(word)

	if result || err != nil {
		t.Fatalf(`input word is palindrome`)
	}
}
