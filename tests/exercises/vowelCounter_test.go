package exercises

import (
	exercise "go/interview/reviews/pkg/exercises"
	"testing"
)

func TestVoweCounter(t *testing.T) {

	word := "paralelepipedo"
	count := exercise.CountVowels(word)
	expectedResult := 7

	if count != expectedResult {
		t.Fatalf(`result is not equals as expected`)
	}
}
