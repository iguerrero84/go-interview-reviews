package exercises

import (
	"errors"
	"sort"
)

func OrderDescendingNumericRunes(r []rune) ([]rune, error) {

	if len(r) == 0 {
		return r, errors.New("array rune length is zero")
	}

	// Convert runes to integers
	ints := make([]int, len(r))
	for i, v := range r {
		ints[i] = int(v - '0')
	}

	// Sort integeers in descending order
	sort.Sort(sort.Reverse(sort.IntSlice(ints)))

	// Convert integers back to runes
	sortedRunes := make([]rune, len(r))
	for i, v := range ints {
		sortedRunes[i] = rune(v + '0')
	}

	return sortedRunes, nil
}
