package main

import (
	"testing"
)

func TestGetExercises(t *testing.T) {
	exercises := GetExercises()
	
	expectedCount := 5
	if len(exercises) != expectedCount {
		t.Errorf("Expected %d exercises, got %d", expectedCount, len(exercises))
	}
	
	for i, exercise := range exercises {
		if exercise.Title == "" {
			t.Errorf("Exercise %d has empty title", i)
		}
		if exercise.Description == "" {
			t.Errorf("Exercise %d has empty description", i)
		}
		if exercise.Example == "" {
			t.Errorf("Exercise %d has empty example", i)
		}
		if exercise.Template == "" {
			t.Errorf("Exercise %d has empty template", i)
		}
		if exercise.Solution == "" {
			t.Errorf("Exercise %d has empty solution", i)
		}
		if exercise.Validator == nil {
			t.Errorf("Exercise %d has nil validator", i)
		}
	}
}

func TestVariablesValidator(t *testing.T) {
	exercises := GetExercises()
	validator := exercises[0].Validator
	
	testCases := []struct {
		input    string
		expected bool
	}{
		{`var name string = "John"`, true},
		{`var age int = 25`, false}, // wrong type
		{`name := "John"`, false},   // no var keyword
		{`var name string`, false},  // no assignment
		{`var name = "John"`, false}, // no explicit type
	}
	
	for _, tc := range testCases {
		result := validator(tc.input)
		if result != tc.expected {
			t.Errorf("Input %q: expected %v, got %v", tc.input, tc.expected, result)
		}
	}
}

func TestFunctionsValidator(t *testing.T) {
	exercises := GetExercises()
	validator := exercises[1].Validator
	
	testCases := []struct {
		input    string
		expected bool
	}{
		{`func add(a, b int) int { return a + b }`, true},
		{`func add(a int, b int) int { return a + b }`, true},
		{`func multiply(a, b int) int { return a * b }`, false}, // wrong name
		{`func add(a, b string) string { return a + b }`, false}, // wrong type
		{`func add(a, b int) { fmt.Println(a + b) }`, false}, // no return type
	}
	
	for _, tc := range testCases {
		result := validator(tc.input)
		if result != tc.expected {
			t.Errorf("Input %q: expected %v, got %v", tc.input, tc.expected, result)
		}
	}
}

func TestSlicesValidator(t *testing.T) {
	exercises := GetExercises()
	validator := exercises[2].Validator
	
	testCases := []struct {
		input    string
		expected bool
	}{
		{`numbers := []int{1, 2, 3, 4, 5}`, true},
		{`var numbers = []int{1, 2, 3, 4, 5}`, true},
		{`numbers := []string{"a", "b", "c"}`, false}, // wrong type
		{`numbers := [5]int{1, 2, 3, 4, 5}`, false}, // array, not slice
		{`numbers := make([]int, 5)`, false}, // no literal braces
	}
	
	for _, tc := range testCases {
		result := validator(tc.input)
		if result != tc.expected {
			t.Errorf("Input %q: expected %v, got %v", tc.input, tc.expected, result)
		}
	}
}

func TestStructsValidator(t *testing.T) {
	exercises := GetExercises()
	validator := exercises[3].Validator
	
	testCases := []struct {
		input    string
		expected bool
	}{
		{`type Person struct { Name string; Age int }`, true},
		{`type Person struct { Name string \n Age int }`, true},
		{`type User struct { Name string; Age int }`, false}, // wrong name
		{`type Person struct { Name string }`, false}, // missing Age
		{`type Person struct { Age int }`, false}, // missing Name
	}
	
	for _, tc := range testCases {
		result := validator(tc.input)
		if result != tc.expected {
			t.Errorf("Input %q: expected %v, got %v", tc.input, tc.expected, result)
		}
	}
}

func TestMethodsValidator(t *testing.T) {
	exercises := GetExercises()
	validator := exercises[4].Validator
	
	testCases := []struct {
		input    string
		expected bool
	}{
		{`func (p Person) Greet() string { return "Hello" }`, true},
		{`func (p *Person) Greet() string { return "Hello" }`, true},
		{`func (u User) Greet() string { return "Hello" }`, false}, // wrong receiver
		{`func (p Person) Hello() string { return "Hello" }`, false}, // wrong name
		{`func (p Person) Greet() { fmt.Println("Hello") }`, false}, // no return type
	}
	
	for _, tc := range testCases {
		result := validator(tc.input)
		if result != tc.expected {
			t.Errorf("Input %q: expected %v, got %v", tc.input, tc.expected, result)
		}
	}
}