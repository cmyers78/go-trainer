package exercises

import (
	"strings"
	"github.com/cmyers78/claude/internal/models"
)

// GetCompositeTypesExercise creates a comprehensive composite types learning module
func GetCompositeTypesExercise() models.Exercise {
	return models.Exercise{
		ID:             "composite-types",
		Title:          "Composite Types",
		Description:    "Master arrays, slices, and maps in Go",
		CognitiveLevel: models.Intermediate,
		ExerciseType:   models.Application,
		Prerequisites:  []string{"variables", "basic-types"},
		LearningGoals: []string{
			"Understand the difference between arrays and slices",
			"Create and manipulate slices effectively",
			"Use maps for key-value storage",
			"Choose appropriate composite types for different scenarios",
		},
		Examples: []models.Example{
			{
				Title: "Arrays (Fixed Size)",
				Code: `var numbers [5]int = [5]int{1, 2, 3, 4, 5}
var names [3]string = [3]string{"Alice", "Bob", "Charlie"}

// Array literal with inferred size
colors := [...]string{"red", "green", "blue"}

fmt.Println("Length:", len(numbers))  // 5
fmt.Println("First:", numbers[0])     // 1`,
				Explanation: "Arrays have fixed size determined at compile time. Size is part of the type. Use [...] to let compiler count elements.",
				Output: "Fixed-size collections with compile-time size",
			},
			{
				Title: "Slices (Dynamic Arrays)",
				Code: `// Create slice from array
numbers := []int{1, 2, 3, 4, 5}

// Add elements
numbers = append(numbers, 6, 7)

// Create slice with make
scores := make([]int, 5)      // length 5, capacity 5
buffer := make([]int, 0, 10)  // length 0, capacity 10

// Slice operations
subset := numbers[1:4]  // [2, 3, 4]`,
				Explanation: "Slices are dynamic arrays. Use append() to add elements. make() creates slices with specific length/capacity.",
				Output: "Dynamic arrays that can grow and shrink",
			},
			{
				Title: "Maps (Key-Value Storage)",
				Code: `// Create and initialize map
ages := map[string]int{
    "Alice": 30,
    "Bob":   25,
    "Carol": 35,
}

// Add/update entries
ages["David"] = 28

// Check if key exists
age, exists := ages["Alice"]
if exists {
    fmt.Println("Alice is", age, "years old")
}

// Delete entry
delete(ages, "Bob")`,
				Explanation: "Maps store key-value pairs. Use comma ok idiom to check key existence. delete() removes entries.",
				Output: "Flexible key-value storage with existence checking",
			},
			{
				Title: "Iterating Collections",
				Code: `numbers := []int{10, 20, 30}
for i, value := range numbers {
    fmt.Printf("Index %d: %d\n", i, value)
}

ages := map[string]int{"Alice": 30, "Bob": 25}
for name, age := range ages {
    fmt.Printf("%s is %d years old\n", name, age)
}

// Use _ to ignore index/key
for _, value := range numbers {
    fmt.Println("Value:", value)
}`,
				Explanation: "Use range to iterate over slices and maps. Get both index/key and value. Use _ to ignore unwanted values.",
				Output: "Efficient iteration over collections",
			},
		},
		Challenges: []models.Challenge{
			{
				Description: "Create a slice of your favorite programming languages and add more languages to it",
				Template: `package main

import "fmt"

func main() {
	// Create a slice with 3 programming languages
	// Add 2 more languages using append
	// Print the final slice and its length
	
	fmt.Println("Languages:", languages)
	fmt.Println("Count:", len(languages))
}`,
				Solution: `languages := []string{"Go", "Python", "JavaScript"}
languages = append(languages, "Rust", "TypeScript")`,
				Hints: []string{
					"Use []string{} to create slice",
					"Use append() to add elements",
					"Remember to assign back to slice",
				},
				Validator: func(code string) bool {
					return strings.Contains(code, "[]string") &&
						   strings.Contains(code, "append(") &&
						   strings.Contains(code, "languages")
				},
			},
			{
				Description: "Create a map of country capitals and look up specific countries",
				Template: `package main

import "fmt"

func main() {
	// Create a map with at least 3 country-capital pairs
	// Look up "France" and check if it exists
	// Print the result
	
	capital, exists := capitals["France"]
	if exists {
		fmt.Printf("Capital of France: %s\n", capital)
	} else {
		fmt.Println("France not found")
	}
}`,
				Solution: `capitals := map[string]string{
	"France": "Paris",
	"Japan": "Tokyo",
	"Brazil": "Brasília",
}`,
				Hints: []string{
					"Use map[string]string{} syntax",
					"Include France with Paris as capital",
					"Use key: value pairs",
				},
				Validator: func(code string) bool {
					return strings.Contains(code, "map[string]string") &&
						   strings.Contains(code, "France") &&
						   strings.Contains(code, "capitals")
				},
			},
			{
				Description: "Process a slice of numbers - find sum and average",
				Template: `package main

import "fmt"

func main() {
	numbers := []int{10, 20, 30, 40, 50}
	
	// Calculate sum using range loop
	// Calculate average (sum / length)
	
	fmt.Printf("Numbers: %v\n", numbers)
	fmt.Printf("Sum: %d\n", sum)
	fmt.Printf("Average: %.1f\n", average)
}`,
				Solution: `sum := 0
for _, num := range numbers {
	sum += num
}
average := float64(sum) / float64(len(numbers))`,
				Hints: []string{
					"Use range to iterate over slice",
					"Accumulate sum in a variable",
					"Convert to float64 for division",
				},
				Validator: func(code string) bool {
					return strings.Contains(code, "range") &&
						   strings.Contains(code, "sum") &&
						   strings.Contains(code, "float64")
				},
			},
			{
				Description: "Create a slice of slices (2D slice) representing a matrix",
				Template: `package main

import "fmt"

func main() {
	// Create a 3x3 matrix using slice of slices
	// Fill it with some numbers
	// Print each row
	
	for i, row := range matrix {
		fmt.Printf("Row %d: %v\n", i, row)
	}
}`,
				Solution: `matrix := [][]int{
	{1, 2, 3},
	{4, 5, 6},
	{7, 8, 9},
}`,
				Hints: []string{
					"Use [][]int for slice of int slices",
					"Each inner slice is a row",
					"Initialize with nested braces",
				},
				Validator: func(code string) bool {
					return strings.Contains(code, "[][]int") &&
						   strings.Contains(code, "matrix") &&
						   strings.Contains(code, "{") &&
						   strings.Count(code, "{") >= 4  // At least outer + 3 inner braces
				},
			},
		},
		EstimatedTime: 20,
	}
}