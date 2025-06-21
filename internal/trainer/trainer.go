package trainer

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/cmyers78/go-trainer/internal/models"
)

// CLTTrainer implements Cognitive Load Theory principles
type CLTTrainer struct {
	config     models.TrainerConfig
	exercises  []models.Exercise
	progress   []models.LearningProgress
	current    int
	startTime  time.Time
}

// NewCLTTrainer creates a new trainer with CLT principles
func NewCLTTrainer(exercises []models.Exercise, config models.TrainerConfig) *CLTTrainer {
	return &CLTTrainer{
		config:    config,
		exercises: exercises,
		progress:  make([]models.LearningProgress, len(exercises)),
		current:   0,
		startTime: time.Now(),
	}
}

// Start begins the training session with CLT-informed pacing
func (t *CLTTrainer) Start() {
	t.showWelcome()
	
	reader := bufio.NewReader(os.Stdin)
	
	for t.current < len(t.exercises) {
		exercise := t.exercises[t.current]
		t.startExercise(exercise)
		
		// Show learning goals first (reduce extraneous load)
		t.showLearningGoals(exercise)
		
		// Progressive disclosure: examples before challenges
		t.showExamples(exercise)
		
		// Wait for learner to process examples
		fmt.Print("\nPress Enter when ready to try the challenges...")
		reader.ReadString('\n')
		
		// Present challenges with faded guidance
		completed := t.runChallenges(exercise, reader)
		
		if completed {
			t.completeExercise(exercise)
			t.current++
		} else {
			break // User quit
		}
	}
	
	t.showFinalResults()
}

// showWelcome introduces the training with clear expectations
func (t *CLTTrainer) showWelcome() {
	fmt.Println("🧠 Go Trainer with Cognitive Load Theory")
	fmt.Println("=========================================")
	fmt.Println()
	fmt.Println("This trainer uses proven learning science principles:")
	fmt.Println("• Worked examples before practice")
	fmt.Println("• Progressive disclosure of complexity")
	fmt.Println("• Multiple practice opportunities")
	fmt.Println("• Adaptive pacing based on your progress")
	fmt.Println()
	fmt.Println("Commands: 'hint', 'skip', 'quit', 'help'")
	fmt.Println()
}

// showLearningGoals clearly states what the learner will achieve
func (t *CLTTrainer) showLearningGoals(exercise models.Exercise) {
	fmt.Printf("📋 %s\n", exercise.Title)
	fmt.Printf("Description: %s\n\n", exercise.Description)
	
	fmt.Println("🎯 Learning Goals:")
	for i, goal := range exercise.LearningGoals {
		fmt.Printf("   %d. %s\n", i+1, goal)
	}
	fmt.Println()
	
	if len(exercise.Prerequisites) > 0 {
		fmt.Println("📚 Prerequisites:")
		for _, prereq := range exercise.Prerequisites {
			fmt.Printf("   • %s\n", prereq)
		}
		fmt.Println()
	}
	
	fmt.Printf("⏱️  Estimated time: %d minutes\n\n", exercise.EstimatedTime)
}

// showExamples implements worked example effect
func (t *CLTTrainer) showExamples(exercise models.Exercise) {
	fmt.Println("📖 Examples (Study these carefully):")
	fmt.Println("=====================================")
	
	for i, example := range exercise.Examples {
		fmt.Printf("\n%d. %s\n", i+1, example.Title)
		fmt.Println(strings.Repeat("-", len(example.Title)+3))
		
		fmt.Printf("Code:\n```go\n%s\n```\n\n", example.Code)
		fmt.Printf("Explanation: %s\n", example.Explanation)
		
		if example.Output != "" {
			fmt.Printf("Output: %s\n", example.Output)
		}
		
		fmt.Println()
	}
}

// runChallenges implements faded guidance and completion effect
func (t *CLTTrainer) runChallenges(exercise models.Exercise, reader *bufio.Reader) bool {
	fmt.Println("🎯 Practice Challenges:")
	fmt.Println("======================")
	
	for i, challenge := range exercise.Challenges {
		fmt.Printf("\nChallenge %d/%d\n", i+1, len(exercise.Challenges))
		fmt.Println(strings.Repeat("-", 15))
		
		if !t.runSingleChallenge(challenge, i, reader) {
			return false // User quit
		}
	}
	
	return true
}

