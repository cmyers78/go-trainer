package unit

import (
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