package unit

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/cmyers78/claude/internal/models"
	"github.com/cmyers78/claude/internal/storage"
)

func TestFileSessionStorage_SaveAndLoad(t *testing.T) {
	// Setup temporary directory for testing
	tempDir := t.TempDir()
	storage := storage.NewFileSessionStorage(tempDir)

	// Create test session
	now := time.Now()
	session := &models.TrainingSession{
		UserID:       "test-user",
		SessionID:    "test-session-123",
		CurrentIndex: 2,
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
				HintsUsed:  1,
			},
		},
	}

	// Test Save
	err := storage.SaveSession(session)
	if err != nil {
		t.Fatalf("Failed to save session: %v", err)
	}

	// Verify file was created
	expectedPath := filepath.Join(tempDir, "test-session-123.json")
	if _, err := os.Stat(expectedPath); os.IsNotExist(err) {
		t.Fatalf("Session file was not created at %s", expectedPath)
	}

	// Test Load
	loadedSession, err := storage.LoadSession("test-session-123")
	if err != nil {
		t.Fatalf("Failed to load session: %v", err)
	}

	// Verify loaded session matches saved session
	if loadedSession.UserID != session.UserID {
		t.Errorf("UserID mismatch: expected %s, got %s", session.UserID, loadedSession.UserID)
	}
	if loadedSession.SessionID != session.SessionID {
		t.Errorf("SessionID mismatch: expected %s, got %s", session.SessionID, loadedSession.SessionID)
	}
	if loadedSession.CurrentIndex != session.CurrentIndex {
		t.Errorf("CurrentIndex mismatch: expected %d, got %d", session.CurrentIndex, loadedSession.CurrentIndex)
	}
	if loadedSession.Status != session.Status {
		t.Errorf("Status mismatch: expected %s, got %s", session.Status, loadedSession.Status)
	}
	if len(loadedSession.Progress) != len(session.Progress) {
		t.Errorf("Progress length mismatch: expected %d, got %d", len(session.Progress), len(loadedSession.Progress))
	}
}

func TestFileSessionStorage_LoadNonExistentSession(t *testing.T) {
	tempDir := t.TempDir()
	storage := storage.NewFileSessionStorage(tempDir)

	_, err := storage.LoadSession("non-existent")
	if err == nil {
		t.Fatal("Expected error when loading non-existent session, got nil")
	}
}

func TestFileSessionStorage_ListSessions(t *testing.T) {
	tempDir := t.TempDir()
	storage := storage.NewFileSessionStorage(tempDir)

	// Create test sessions for different users
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
			Status:    models.SessionActive,
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
		if err := storage.SaveSession(session); err != nil {
			t.Fatalf("Failed to save session %s: %v", session.SessionID, err)
		}
	}

	// Test listing sessions for user1
	user1Sessions, err := storage.ListSessions("user1")
	if err != nil {
		t.Fatalf("Failed to list sessions for user1: %v", err)
	}

	if len(user1Sessions) != 2 {
		t.Errorf("Expected 2 sessions for user1, got %d", len(user1Sessions))
	}

	// Verify sessions belong to user1
	for _, session := range user1Sessions {
		if session.UserID != "user1" {
			t.Errorf("Found session for wrong user: expected user1, got %s", session.UserID)
		}
	}

	// Test listing sessions for user2
	user2Sessions, err := storage.ListSessions("user2")
	if err != nil {
		t.Fatalf("Failed to list sessions for user2: %v", err)
	}

	if len(user2Sessions) != 1 {
		t.Errorf("Expected 1 session for user2, got %d", len(user2Sessions))
	}
}

func TestFileSessionStorage_DeleteSession(t *testing.T) {
	tempDir := t.TempDir()
	storage := storage.NewFileSessionStorage(tempDir)

	// Create and save test session
	session := &models.TrainingSession{
		UserID:    "test-user",
		SessionID: "test-session-delete",
		Status:    models.SessionPaused,
		StartTime: time.Now(),
	}

	err := storage.SaveSession(session)
	if err != nil {
		t.Fatalf("Failed to save session: %v", err)
	}

	// Verify file exists
	sessionPath := filepath.Join(tempDir, "test-session-delete.json")
	if _, err := os.Stat(sessionPath); os.IsNotExist(err) {
		t.Fatalf("Session file should exist before deletion")
	}

	// Delete session
	err = storage.DeleteSession("test-session-delete")
	if err != nil {
		t.Fatalf("Failed to delete session: %v", err)
	}

	// Verify file is deleted
	if _, err := os.Stat(sessionPath); !os.IsNotExist(err) {
		t.Errorf("Session file should be deleted")
	}

	// Test deleting non-existent session (should not error)
	err = storage.DeleteSession("non-existent")
	if err != nil {
		t.Errorf("Deleting non-existent session should not error, got: %v", err)
	}
}

func TestTrainingSession_StatusConstants(t *testing.T) {
	testCases := []struct {
		status   models.SessionStatus
		expected string
	}{
		{models.SessionActive, "active"},
		{models.SessionPaused, "paused"},
		{models.SessionCompleted, "completed"},
		{models.SessionAbandoned, "abandoned"},
	}

	for _, tc := range testCases {
		if string(tc.status) != tc.expected {
			t.Errorf("Status constant mismatch: expected %s, got %s", tc.expected, string(tc.status))
		}
	}
}

func TestTrainingSession_JSONSerialization(t *testing.T) {
	now := time.Now()
	pausedAt := now.Add(time.Minute)
	
	session := &models.TrainingSession{
		UserID:       "test-user",
		SessionID:    "test-session",
		CurrentIndex: 1,
		StartTime:    now,
		LastActivity: now,
		PausedAt:     &pausedAt,
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
				Attempts:   1,
				Score:      90.0,
				TimeSpent:  time.Minute * 5,
				HintsUsed:  0,
			},
		},
	}

	// Test with temporary storage to verify JSON serialization works
	tempDir := t.TempDir()
	storage := storage.NewFileSessionStorage(tempDir)

	// Save (serializes to JSON)
	err := storage.SaveSession(session)
	if err != nil {
		t.Fatalf("Failed to serialize session to JSON: %v", err)
	}

	// Load (deserializes from JSON)
	loadedSession, err := storage.LoadSession("test-session")
	if err != nil {
		t.Fatalf("Failed to deserialize session from JSON: %v", err)
	}

	// Verify complex fields are preserved
	if loadedSession.PausedAt == nil {
		t.Error("PausedAt should not be nil after deserialization")
	} else if !loadedSession.PausedAt.Equal(pausedAt) {
		t.Errorf("PausedAt mismatch: expected %v, got %v", pausedAt, *loadedSession.PausedAt)
	}

	if loadedSession.Config.TimeLimit != time.Hour {
		t.Errorf("Config.TimeLimit mismatch: expected %v, got %v", time.Hour, loadedSession.Config.TimeLimit)
	}

	if len(loadedSession.Progress) != 1 {
		t.Fatalf("Expected 1 progress entry, got %d", len(loadedSession.Progress))
	}

	progress := loadedSession.Progress[0]
	if progress.TimeSpent != time.Minute*5 {
		t.Errorf("Progress.TimeSpent mismatch: expected %v, got %v", time.Minute*5, progress.TimeSpent)
	}
}