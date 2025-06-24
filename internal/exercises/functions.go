package exercises

import (
	"strings"
	"github.com/cmyers78/claude/internal/models"
)

// GetFunctionsExercise creates a comprehensive functions learning module
// Applies progressive disclosure and worked examples
func GetFunctionsExercise() models.Exercise {
	return models.Exercise{
		ID:             "functions",
		Title:          "Functions",
		Description:    "Learn to create and use functions effectively",
		CognitiveLevel: models.Beginner,
		ExerciseType:   models.Application,
		Prerequisites:  []string{"variables", "basic-types", "composite-types"},
		LearningGoals: []string{
			"Write functions with parameters and return values",
			"Understand function signatures and naming conventions",
			"Apply functions to solve problems",
		},
		Examples: []models.Example{
			{
				Title: "Basic Function (No Parameters, No Return)",
				Code: `func sayHello() {
	fmt.Println("Hello, World!")
}

func main() {
	sayHello() // Call the function
}`,
				Explanation: "The simplest function form. Uses 'func' keyword, followed by name and parentheses.",
				Output: "Hello, World!",
			},
			{
				Title: "Function with Parameters",
				Code: `func greet(name string) {
	fmt.Println("Hello,", name)
}

func main() {
	greet("Alice")
	greet("Bob")
}`,
				Explanation: "Parameters go inside parentheses with their types. Call function by passing arguments.",
				Output: "Hello, Alice\nHello, Bob",
			},
			{
				Title: "Function with Return Value",
				Code: `func add(a, b int) int {
	return a + b
}

func main() {
	result := add(5, 3)
	fmt.Println("Sum:", result)
}`,
				Explanation: "Return type comes after parameters. Use 'return' to send value back to caller.",
				Output: "Sum: 8",
			},
			{
				Title: "Multiple Return Values (Go Specialty)",
				Code: `func divmod(a, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

func main() {
	q, r := divmod(17, 5)
	fmt.Printf("17 ÷ 5 = %d remainder %d\n", q, r)
}`,
				Explanation: "Go functions can return multiple values. Very useful for error handling patterns.",
				Output: "17 ÷ 5 = 3 remainder 2",
			},
		},
		Challenges: []models.Challenge{
			{
				Description: "Create a function 'add' that takes two integers and returns their sum",
				Template: `package main

import "fmt"

func main() {
	result := add(5, 3)
	fmt.Println("Result:", result)
}

// Your function here`,
				Solution: `func add(a, b int) int {
	return a + b
}`,
				Hints: []string{
					"Function name should be 'add'",
					"Takes two int parameters",
					"Returns one int value",
					"Use 'return' to send back the sum",
				},
				Validator: func(code string) bool {
					return strings.Contains(code, "func add") && 
						   strings.Contains(code, "int") && 
						   strings.Contains(code, "return")
				},
			},
			{
				Description: "Create a function 'multiply' that multiplies two numbers",
				Template: `package main

import "fmt"

func main() {
	result := multiply(4, 7)
	fmt.Println("Result:", result)
}

// Your function here`,
				Solution: `func multiply(a, b int) int {
	return a * b
}`,
				Hints: []string{
					"Similar to add, but use * operator",
					"Same pattern: func name(params) returnType { return value }",
				},
				Validator: func(code string) bool {
					return strings.Contains(code, "func multiply") && 
						   strings.Contains(code, "int") && 
						   strings.Contains(code, "return") &&
						   strings.Contains(code, "*")
				},
			},
			{
				Description: "Create a function that returns both quotient and remainder (division)",
				Template: `package main

import "fmt"

func main() {
	q, r := divide(17, 5)
	fmt.Printf("17 ÷ 5 = %d remainder %d\n", q, r)
}

// Your function here - return TWO values`,
				Solution: `func divide(a, b int) (int, int) {
	return a / b, a % b
}`,
				Hints: []string{
					"Return type should be (int, int) for two values",
					"Use a/b for quotient, a%b for remainder",
					"Return both values separated by comma",
				},
				Validator: func(code string) bool {
					return strings.Contains(code, "func divide") && 
						   strings.Contains(code, "(int, int)") && 
						   strings.Contains(code, "return") &&
						   (strings.Contains(code, "/") && strings.Contains(code, "%"))
				},
			},
		},
		EstimatedTime: 15,
	}
}