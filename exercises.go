package main

import "strings"

// GetExercises returns all available exercises
func GetExercises() []Exercise {
	return []Exercise{
		{
			Title:       "Variables and Types",
			Description: "Declare a variable 'name' of type string and assign it your name",
			Example:     "Variables in Go can be declared in several ways:\n\n// Method 1: var keyword with explicit type\nvar message string = \"Hello, World!\"\n\n// Method 2: var keyword with type inference\nvar count = 42\n\n// Method 3: short declaration (inside functions only)\nage := 25\n\n// Basic types: string, int, float64, bool\nvar isActive bool = true",
			Template:    "package main\n\nfunc main() {\n\t// Your code here\n\tfmt.Println(name)\n}",
			Solution:    `var name string = "YourName"`,
			Validator: func(code string) bool {
				return strings.Contains(code, "var") && strings.Contains(code, "string") && strings.Contains(code, "=")
			},
		},
		{
			Title:       "Functions",
			Description: "Create a function 'add' that takes two integers and returns their sum",
			Example:     "Functions in Go are defined with the 'func' keyword:\n\n// Basic function\nfunc sayHello() {\n    fmt.Println(\"Hello!\")\n}\n\n// Function with parameters\nfunc greet(name string) {\n    fmt.Println(\"Hello,\", name)\n}\n\n// Function with return value\nfunc multiply(a, b int) int {\n    return a * b\n}\n\n// Multiple return values\nfunc divmod(a, b int) (int, int) {\n    return a / b, a % b\n}",
			Template:    "package main\n\nfunc main() {\n\tresult := add(5, 3)\n\tfmt.Println(result)\n}\n\n// Your function here",
			Solution:    "func add(a, b int) int {\n\treturn a + b\n}",
			Validator: func(code string) bool {
				return strings.Contains(code, "func add") && strings.Contains(code, "int") && strings.Contains(code, "return")
			},
		},
		{
			Title:       "Slices",
			Description: "Create a slice of integers with values 1, 2, 3, 4, 5 and print its length",
			Example:     "Slices are dynamic arrays in Go:\n\n// Creating slices\nvar numbers []int                    // empty slice\nscores := []int{95, 87, 92}         // slice literal\ndata := make([]int, 5)               // slice with length 5\n\n// Slice operations\nlen(scores)                          // length: 3\nscores = append(scores, 98)          // add element\nsubset := scores[1:3]                // slice from index 1 to 2",
			Template:    "package main\n\nfunc main() {\n\t// Your code here\n\tfmt.Println(len(numbers))\n}",
			Solution:    "numbers := []int{1, 2, 3, 4, 5}",
			Validator: func(code string) bool {
				return strings.Contains(code, "[]int") && strings.Contains(code, "{") && strings.Contains(code, "}")
			},
		},
		{
			Title:       "Structs",
			Description: "Define a struct 'Person' with fields Name (string) and Age (int)",
			Example:     "Structs group related data together:\n\n// Define a struct\ntype Car struct {\n    Brand string\n    Model string\n    Year  int\n}\n\n// Create and use structs\nvar myCar Car\nmyCar.Brand = \"Toyota\"\n\n// Struct literal\ncar2 := Car{Brand: \"Honda\", Model: \"Civic\", Year: 2023}\n\n// Short form (order matters)\ncar3 := Car{\"Ford\", \"Mustang\", 2022}",
			Template:    "package main\n\nfunc main() {\n\tp := Person{Name: \"Alice\", Age: 30}\n\tfmt.Println(p)\n}\n\n// Your struct here",
			Solution:    "type Person struct {\n\tName string\n\tAge  int\n}",
			Validator: func(code string) bool {
				return strings.Contains(code, "type Person struct") && strings.Contains(code, "Name string") && strings.Contains(code, "Age")
			},
		},
		{
			Title:       "Methods",
			Description: "Add a method 'Greet' to the Person struct that returns a greeting string",
			Example:     "Methods are functions with a receiver:\n\n// Method with value receiver\nfunc (c Car) GetInfo() string {\n    return c.Brand + \" \" + c.Model\n}\n\n// Method with pointer receiver (can modify)\nfunc (c *Car) UpdateYear(year int) {\n    c.Year = year\n}\n\n// Usage\ncar := Car{\"BMW\", \"X5\", 2020}\ninfo := car.GetInfo()        // call method\ncar.UpdateYear(2023)         // modify through pointer",
			Template:    "package main\n\ntype Person struct {\n\tName string\n\tAge  int\n}\n\nfunc main() {\n\tp := Person{Name: \"Alice\", Age: 30}\n\tfmt.Println(p.Greet())\n}\n\n// Your method here",
			Solution:    "func (p Person) Greet() string {\n\treturn \"Hello, I'm \" + p.Name\n}",
			Validator: func(code string) bool {
				return strings.Contains(code, "func (p Person) Greet()") && strings.Contains(code, "string") && strings.Contains(code, "return")
			},
		},
	}
}
