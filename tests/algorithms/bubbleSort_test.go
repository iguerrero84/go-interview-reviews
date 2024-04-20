package algorithms

import (
	"errors"
	"testing"

	algorithm "go/interview/reviews/pkg/algorithms"
)

func TestBubbleSort(t *testing.T) {
	inputArr := []int{56, 78, 98, 43, 39, 50, 23, 14, 67}

	expectedArr := []int{14, 23, 39, 43, 50, 56, 67, 78, 98}

	resultArr, err := algorithm.BubbleSort(inputArr...)

	if len(resultArr) == 0 || err != nil {
		t.Fatalf(`BubbleSort("") = %q, %v, expectedArr "", error`, resultArr, err)
	}

	for i := 0; i < len(inputArr)-1; i++ {
		if resultArr[i] != expectedArr[i] {
			t.Fatalf(`BubbleSort("") = %q, %v, expectedArr "", error`, resultArr, errors.New("Ints are different"))
		}
	}

}

func TestBubbleSortWithEmptyInputArray(t *testing.T) {
	inputArr := []int{}

	expectedArr := []int{}

	resultArr, err := algorithm.BubbleSort(inputArr...)

	if len(resultArr) != 0 || err == nil {
		t.Fatalf(`BubbleSort("") = %q, %v, expectedArr "", error`, resultArr, err)
	}

	if len(resultArr) != len(expectedArr) {
		t.Fatalf(`BubbleSort("") = %q, %v, expectedArr "", error`, resultArr, errors.New("arrays length different"))
	}
}
