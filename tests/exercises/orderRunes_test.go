package exercises

import (
	"testing"

	exercise "go/interview/reviews/pkg/exercises"
)

func TestOrderAscendingNumericRunes(t *testing.T) {
	inputRune := []rune{'3', '1', '5', '9', '2', '6'}

	length := len(inputRune)

	expectedResult := []int{1, 2, 3, 5, 6, 9}

	orderedRune, err := exercise.OrderAscendingNumericRunes(inputRune)

	if length == 0 || err != nil {
		t.Fatalf(`OrderAscendingRune("") = %q, %v, expectedResult "", error`, expectedResult, err)
	}

	for i := 0; i < length-1; i++ {
		if int(orderedRune[i]-'0') != expectedResult[i] {
			t.Fatalf(`OrderAscendingRune("") = %q, %v, expectedResult "", error`, expectedResult, err)
		}
	}
}

func TestOrderDescendingNumericRunes(t *testing.T) {
	inputRune := []rune{'3', '1', '5', '9', '2', '6'}

	length := len(inputRune)

	expectedResult := []int{9, 6, 5, 3, 2, 1}

	orderedRune, err := exercise.OrderDescendingNumericRunes(inputRune)

	if length == 0 || err != nil {
		t.Fatalf(`OrderAscendingRune("") = %q, %v, expectedResult "", error`, expectedResult, err)
	}

	for i := 0; i < length-1; i++ {
		if int(orderedRune[i]-'0') != expectedResult[i] {
			t.Fatalf(`OrderAscendingRune("") = %q, %v, expectedResult "", error`, expectedResult, err)
		}
	}
}
