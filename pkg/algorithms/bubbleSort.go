package searchingalgorithms

import "errors"

func BubbleSort(arr ...int) ([]int, error) {
	n := len(arr)
	if n == 0 {
		return arr, errors.New("empty array")
	}
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	return arr, nil
}