// runSingleChallenge handles individual challenge with adaptive support
func (t *CLTTrainer) runSingleChallenge(challenge models.Challenge, challengeNum int, reader *bufio.Reader) bool {
	fmt.Printf("Task: %s\n\n", challenge.Description)
	fmt.Printf("Template:\n```go\n%s\n```\n\n", challenge.Template)
	
	attempts := 0
	hintsUsed := 0
	
	for attempts < t.config.MaxAttempts {
		fmt.Print("Your solution: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		
		switch strings.ToLower(input) {
		case "quit":
			return false
		case "help":
			t.showHelp()
			continue
		case "hint":
			if hintsUsed < len(challenge.Hints) {
				fmt.Printf("💡 Hint: %s\n", challenge.Hints[hintsUsed])
				hintsUsed++
			} else {
				fmt.Printf("💡 Solution: %s\n", challenge.Solution)
			}
			continue
		case "skip":
			fmt.Printf("⏭️  Skipped. Solution: %s\n", challenge.Solution)
			return true
		default:
			attempts++
			if challenge.Validator(input) {
				fmt.Println("✅ Excellent! That's correct!")
				
				// Provide elaborative feedback for learning
				if attempts == 1 && hintsUsed == 0 {
					fmt.Println("🌟 Perfect on first try!")
				} else if attempts <= 2 {
					fmt.Println("👍 Good work!")
				} else {
					fmt.Println("💪 Great persistence!")
				}
				return true
			} else {
				t.provideAdaptiveFeedback(attempts, challengeNum, challenge)
			}
		}
	}
	
	fmt.Printf("Max attempts reached. Solution: %s\n", challenge.Solution)
	return true
}

// provideAdaptiveFeedback gives targeted help based on CLT principles
func (t *CLTTrainer) provideAdaptiveFeedback(attempts int, challengeNum int, challenge models.Challenge) {
	if attempts == 1 {
		// First mistake: gentle guidance
		fmt.Println("❌ Not quite right. Compare your answer with the examples above.")
	} else if attempts == 2 {
		// Second mistake: more specific help
		fmt.Println("❌ Still not correct. Type 'hint' for guidance, or review the examples.")
	} else {
		// Multiple mistakes: direct support
		fmt.Println("❌ Let's break this down. Type 'hint' for step-by-step help.")
	}
}

// startExercise initializes tracking for an exercise
func (t *CLTTrainer) startExercise(exercise models.Exercise) {
	t.progress[t.current] = models.LearningProgress{
		ExerciseID: exercise.ID,
		StartTime:  time.Now(),
		Attempts:   0,
	}
}

// completeExercise finalizes tracking for an exercise
func (t *CLTTrainer) completeExercise(exercise models.Exercise) {
	now := time.Now()
	t.progress[t.current].CompletedAt = &now
	t.progress[t.current].TimeSpent = now.Sub(t.progress[t.current].StartTime)
	
	fmt.Printf("✅ %s completed!\n", exercise.Title)
	fmt.Printf("Time spent: %.1f minutes\n\n", t.progress[t.current].TimeSpent.Minutes())
}

// showHelp provides contextual assistance
func (t *CLTTrainer) showHelp() {
	fmt.Println("\n📚 Available Commands:")
	fmt.Println("  hint  - Get a helpful hint for the current challenge")
	fmt.Println("  skip  - Skip the current challenge and see the solution")
	fmt.Println("  quit  - Exit the trainer")
	fmt.Println("  help  - Show this help message")
	fmt.Println()
}

// showFinalResults provides comprehensive learning summary
func (t *CLTTrainer) showFinalResults() {
	fmt.Println("\n🎉 Training Complete!")
	fmt.Println("====================")
	
	totalTime := time.Since(t.startTime)
	completed := 0
	
	for _, progress := range t.progress {
		if progress.CompletedAt != nil {
			completed++
		}
	}
	
	fmt.Printf("Exercises completed: %d/%d\n", completed, len(t.exercises))
	fmt.Printf("Total time: %.1f minutes\n", totalTime.Minutes())
	
	// Learning reinforcement
	fmt.Println("\n🧠 Key Concepts Learned:")
	for i, exercise := range t.exercises[:completed] {
		fmt.Printf("  %d. %s\n", i+1, exercise.Title)
		for _, goal := range exercise.LearningGoals {
			fmt.Printf("     • %s\n", goal)
		}
	}
	
	fmt.Println("\n🚀 Next Steps:")
	fmt.Println("  • Practice these concepts in your own projects")
	fmt.Println("  • Explore Go's standard library")
	fmt.Println("  • Join the Go community online")
}