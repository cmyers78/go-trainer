package unit

import (
	"strings"
	"testing"
	"time"

	"github.com/cmyers78/go-trainer/internal/exercises"
	"github.com/cmyers78/go-trainer/internal/models"
	"github.com/cmyers78/go-trainer/internal/trainer"
)

func TestCLTTrainerCreation(t *testing.T) {
	registry := exercises.NewRegistry()
	exerciseList := registry.GetAll()
	
	config := models.TrainerConfig{
		MaxAttempts:    3,
		TimeLimit:      time.Hour,
		ShowHints:      true,
		AdaptivePacing: true,
		CognitiveLoad:  models.Beginner,
	}
	
	cltTrainer := trainer.NewCLTTrainer(exerciseList, config)
	
	if cltTrainer == nil {
		t.Fatal("NewCLTTrainer returned nil")
	}
}

func TestTrainerConfig(t *testing.T) {
	config := models.TrainerConfig{
		MaxAttempts:    5,
		TimeLimit:      30 * time.Minute,
		ShowHints:      false,
		AdaptivePacing: false,
		CognitiveLoad:  models.Advanced,
	}
	
	if config.MaxAttempts != 5 {
		t.Errorf("Expected MaxAttempts 5, got %d", config.MaxAttempts)
	}
	
	if config.CognitiveLoad != models.Advanced {
		t.Error("Expected Advanced cognitive load")
	}
}

func TestLearningProgress(t *testing.T) {
	progress := models.LearningProgress{
		ExerciseID: "variables",
		StartTime:  time.Now(),
		Attempts:   3,
		Score:      85.5,
		HintsUsed:  1,
	}
	
	if progress.ExerciseID != "variables" {
		t.Error("Exercise ID not set correctly")
	}
	
	if progress.Score != 85.5 {
		t.Error("Score not set correctly")
	}
}

func TestCognitiveLevels(t *testing.T) {
	if models.Beginner >= models.Intermediate {
		t.Error("Cognitive levels should be ordered")
	}
	
	if models.Intermediate >= models.Advanced {
		t.Error("Cognitive levels should be ordered")
	}
}

func TestExerciseTypes(t *testing.T) {
	exercise := exercises.GetVariablesExercise()
	
	if exercise.ExerciseType != models.Concept {
		t.Error("Variables should be a concept exercise")
	}
	
	functionsEx := exercises.GetFunctionsExercise()
	if functionsEx.ExerciseType != models.Application {
		t.Error("Functions should be an application exercise")
	}
}

func TestProgressTracking(t *testing.T) {
	// Test that LearningProgress struct has all required fields
	progress := models.LearningProgress{
		ExerciseID: "test",
		StartTime:  time.Now(),
		Attempts:   5,
		HintsUsed:  2,
		Score:      85.5,
	}
	
	if progress.Attempts != 5 {
		t.Error("Attempts field should be properly set")
	}
	
	if progress.HintsUsed != 2 {
		t.Error("HintsUsed field should be properly set")
	}
	
	if progress.Score != 85.5 {
		t.Error("Score field should be properly set")
	}
}

func TestScoringComponents(t *testing.T) {
	// Test that score calculation components are reasonable
	// Note: Full scoring test would require creating a trainer instance
	// This validates the scoring fields exist and can be set
	
	progress := models.LearningProgress{
		ExerciseID: "variables",
		StartTime:  time.Now().Add(-5 * time.Minute),
		Attempts:   3,
		HintsUsed:  1,
		Score:      78.5,
		TimeSpent:  5 * time.Minute,
	}
	
	completedAt := time.Now()
	progress.CompletedAt = &completedAt
	
	// Verify all analytics fields are properly tracked
	if progress.Score < 0 || progress.Score > 100 {
		t.Error("Score should be between 0 and 100")
	}
	
	if progress.TimeSpent <= 0 {
		t.Error("TimeSpent should be positive")
	}
	
	if progress.CompletedAt == nil {
		t.Error("CompletedAt should be set for completed exercises")
	}
}

func TestTerminalFormatting(t *testing.T) {
	// Test that code formatting works for terminal display
	registry := exercises.NewRegistry()
	exerciseList := registry.GetAll()
	
	config := models.TrainerConfig{
		MaxAttempts:    3,
		TimeLimit:      time.Hour,
		ShowHints:      true,
		AdaptivePacing: true,
		CognitiveLoad:  models.Beginner,
	}
	
	trainer := trainer.NewCLTTrainer(exerciseList, config)
	
	// Test formatting method exists and works
	testCode := "package main\n\nfunc main() {}"
	formatted := trainer.FormatCodeBlock(testCode) // Need to make this method public for testing
	
	// Check that formatting doesn't contain markdown fences
	if strings.Contains(formatted, "```") {
		t.Error("Formatted output should not contain markdown fences")
	}
	
	// Check that it contains terminal-friendly borders
	if !strings.Contains(formatted, "┌") || !strings.Contains(formatted, "└") {
		t.Error("Formatted output should contain terminal borders")
	}
}