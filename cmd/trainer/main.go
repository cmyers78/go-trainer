package main

import (
	"time"

	"github.com/cmyers78/go-trainer/internal/exercises"
	"github.com/cmyers78/go-trainer/internal/models"
	"github.com/cmyers78/go-trainer/internal/trainer"
)

func main() {
	// Initialize exercise registry
	registry := exercises.NewRegistry()
	exerciseList := registry.GetAll()
	
	// Configure trainer with CLT principles
	config := models.TrainerConfig{
		MaxAttempts:    3,
		TimeLimit:      time.Hour,
		ShowHints:      true,
		AdaptivePacing: true,
		CognitiveLoad:  models.Beginner,
	}
	
	// Create and start the CLT-based trainer
	cltTrainer := trainer.NewCLTTrainer(exerciseList, config)
	cltTrainer.Start()
}