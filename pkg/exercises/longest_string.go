package exercises

func LengthOfLongestSubstring(s string) (int, error) {
	charIndexMap := make(map[rune]int)
	maxLength := 0
	start := 0

	for i, char := range s {
		if index, ok := charIndexMap[char]; ok && index > start {
			start = index + 1
		}
		charIndexMap[char] = i
		currentLength := i - start + 1

		if currentLength > maxLength {
			maxLength = currentLength
		}
	}

	return maxLength, nil
}
