package main

// Exercise represents a single learning exercise
type Exercise struct {
	Title       string
	Description string
	Example     string
	Template    string
	Solution    string
	Validator   func(string) bool
}

// Trainer manages the learning session
type Trainer struct {
	exercises []Exercise
	current   int
	score     int
}