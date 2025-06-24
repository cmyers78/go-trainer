package exercises

import (
	"strings"
	"github.com/cmyers78/claude/internal/models"
)

// GetBasicTypesExercise creates a comprehensive basic types learning module
func GetBasicTypesExercise() models.Exercise {
	return models.Exercise{
		ID:             "basic-types",
		Title:          "Basic Data Types",
		Description:    "Master Go's fundamental data types and constants",
		CognitiveLevel: models.Beginner,
		ExerciseType:   models.Concept,
		Prerequisites:  []string{"variables"},
		LearningGoals: []string{
			"Understand Go's numeric types and their ranges",
			"Work with strings and string operations",
			"Use constants effectively",
			"Choose appropriate types for different use cases",
		},
		Examples: []models.Example{
			{
				Title: "Numeric Types",
				Code: `var age int = 25           // Platform-dependent size
var count int32 = 1000       // Exactly 32 bits
var distance int64 = 384400  // Exactly 64 bits
var temperature float32 = 98.6
var precision float64 = 3.14159265359`,
				Explanation: "Go has specific numeric types. int is platform-dependent, while int32/int64 are fixed sizes. float64 is preferred for most floating-point calculations.",
				Output: "Different numeric types with specific bit sizes",
			},
			{
				Title: "String Operations",
				Code: `name := "Go"
greeting := "Hello, " + name + "!"
length := len(greeting)
first := greeting[0]        // byte value
substring := greeting[0:5]  // "Hello"`,
				Explanation: "Strings are immutable byte sequences. Use + for concatenation, len() for length, and slicing for substrings.",
				Output: "String manipulation and access operations",
			},
			{
				Title: "Constants",
				Code: `const Pi = 3.14159
const MaxUsers = 100
const (
    StatusOK = 200
    StatusNotFound = 404
    StatusError = 500
)`,
				Explanation: "Constants are compile-time values that cannot change. Group related constants in blocks.",
				Output: "Constants for fixed values and enumerations",
			},
			{
				Title: "Type Conversions",
				Code: `var i int = 42
var f float64 = float64(i)  // Explicit conversion required
var s string = fmt.Sprintf("%d", i)  // Convert to string
var b byte = byte(i)        // Convert to byte`,
				Explanation: "Go requires explicit type conversions. No automatic conversion between different numeric types.",
				Output: "Safe type conversions between compatible types",
			},
		},
		Challenges: []models.Challenge{
			{
				Description: "Declare constants for a simple HTTP status system",
				Template: `package main

import "fmt"

// Declare constants here using a const block
// StatusOK = 200
// StatusNotFound = 404  
// StatusError = 500

func main() {
	fmt.Println("OK:", StatusOK)
	fmt.Println("Not Found:", StatusNotFound)
	fmt.Println("Error:", StatusError)
}`,
				Solution: `const (
	StatusOK = 200
	StatusNotFound = 404
	StatusError = 500
)`,
				Hints: []string{
					"Use const block with parentheses",
					"Each constant on its own line",
					"No var keyword needed",
				},
				Validator: func(code string) bool {
					return strings.Contains(code, "const (") &&
						   strings.Contains(code, "StatusOK") &&
						   strings.Contains(code, "200") &&
						   strings.Contains(code, "StatusNotFound") &&
						   strings.Contains(code, "404")
				},
			},
			{
				Description: "Create variables with specific numeric types and convert between them",
				Template: `package main

import "fmt"

func main() {
	// Declare age as int32 with value 25
	// Declare height as float64 with value 5.9
	// Convert age to float64 and store in ageFloat
	
	fmt.Printf("Age: %d, Height: %.1f, Age as float: %.1f\n", age, height, ageFloat)
}`,
				Solution: `var age int32 = 25
var height float64 = 5.9
ageFloat := float64(age)`,
				Hints: []string{
					"Use int32 for age",
					"Use float64 for height",
					"Use float64(age) for conversion",
				},
				Validator: func(code string) bool {
					return strings.Contains(code, "int32") &&
						   strings.Contains(code, "float64") &&
						   strings.Contains(code, "float64(age)")
				},
			},
			{
				Description: "Work with strings - create a full name from first and last name",
				Template: `package main

import "fmt"

func main() {
	firstName := "John"
	lastName := "Doe"
	
	// Create fullName by concatenating firstName + " " + lastName
	// Get the length of fullName
	// Get first character of fullName
	
	fmt.Printf("Full name: %s\n", fullName)
	fmt.Printf("Length: %d\n", nameLength)
	fmt.Printf("First character: %c\n", firstChar)
}`,
				Solution: `fullName := firstName + " " + lastName
nameLength := len(fullName)
firstChar := fullName[0]`,
				Hints: []string{
					"Use + to concatenate strings",
					"Use len() function for length",
					"Use [0] to get first character",
				},
				Validator: func(code string) bool {
					return strings.Contains(code, "firstName + ") &&
						   strings.Contains(code, "len(") &&
						   strings.Contains(code, "[0]")
				},
			},
		},
		EstimatedTime: 12,
	}
}