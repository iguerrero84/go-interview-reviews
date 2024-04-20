package exercises

import (
	"errors"
	"sort"
)

func OrderAscendingNumericRunes(r []rune) ([]rune, error) {

	if len(r) == 0 {
		return r, errors.New("rune array empty")
	}

	// Convert runes to integers
	ints := make([]int, len(r))
	for i, v := range r {
		ints[i] = int(v - '0')
	}

	// Sort integeers in descending order
	sort.Ints(ints)

	// Convert integers back to runes
	sortedRunes := make([]rune, len(r))
	for i, v := range ints {
		sortedRunes[i] = rune(v + '0')
	}

	return sortedRunes, nil
}
