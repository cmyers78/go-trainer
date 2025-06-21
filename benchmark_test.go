package main

import (
	"testing"
)

func BenchmarkNewTrainer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewTrainer()
	}
}

func BenchmarkGetExercises(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetExercises()
	}
}

func BenchmarkValidators(b *testing.B) {
	exercises := GetExercises()
	
	testInputs := []string{
		`var name string = "John"`,
		`func add(a, b int) int { return a + b }`,
		`numbers := []int{1, 2, 3, 4, 5}`,
		`type Person struct { Name string; Age int }`,
		`func (p Person) Greet() string { return "Hello" }`,
	}
	
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		for j, input := range testInputs {
			if j < len(exercises) {
				exercises[j].Validator(input)
			}
		}
	}
}

func BenchmarkCheckSolution(b *testing.B) {
	trainer := NewTrainer()
	solution := `var name string = "John"`
	
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		trainer.checkSolution(solution)
	}
}