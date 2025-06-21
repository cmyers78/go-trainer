package main

import (
	"testing"
)

func TestNewTrainer(t *testing.T) {
	trainer := NewTrainer()
	
	if trainer == nil {
		t.Fatal("NewTrainer returned nil")
	}
	
	if trainer.current != 0 {
		t.Errorf("Expected current to be 0, got %d", trainer.current)
	}
	
	if trainer.score != 0 {
		t.Errorf("Expected score to be 0, got %d", trainer.score)
	}
	
	if len(trainer.exercises) == 0 {
		t.Error("Expected exercises to be populated")
	}
}

func TestCheckSolution(t *testing.T) {
	trainer := NewTrainer()
	
	// Test valid solution for first exercise (variables)
	result := trainer.checkSolution(`var name string = "John"`)
	if !result {
		t.Error("Valid solution was rejected")
	}
	
	// Test invalid solution
	result = trainer.checkSolution(`name := "John"`)
	if result {
		t.Error("Invalid solution was accepted")
	}
	
	// Test out of bounds
	trainer.current = len(trainer.exercises)
	result = trainer.checkSolution(`var name string = "John"`)
	if result {
		t.Error("Out of bounds check failed")
	}
}

func TestTrainerProgression(t *testing.T) {
	trainer := NewTrainer()
	
	originalCurrent := trainer.current
	originalScore := trainer.score
	
	// Simulate correct answer
	if trainer.checkSolution(`var name string = "John"`) {
		trainer.score++
		trainer.current++
	}
	
	if trainer.current != originalCurrent+1 {
		t.Errorf("Expected current to increment from %d to %d, got %d", 
			originalCurrent, originalCurrent+1, trainer.current)
	}
	
	if trainer.score != originalScore+1 {
		t.Errorf("Expected score to increment from %d to %d, got %d", 
			originalScore, originalScore+1, trainer.score)
	}
}

func TestTrainerBounds(t *testing.T) {
	trainer := NewTrainer()
	
	// Test within bounds
	if trainer.current >= len(trainer.exercises) {
		t.Error("Trainer should start within bounds")
	}
	
	// Test progression to completion
	trainer.current = len(trainer.exercises) - 1
	trainer.score = len(trainer.exercises) - 1
	
	// Simulate final correct answer
	if trainer.checkSolution(trainer.exercises[trainer.current].Solution) {
		trainer.score++
		trainer.current++
	}
	
	// Should now be at completion
	if trainer.current != len(trainer.exercises) {
		t.Error("Trainer should complete when all exercises are done")
	}
	
	if trainer.score != len(trainer.exercises) {
		t.Error("Score should equal exercise count on perfect completion")
	}
}