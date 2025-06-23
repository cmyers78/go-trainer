package exercises

import (
	"strings"
	"github.com/cmyers78/go-trainer/internal/models"
)

// GetStructsExercise creates a comprehensive structs learning module
func GetStructsExercise() models.Exercise {
	return models.Exercise{
		ID:             "structs",
		Title:          "Structs and Methods",
		Description:    "Learn to create custom types with structs and methods",
		CognitiveLevel: models.Intermediate,
		ExerciseType:   models.Application,
		Prerequisites:  []string{"variables", "basic-types", "composite-types", "functions"},
		LearningGoals: []string{
			"Define custom types using structs",
			"Create and initialize struct instances",
			"Add methods to structs",
			"Understand value vs pointer receivers",
			"Use struct embedding for composition",
		},
		Examples: []models.Example{
			{
				Title: "Basic Struct Definition",
				Code: `type Person struct {
    Name string
    Age  int
    City string
}

func main() {
    // Different ways to create struct instances
    p1 := Person{Name: "Alice", Age: 30, City: "New York"}
    p2 := Person{"Bob", 25, "Boston"}  // Positional
    
    var p3 Person  // Zero value
    p3.Name = "Carol"
    p3.Age = 35
    
    fmt.Printf("%+v\n", p1)  // {Name:Alice Age:30 City:New York}
}`,
				Explanation: "Structs group related data. Use named fields for clarity. Zero value creates struct with field zero values.",
				Output: "Custom data types with grouped fields",
			},
			{
				Title: "Methods on Structs",
				Code: `type Rectangle struct {
    Width, Height float64
}

// Method with value receiver
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Method with pointer receiver (can modify)
func (r *Rectangle) Scale(factor float64) {
    r.Width *= factor
    r.Height *= factor
}

func main() {
    rect := Rectangle{Width: 10, Height: 5}
    fmt.Println("Area:", rect.Area())
    
    rect.Scale(2)  // Modifies original
    fmt.Println("New area:", rect.Area())
}`,
				Explanation: "Methods are functions with receivers. Value receivers get copies, pointer receivers can modify the original.",
				Output: "Behavior attached to custom types",
			},
			{
				Title: "Struct Embedding (Composition)",
				Code: `type Address struct {
    Street, City, State string
    ZipCode int
}

type Person struct {
    Name string
    Age  int
    Address  // Embedded struct
}

func main() {
    p := Person{
        Name: "Alice",
        Age:  30,
        Address: Address{
            Street:  "123 Main St",
            City:    "Boston",
            State:   "MA",
            ZipCode: 02101,
        },
    }
    
    // Access embedded fields directly
    fmt.Println(p.Street)  // Same as p.Address.Street
    fmt.Println(p.City)    // Same as p.Address.City
}`,
				Explanation: "Embedding promotes fields from embedded struct. Provides composition-based inheritance alternative.",
				Output: "Composition through struct embedding",
			},
			{
				Title: "Struct Tags and JSON",
				Code: `import "encoding/json"

type User struct {
    ID       int    \`json:"id"\`
    Username string \`json:"username"\`
    Email    string \`json:"email,omitempty"\`
    password string // lowercase = private, won't be exported
}

func main() {
    user := User{
        ID:       1,
        Username: "alice",
        Email:    "alice@example.com",
    }
    
    jsonData, _ := json.Marshal(user)
    fmt.Println(string(jsonData))
    // Output: {"id":1,"username":"alice","email":"alice@example.com"}
}`,
				Explanation: "Struct tags provide metadata. JSON tags control serialization. Uppercase fields are exported (public).",
				Output: "Metadata-driven serialization and encapsulation",
			},
		},
		Challenges: []models.Challenge{
			{
				Description: "Create a Book struct and a method to display book information",
				Template: `package main

import "fmt"

// Define Book struct with Title, Author, Pages, and Year fields

// Add a method Info() that returns a formatted string describing the book

func main() {
    book := Book{
        Title:  "The Go Programming Language",
        Author: "Donovan & Kernighan",
        Pages:  380,
        Year:   2015,
    }
    
    fmt.Println(book.Info())
}`,
				Solution: `type Book struct {
    Title  string
    Author string
    Pages  int
    Year   int
}

func (b Book) Info() string {
    return fmt.Sprintf("%s by %s (%d, %d pages)", b.Title, b.Author, b.Year, b.Pages)
}`,
				Hints: []string{
					"Define struct with field names and types",
					"Method has (b Book) receiver",
					"Use fmt.Sprintf for formatted string",
				},
				Validator: func(code string) bool {
					return strings.Contains(code, "type Book struct") &&
						   strings.Contains(code, "func (") &&
						   strings.Contains(code, "Info()") &&
						   strings.Contains(code, "Title") &&
						   strings.Contains(code, "Author")
				},
			},
			{
				Description: "Create a BankAccount struct with methods to deposit and withdraw money",
				Template: `package main

import "fmt"

// Define BankAccount struct with Owner (string) and Balance (float64)

// Add Deposit method that adds amount to balance (use pointer receiver)

// Add Withdraw method that subtracts amount from balance (use pointer receiver)
// Return error if insufficient funds

// Add GetBalance method that returns current balance (value receiver)

func main() {
    account := BankAccount{Owner: "Alice", Balance: 100.0}
    
    account.Deposit(50.0)
    fmt.Printf("After deposit: $%.2f\n", account.GetBalance())
    
    err := account.Withdraw(30.0)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Printf("After withdrawal: $%.2f\n", account.GetBalance())
    }
}`,
				Solution: `type BankAccount struct {
    Owner   string
    Balance float64
}

func (b *BankAccount) Deposit(amount float64) {
    b.Balance += amount
}

func (b *BankAccount) Withdraw(amount float64) error {
    if amount > b.Balance {
        return fmt.Errorf("insufficient funds")
    }
    b.Balance -= amount
    return nil
}

func (b BankAccount) GetBalance() float64 {
    return b.Balance
}`,
				Hints: []string{
					"Use pointer receivers for methods that modify",
					"Return error for invalid operations",
					"Value receiver for read-only methods",
				},
				Validator: func(code string) bool {
					return strings.Contains(code, "type BankAccount struct") &&
						   strings.Contains(code, "*BankAccount) Deposit") &&
						   strings.Contains(code, "*BankAccount) Withdraw") &&
						   strings.Contains(code, "error") &&
						   strings.Contains(code, "GetBalance")
				},
			},
			{
				Description: "Create an Employee struct that embeds a Person struct",
				Template: `package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

// Define Employee struct that embeds Person and adds JobTitle and Salary fields

// Add a method Introduction() that returns a string like "Hi, I'm Alice, a Software Engineer"

func main() {
    emp := Employee{
        Person: Person{Name: "Alice", Age: 30},
        JobTitle: "Software Engineer",
        Salary:   75000,
    }
    
    fmt.Println(emp.Introduction())
    fmt.Printf("Name: %s, Age: %d, Salary: $%d\n", emp.Name, emp.Age, emp.Salary)
}`,
				Solution: `type Employee struct {
    Person   // Embedded struct
    JobTitle string
    Salary   int
}

func (e Employee) Introduction() string {
    return fmt.Sprintf("Hi, I'm %s, a %s", e.Name, e.JobTitle)
}`,
				Hints: []string{
					"Embed Person struct without field name",
					"Access embedded fields directly (e.Name)",
					"Method receiver works on Employee type",
				},
				Validator: func(code string) bool {
					return strings.Contains(code, "type Employee struct") &&
						   strings.Contains(code, "Person") &&
						   strings.Contains(code, "JobTitle") &&
						   strings.Contains(code, "func (e Employee) Introduction()") &&
						   strings.Contains(code, "e.Name")
				},
			},
		},
		EstimatedTime: 25,
	}
}