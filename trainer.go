package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// NewTrainer creates a new trainer instance
func NewTrainer() *Trainer {
	return &Trainer{
		exercises: GetExercises(),
		current:   0,
		score:     0,
	}
}

// Start begins the training session
func (t *Trainer) Start() {
	fmt.Println("🎯 Welcome to the Go Trainer!")
	fmt.Println("This interactive trainer will help you learn Go step by step.")
	fmt.Println("Type 'help' for commands, 'quit' to exit.\n")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("\n--- Exercise %d/%d ---\n", t.current+1, len(t.exercises))
		fmt.Printf("Score: %d/%d\n\n", t.score, len(t.exercises))

		if t.current >= len(t.exercises) {
			t.showFinalScore()
			break
		}

		exercise := t.exercises[t.current]
		fmt.Printf("📚 %s\n", exercise.Title)
		fmt.Printf("%s\n\n", exercise.Description)
		fmt.Printf("📖 Examples:\n%s\n\n", exercise.Example)
		fmt.Printf("Template:\n%s\n\n", exercise.Template)

		for {
			fmt.Print("Enter your solution (or 'hint', 'skip', 'quit'): ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			switch strings.ToLower(input) {
			case "quit":
				fmt.Println("Thanks for using Go Trainer! Keep practicing!")
				return
			case "help":
				t.showHelp()
			case "hint":
				t.showHint()
			case "skip":
				fmt.Printf("Skipped. The solution was:\n%s\n", exercise.Solution)
				t.current++
				goto nextExercise
			default:
				if t.checkSolution(input) {
					fmt.Println("✅ Correct! Well done!")
					t.score++
					t.current++
					goto nextExercise
				} else {
					fmt.Println("❌ Not quite right. Try again or type 'hint' for help.")
				}
			}
		}
		nextExercise:
	}
}

// checkSolution validates the user's solution
func (t *Trainer) checkSolution(solution string) bool {
	if t.current >= len(t.exercises) {
		return false
	}
	return t.exercises[t.current].Validator(solution)
}

// showHint displays the solution for the current exercise
func (t *Trainer) showHint() {
	if t.current >= len(t.exercises) {
		return
	}
	exercise := t.exercises[t.current]
	fmt.Printf("💡 Hint: %s\n", exercise.Solution)
}

// showHelp displays available commands
func (t *Trainer) showHelp() {
	fmt.Println("\nAvailable commands:")
	fmt.Println("  help  - Show this help message")
	fmt.Println("  hint  - Show the solution for current exercise")
	fmt.Println("  skip  - Skip current exercise")
	fmt.Println("  quit  - Exit the trainer")
}

// showFinalScore displays the final results
func (t *Trainer) showFinalScore() {
	fmt.Println("\n🎉 Congratulations! You've completed all exercises!")
	fmt.Printf("Final Score: %d/%d\n", t.score, len(t.exercises))

	percentage := float64(t.score) / float64(len(t.exercises)) * 100

	if percentage >= 80 {
		fmt.Println("🌟 Excellent work! You have a solid understanding of Go basics!")
	} else if percentage >= 60 {
		fmt.Println("👍 Good job! Review the concepts you missed and try again.")
	} else {
		fmt.Println("💪 Keep practicing! Consider reviewing Go fundamentals.")
	}

	fmt.Println("\nNext steps:")
	fmt.Println("- Practice more with goroutines and channels")
	fmt.Println("- Build a small project using these concepts")
	fmt.Println("- Explore the Go standard library")
}