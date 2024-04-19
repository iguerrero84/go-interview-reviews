package main

import (
	"fmt"
	exercise "go/interview/reviews/pkg/exercises"
)

func main() {
	runOrderAscendingNumericRune()
}

func runOrderAscendingNumericRune() {
	r := []rune{'3', '1', '5', '9', '2', '6'}
	orderedRune := exercise.OrderAscendingNumericRunes(r)

	for i := 0; i < len(orderedRune); i++ {
		fmt.Println(orderedRune[i] - '0')
	}
}

func runOrderDescendingNumericRune() {
	r := []rune{'3', '1', '5', '9', '2', '6'}
	orderedRune := exercise.OrderDescendingNumericRunes(r)

	for i := 0; i < len(orderedRune); i++ {
		fmt.Println(orderedRune[i] - '0')
	}
}

func runPalindrome() {
	word := "level"
	fmt.Println(exercise.IsPalindrome(word))
}

func runVowelCounter() {
	word := "paralelepipedo"
	fmt.Println(exercise.CountVowels(word))
}

func runFizzBuzz() {
	exercise.FizzBuzz()
}
