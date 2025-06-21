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

// Trainer manages the learning session
type Trainer struct {
	Exercises    []Exercise
	Progress     []LearningProgress
	CurrentIndex int
	StartTime    time.Time
}

// TrainerConfig holds configuration for the training session
type TrainerConfig struct {
	MaxAttempts     int
	TimeLimit       time.Duration
	ShowHints       bool
	AdaptivePacing  bool
	CognitiveLoad   CognitiveLevel
}