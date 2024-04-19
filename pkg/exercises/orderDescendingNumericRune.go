package exercises

import "sort"

func OrderDescendingNumericRunes(r []rune) []rune {

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

	return sortedRunes
}
