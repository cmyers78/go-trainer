package exercises

import (
	"strings"
	"github.com/cmyers78/go-trainer/internal/models"
)

// GetVariablesExercise creates a comprehensive variables learning module
// Applies CLT principles: worked examples, progressive disclosure, variability
func GetVariablesExercise() models.Exercise {
	return models.Exercise{
		ID:             "variables",
		Title:          "Variables and Types",
		Description:    "Master Go's variable declarations and type system",
		CognitiveLevel: models.Beginner,
		ExerciseType:   models.Concept,
		Prerequisites:  []string{},
		LearningGoals: []string{
			"Understand different variable declaration methods",
			"Choose appropriate declaration style for different contexts",
			"Work with Go's basic types confidently",
		},
		Examples: []models.Example{
			{
				Title: "Explicit Type Declaration",
				Code: `var message string = "Hello, World!"
var count int = 42
var isReady bool = true`,
				Explanation: "Use 'var' with explicit type when you need to be clear about the type or declare without immediate assignment.",
				Output: "Variables declared with explicit types",
			},
			{
				Title: "Type Inference",
				Code: `var message = "Hello, World!"  // string inferred
var count = 42              // int inferred  
var temperature = 98.6      // float64 inferred`,
				Explanation: "Go can infer types from the assigned values. This reduces verbosity while maintaining type safety.",
				Output: "Types automatically inferred from values",
			},
			{
				Title: "Short Declaration (Most Common)",
				Code: `message := "Hello, World!"
count := 42
temperature := 98.6
isReady := true`,
				Explanation: "Short declaration (:=) is the most concise form. Can only be used inside functions.",
				Output: "Concise variable declaration inside functions",
			},
			{
				Title: "Zero Values",
				Code: `var name string    // "" (empty string)
var age int        // 0
var height float64 // 0.0
var isActive bool  // false`,
				Explanation: "Variables declared without initialization get their type's zero value. This prevents undefined behavior.",
				Output: "Variables initialized to zero values",
			},
		},
		Challenges: []models.Challenge{
			{
				Description: "Declare a variable 'name' of type string and assign it your name using explicit type declaration",
				Template: `package main

import "fmt"

func main() {
	// Your code here - use var with explicit type
	fmt.Println("Hello,", name)
}`,
				Solution: `var name string = "YourName"`,
				Hints: []string{
					"Use the 'var' keyword",
					"Specify 'string' as the type",
					"Don't forget the assignment with =",
				},
				Validator: func(code string) bool {
					return strings.Contains(code, "var") && 
						   strings.Contains(code, "string") && 
						   strings.Contains(code, "=") &&
						   strings.Contains(code, "name")
				},
			},
			{
				Description: "Declare the same variable using type inference (no explicit type)",
				Template: `package main

import "fmt"

func main() {
	// Your code here - let Go infer the type
	fmt.Println("Hello,", name)
}`,
				Solution: `var name = "YourName"`,
				Hints: []string{
					"Use 'var' but omit the type",
					"Go will infer string from the value",
				},
				Validator: func(code string) bool {
					return strings.Contains(code, "var") && 
						   strings.Contains(code, "name") &&
						   strings.Contains(code, "=") &&
						   !strings.Contains(code, "string")
				},
			},
			{
				Description: "Now use short declaration syntax (most common in Go)",
				Template: `package main

import "fmt"

func main() {
	// Your code here - use := syntax
	fmt.Println("Hello,", name)
}`,
				Solution: `name := "YourName"`,
				Hints: []string{
					"Use := instead of var",
					"This is the most concise form",
				},
				Validator: func(code string) bool {
					return strings.Contains(code, ":=") && 
						   strings.Contains(code, "name")
				},
			},
		},
		EstimatedTime: 10,
	}
}