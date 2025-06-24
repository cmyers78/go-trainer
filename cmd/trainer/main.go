package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/cmyers78/claude/internal/exercises"
	"github.com/cmyers78/claude/internal/models"
	"github.com/cmyers78/claude/internal/storage"
	"github.com/cmyers78/claude/internal/trainer"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "resume" {
		handleResume()
		return
	}

	if len(os.Args) > 1 && os.Args[1] == "list" {
		handleList()
		return
	}

	if len(os.Args) > 1 && os.Args[1] == "delete" {
		handleDelete()
		return
	}

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

	// Setup session storage
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Error getting home directory: %v\n", err)
		os.Exit(1)
	}
	
	sessionStorage := storage.NewFileSessionStorage(filepath.Join(homeDir, ".claude-trainer", "sessions"))
	userID := "default" // In a real app, this would be from authentication

	// Create and start the CLT-based trainer
	cltTrainer := trainer.NewCLTTrainer(exerciseList, config, userID, sessionStorage)
	cltTrainer.Start()
}

// handleResume manages session resumption
func handleResume() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Error getting home directory: %v\n", err)
		os.Exit(1)
	}
	
	sessionStorage := storage.NewFileSessionStorage(filepath.Join(homeDir, ".claude-trainer", "sessions"))
	userID := "default"
	
	// List available sessions to resume
	sessions, err := trainer.ListUserSessions(userID, sessionStorage)
	if err != nil {
		fmt.Printf("Error listing sessions: %v\n", err)
		os.Exit(1)
	}
	
	var pausedSessions []*models.TrainingSession
	for _, session := range sessions {
		if session.Status == models.SessionPaused {
			pausedSessions = append(pausedSessions, session)
		}
	}
	
	if len(pausedSessions) == 0 {
		fmt.Println("No paused training sessions found.")
		return
	}
	
	if len(pausedSessions) == 1 {
		// Resume the only paused session
		resumeSession(pausedSessions[0].SessionID, sessionStorage)
		return
	}
	
	// Multiple sessions - let user choose
	fmt.Println("Multiple paused sessions found:")
	for i, session := range pausedSessions {
		fmt.Printf("%d. Session from %s (Exercise %d)\n", 
			i+1, session.PausedAt.Format("2006-01-02 15:04"), session.CurrentIndex+1)
	}
	
	fmt.Print("Choose session to resume (1-", len(pausedSessions), "): ")
	var choice int
	if _, err := fmt.Scanf("%d", &choice); err != nil || choice < 1 || choice > len(pausedSessions) {
		fmt.Println("Invalid choice")
		os.Exit(1)
	}
	
	resumeSession(pausedSessions[choice-1].SessionID, sessionStorage)
}

// resumeSession continues a specific training session
func resumeSession(sessionID string, sessionStorage storage.SessionStorage) {
	registry := exercises.NewRegistry()
	exerciseList := registry.GetAll()
	
	cltTrainer, err := trainer.ResumeSession(sessionID, exerciseList, sessionStorage)
	if err != nil {
		fmt.Printf("Error resuming session: %v\n", err)
		os.Exit(1)
	}
	
	cltTrainer.Start()
}

// handleList shows all user sessions
func handleList() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Error getting home directory: %v\n", err)
		os.Exit(1)
	}
	
	sessionStorage := storage.NewFileSessionStorage(filepath.Join(homeDir, ".claude-trainer", "sessions"))
	userID := "default"
	
	sessions, err := trainer.ListUserSessions(userID, sessionStorage)
	if err != nil {
		fmt.Printf("Error listing sessions: %v\n", err)
		os.Exit(1)
	}
	
	if len(sessions) == 0 {
		fmt.Println("No training sessions found.")
		return
	}
	
	fmt.Println("Training Sessions:")
	fmt.Println("==================")
	
	for _, session := range sessions {
		fmt.Printf("Session ID: %s\n", session.SessionID)
		fmt.Printf("Status: %s\n", session.Status)
		fmt.Printf("Started: %s\n", session.StartTime.Format("2006-01-02 15:04:05"))
		
		if session.PausedAt != nil {
			fmt.Printf("Paused: %s\n", session.PausedAt.Format("2006-01-02 15:04:05"))
		}
		
		fmt.Printf("Progress: Exercise %d/%d\n", session.CurrentIndex+1, len(session.Progress))
		fmt.Println("---")
	}
}

// handleDelete manages session deletion
func handleDelete() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Error getting home directory: %v\n", err)
		os.Exit(1)
	}
	
	sessionStorage := storage.NewFileSessionStorage(filepath.Join(homeDir, ".claude-trainer", "sessions"))
	userID := "default"
	
	sessions, err := trainer.ListUserSessions(userID, sessionStorage)
	if err != nil {
		fmt.Printf("Error listing sessions: %v\n", err)
		os.Exit(1)
	}
	
	if len(sessions) == 0 {
		fmt.Println("No training sessions found.")
		return
	}
	
	fmt.Println("Select session to delete:")
	fmt.Println("========================")
	
	for i, session := range sessions {
		status := string(session.Status)
		if session.PausedAt != nil {
			status += fmt.Sprintf(" (paused %s)", session.PausedAt.Format("2006-01-02 15:04"))
		}
		fmt.Printf("%d. %s - %s - Exercise %d\n", 
			i+1, session.SessionID, status, session.CurrentIndex+1)
	}
	
	fmt.Print("Choose session to delete (1-", len(sessions), ") or 0 to cancel: ")
	var choice int
	if _, err := fmt.Scanf("%d", &choice); err != nil {
		fmt.Println("Invalid input")
		os.Exit(1)
	}
	
	if choice == 0 {
		fmt.Println("Cancelled.")
		return
	}
	
	if choice < 1 || choice > len(sessions) {
		fmt.Println("Invalid choice")
		os.Exit(1)
	}
	
	sessionToDelete := sessions[choice-1]
	
	// Confirm deletion
	fmt.Printf("Are you sure you want to delete session '%s'? (y/N): ", sessionToDelete.SessionID)
	var confirm string
	fmt.Scanf("%s", &confirm)
	
	if confirm != "y" && confirm != "Y" {
		fmt.Println("Cancelled.")
		return
	}
	
	if err := sessionStorage.DeleteSession(sessionToDelete.SessionID); err != nil {
		fmt.Printf("Error deleting session: %v\n", err)
		os.Exit(1)
	}
	
	fmt.Printf("Session '%s' deleted successfully.\n", sessionToDelete.SessionID)
}