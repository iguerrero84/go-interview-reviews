package main

import (
	"fmt"
	exercise "go/interview/reviews/pkg/exercises"
)

func main() {
	runPalindrome()
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
