package benchmark

import (
	"strconv"
	"testing"
	"time"

	"github.com/cmyers78/claude/internal/exercises"
	"github.com/cmyers78/claude/internal/models"
	"github.com/cmyers78/claude/internal/storage"
	"github.com/cmyers78/claude/internal/trainer"
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
		tempDir := b.TempDir()
		sessionStorage := storage.NewFileSessionStorage(tempDir)
		trainer.NewCLTTrainer(exerciseList, config, "test-user", sessionStorage)
	}
}

func BenchmarkSessionSave(b *testing.B) {
	tempDir := b.TempDir()
	sessionStorage := storage.NewFileSessionStorage(tempDir)
	
	session := &models.TrainingSession{
		UserID:       "bench-user",
		SessionID:    "bench-session",
		CurrentIndex: 2,
		StartTime:    time.Now(),
		LastActivity: time.Now(),
		Status:       models.SessionPaused,
		Config: models.TrainerConfig{
			MaxAttempts:    3,
			TimeLimit:      time.Hour,
			ShowHints:      true,
			AdaptivePacing: true,
			CognitiveLoad:  models.Beginner,
		},
		Progress: []models.LearningProgress{
			{
				ExerciseID: "variables",
				StartTime:  time.Now(),
				Attempts:   2,
				Score:      85.0,
				TimeSpent:  time.Minute * 3,
				HintsUsed:  1,
			},
			{
				ExerciseID: "basic-types",
				StartTime:  time.Now(),
				Attempts:   1,
				Score:      92.0,
				TimeSpent:  time.Minute * 2,
				HintsUsed:  0,
			},
		},
	}
	
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		session.SessionID = "bench-session-" + strconv.Itoa(i) // Make unique
		sessionStorage.SaveSession(session)
	}
}

func BenchmarkSessionLoad(b *testing.B) {
	tempDir := b.TempDir()
	sessionStorage := storage.NewFileSessionStorage(tempDir)
	
	// Pre-create session to load
	session := &models.TrainingSession{
		UserID:       "bench-user",
		SessionID:    "bench-load-session",
		CurrentIndex: 1,
		StartTime:    time.Now(),
		Status:       models.SessionPaused,
		Config: models.TrainerConfig{
			MaxAttempts:    3,
			TimeLimit:      time.Hour,
			ShowHints:      true,
			AdaptivePacing: true,
			CognitiveLoad:  models.Beginner,
		},
		Progress: []models.LearningProgress{
			{
				ExerciseID: "variables",
				StartTime:  time.Now(),
				Attempts:   2,
				Score:      85.0,
				TimeSpent:  time.Minute * 3,
				HintsUsed:  1,
			},
		},
	}
	
	sessionStorage.SaveSession(session)
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		sessionStorage.LoadSession("bench-load-session")
	}
}

func BenchmarkListSessions(b *testing.B) {
	tempDir := b.TempDir()
	sessionStorage := storage.NewFileSessionStorage(tempDir)
	
	// Pre-create multiple sessions
	for i := 0; i < 10; i++ {
		session := &models.TrainingSession{
			UserID:    "bench-user",
			SessionID: "session-" + strconv.Itoa(i),
			Status:    models.SessionPaused,
			StartTime: time.Now(),
		}
		sessionStorage.SaveSession(session)
	}
	
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		sessionStorage.ListSessions("bench-user")
	}
}

func BenchmarkSessionResume(b *testing.B) {
	tempDir := b.TempDir()
	sessionStorage := storage.NewFileSessionStorage(tempDir)
	
	registry := exercises.NewRegistry()
	exerciseList := registry.GetAll()
	
	// Pre-create session to resume
	session := &models.TrainingSession{
		UserID:       "bench-user",
		SessionID:    "bench-resume-session",
		CurrentIndex: 1,
		StartTime:    time.Now(),
		Status:       models.SessionPaused,
		Config: models.TrainerConfig{
			MaxAttempts:    3,
			TimeLimit:      time.Hour,
			ShowHints:      true,
			AdaptivePacing: true,
			CognitiveLoad:  models.Beginner,
		},
		Progress: []models.LearningProgress{
			{
				ExerciseID: "variables",
				StartTime:  time.Now(),
				Attempts:   1,
				Score:      88.0,
				TimeSpent:  time.Minute * 4,
				HintsUsed:  0,
			},
		},
	}
	
	sessionStorage.SaveSession(session)
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		// Reset session to paused state for each iteration
		session.Status = models.SessionPaused
		sessionStorage.SaveSession(session)
		
		trainer.ResumeSession("bench-resume-session", exerciseList, sessionStorage)
	}
}