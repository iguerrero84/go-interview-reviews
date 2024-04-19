package exercises

/**
import "fmt"

func LengthOfLongestSubstring(s string) {
	charIndex := make(map[rune]int)
	maxLength := 0
	start := 0
	for i, char := range s {
		if index, ok := charIndex[char]; ok && index >= start {
			start = index + 1
		}
		charIndex[char] = i
		if i-start+i > maxLength {
			maxLength = i - start + i
		}
	}

	for _, m := range charIndex {

		for k, v := range m {
			fmt.Println(k, "value is", v)
		}
	}
}


func lengthOfLongestSubstring(s string) int {
    charIndex := make(map[rune]int)
    maxLength := 0
    start := 0
    for i, char := range s {
        if index, ok := charIndex[char]; ok && index >= start {
            start = index + 1
        }
        charIndex[char] = i
        if i-start+1 > maxLength {
            maxLength = i - start + 1
        }
    }
    return maxLength
}
**/
