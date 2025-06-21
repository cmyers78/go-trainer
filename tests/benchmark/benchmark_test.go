package benchmark

import (
	"testing"
	"time"

	"github.com/cmyers78/go-trainer/internal/exercises"
	"github.com/cmyers78/go-trainer/internal/models"
	"github.com/cmyers78/go-trainer/internal/trainer"
)

func BenchmarkExerciseRegistry(b *testing.B) {
	for i := 0; i < b.N; i++ {
		exercises.NewRegistry()
	}
}

func BenchmarkGetAllExercises(b *testing.B) {
	registry := exercises.NewRegistry()
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		registry.GetAll()
	}
}

func BenchmarkVariablesValidator(b *testing.B) {
	exercise := exercises.GetVariablesExercise()
	validator := exercise.Challenges[0].Validator
	solution := `var name string = "John"`
	
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		validator(solution)
	}
}

func BenchmarkFunctionsValidator(b *testing.B) {
	exercise := exercises.GetFunctionsExercise()
	validator := exercise.Challenges[0].Validator
	solution := `func add(a, b int) int { return a + b }`
	
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		validator(solution)
	}
}

func BenchmarkTrainerCreation(b *testing.B) {
	registry := exercises.NewRegistry()
	exerciseList := registry.GetAll()
	config := models.TrainerConfig{
		MaxAttempts:    3,
		TimeLimit:      time.Hour,
		ShowHints:      true,
		AdaptivePacing: true,
		CognitiveLoad:  models.Beginner,
	}
	
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		trainer.NewCLTTrainer(exerciseList, config)
	}
}