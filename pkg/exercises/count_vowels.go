package exercises

import "strings"

func CountVowels(word string) int {
	vowels := "aeiou"
	count := 0
	for _, char := range word {
		if strings.ContainsRune(vowels, char) {
			count++
		}
	}
	return count
}
