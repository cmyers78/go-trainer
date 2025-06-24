package unit

import (
	"strings"
	"testing"
	"time"

	"github.com/cmyers78/claude/internal/exercises"
	"github.com/cmyers78/claude/internal/models"
	"github.com/cmyers78/claude/internal/storage"
	"github.com/cmyers78/claude/internal/trainer"
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
	
	tempDir := t.TempDir()
	sessionStorage := storage.NewFileSessionStorage(tempDir)
	cltTrainer := trainer.NewCLTTrainer(exerciseList, config, "test-user", sessionStorage)
	
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
	
	tempDir := t.TempDir()
	sessionStorage := storage.NewFileSessionStorage(tempDir)
	trainer := trainer.NewCLTTrainer(exerciseList, config, "test-user", sessionStorage)
	
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

func TestCLTTrainerWithSessionStorage(t *testing.T) {
	tempDir := t.TempDir()
	sessionStorage := storage.NewFileSessionStorage(tempDir)
	
	registry := exercises.NewRegistry()
	exerciseList := registry.GetAll()
	
	config := models.TrainerConfig{
		MaxAttempts:    3,
		TimeLimit:      time.Hour,
		ShowHints:      true,
		AdaptivePacing: true,
		CognitiveLoad:  models.Beginner,
	}
	
	trainer := trainer.NewCLTTrainer(exerciseList, config, "test-user", sessionStorage)
	
	if trainer == nil {
		t.Fatal("NewCLTTrainer with session storage returned nil")
	}
}

func TestTrainerSessionResume(t *testing.T) {
	tempDir := t.TempDir()
	sessionStorage := storage.NewFileSessionStorage(tempDir)
	
	registry := exercises.NewRegistry()
	exerciseList := registry.GetAll()
	
	// Create a session to resume
	now := time.Now()
	session := &models.TrainingSession{
		UserID:       "test-user",
		SessionID:    "test-session-resume",
		CurrentIndex: 1, // Start at second exercise
		StartTime:    now,
		LastActivity: now,
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
				StartTime:  now,
				Attempts:   2,
				Score:      85.0,
				TimeSpent:  time.Minute * 3,
				HintsUsed:  1,
			},
		},
	}
	
	// Save the session
	err := sessionStorage.SaveSession(session)
	if err != nil {
		t.Fatalf("Failed to save test session: %v", err)
	}
	
	// Test resuming the session
	resumedTrainer, err := trainer.ResumeSession("test-session-resume", exerciseList, sessionStorage)
	if err != nil {
		t.Fatalf("Failed to resume session: %v", err)
	}
	
	if resumedTrainer == nil {
		t.Fatal("ResumeSession returned nil trainer")
	}
	
	// Verify session status was updated to active
	updatedSession, err := sessionStorage.LoadSession("test-session-resume")
	if err != nil {
		t.Fatalf("Failed to load updated session: %v", err)
	}
	
	if updatedSession.Status != models.SessionActive {
		t.Errorf("Session status should be active after resume, got %s", updatedSession.Status)
	}
	
	if updatedSession.PausedAt != nil {
		t.Error("PausedAt should be nil after resume")
	}
}

func TestResumeNonExistentSession(t *testing.T) {
	tempDir := t.TempDir()
	sessionStorage := storage.NewFileSessionStorage(tempDir)
	
	registry := exercises.NewRegistry()
	exerciseList := registry.GetAll()
	
	_, err := trainer.ResumeSession("non-existent", exerciseList, sessionStorage)
	if err == nil {
		t.Fatal("Expected error when resuming non-existent session")
	}
}

func TestResumeNonPausedSession(t *testing.T) {
	tempDir := t.TempDir()
	sessionStorage := storage.NewFileSessionStorage(tempDir)
	
	registry := exercises.NewRegistry()
	exerciseList := registry.GetAll()
	
	// Create an active session (not paused)
	session := &models.TrainingSession{
		UserID:    "test-user",
		SessionID: "active-session",
		Status:    models.SessionActive, // Not paused
		StartTime: time.Now(),
		Config: models.TrainerConfig{
			MaxAttempts:    3,
			TimeLimit:      time.Hour,
			ShowHints:      true,
			AdaptivePacing: true,
			CognitiveLoad:  models.Beginner,
		},
	}
	
	err := sessionStorage.SaveSession(session)
	if err != nil {
		t.Fatalf("Failed to save test session: %v", err)
	}
	
	// Try to resume non-paused session
	_, err = trainer.ResumeSession("active-session", exerciseList, sessionStorage)
	if err == nil {
		t.Fatal("Expected error when resuming non-paused session")
	}
}

func TestListUserSessions(t *testing.T) {
	tempDir := t.TempDir()
	sessionStorage := storage.NewFileSessionStorage(tempDir)
	
	// Create test sessions
	sessions := []*models.TrainingSession{
		{
			UserID:    "user1",
			SessionID: "session1",
			Status:    models.SessionPaused,
			StartTime: time.Now(),
		},
		{
			UserID:    "user1",
			SessionID: "session2",
			Status:    models.SessionCompleted,
			StartTime: time.Now(),
		},
		{
			UserID:    "user2",
			SessionID: "session3",
			Status:    models.SessionPaused,
			StartTime: time.Now(),
		},
	}
	
	// Save all sessions
	for _, session := range sessions {
		if err := sessionStorage.SaveSession(session); err != nil {
			t.Fatalf("Failed to save session %s: %v", session.SessionID, err)
		}
	}
	
	// Test listing sessions for user1
	user1Sessions, err := trainer.ListUserSessions("user1", sessionStorage)
	if err != nil {
		t.Fatalf("Failed to list sessions for user1: %v", err)
	}
	
	if len(user1Sessions) != 2 {
		t.Errorf("Expected 2 sessions for user1, got %d", len(user1Sessions))
	}
	
	// Test listing sessions for non-existent user
	noSessions, err := trainer.ListUserSessions("non-existent", sessionStorage)
	if err != nil {
		t.Fatalf("Failed to list sessions for non-existent user: %v", err)
	}
	
	if len(noSessions) != 0 {
		t.Errorf("Expected 0 sessions for non-existent user, got %d", len(noSessions))
	}
}

func TestNewCLTTrainerFromSession(t *testing.T) {
	tempDir := t.TempDir()
	sessionStorage := storage.NewFileSessionStorage(tempDir)
	
	registry := exercises.NewRegistry()
	exerciseList := registry.GetAll()
	
	// Create test session
	session := &models.TrainingSession{
		UserID:       "test-user",
		SessionID:    "test-session",
		CurrentIndex: 2,
		StartTime:    time.Now().Add(-time.Hour),
		Status:       models.SessionPaused,
		Config: models.TrainerConfig{
			MaxAttempts:    5,
			TimeLimit:      30 * time.Minute,
			ShowHints:      false,
			AdaptivePacing: false,
			CognitiveLoad:  models.Advanced,
		},
		Progress: []models.LearningProgress{
			{ExerciseID: "variables", Score: 95.0},
			{ExerciseID: "basic-types", Score: 88.0},
		},
	}
	
	// Create trainer from session
	trainer := trainer.NewCLTTrainerFromSession(session, exerciseList, sessionStorage)
	
	if trainer == nil {
		t.Fatal("NewCLTTrainerFromSession returned nil")
	}
	
	// Verify trainer was created with session data
	// Note: We can't directly access private fields, but we can test the functionality
	// by ensuring the trainer can be created without error
}