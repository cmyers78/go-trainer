package models

import "time"

// LearningProgress tracks a learner's progress through exercises
type LearningProgress struct {
	ExerciseID    string
	StartTime     time.Time
	CompletedAt   *time.Time
	Attempts      int
	Score         float64
	TimeSpent     time.Duration
	HintsUsed     int
}


// TrainerConfig holds configuration for the training session
type TrainerConfig struct {
	MaxAttempts     int
	TimeLimit       time.Duration
	ShowHints       bool
	AdaptivePacing  bool
	CognitiveLoad   CognitiveLevel
}

// TrainingSession represents a saved training session that can be resumed
type TrainingSession struct {
	UserID         string              `json:"user_id"`
	SessionID      string              `json:"session_id"`
	Config         TrainerConfig       `json:"config"`
	Progress       []LearningProgress  `json:"progress"`
	CurrentIndex   int                 `json:"current_index"`
	StartTime      time.Time           `json:"start_time"`
	LastActivity   time.Time           `json:"last_activity"`
	PausedAt       *time.Time          `json:"paused_at,omitempty"`
	Status         SessionStatus       `json:"status"`
}

// SessionStatus represents the current state of a training session
type SessionStatus string

const (
	SessionActive    SessionStatus = "active"
	SessionPaused    SessionStatus = "paused"
	SessionCompleted SessionStatus = "completed"
	SessionAbandoned SessionStatus = "abandoned"
)