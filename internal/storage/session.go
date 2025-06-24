package storage

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/cmyers78/claude/internal/models"
)

// SessionStorage handles persistence of training sessions
type SessionStorage interface {
	SaveSession(session *models.TrainingSession) error
	LoadSession(sessionID string) (*models.TrainingSession, error)
	ListSessions(userID string) ([]*models.TrainingSession, error)
	DeleteSession(sessionID string) error
}

// FileSessionStorage implements SessionStorage using local file system
type FileSessionStorage struct {
	basePath string
}

// NewFileSessionStorage creates a new file-based session storage
func NewFileSessionStorage(basePath string) *FileSessionStorage {
	return &FileSessionStorage{
		basePath: basePath,
	}
}

// SaveSession saves a training session to disk
func (fs *FileSessionStorage) SaveSession(session *models.TrainingSession) error {
	if err := fs.ensureBasePath(); err != nil {
		return fmt.Errorf("failed to ensure base path: %w", err)
	}

	session.LastActivity = time.Now()
	
	data, err := json.MarshalIndent(session, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal session: %w", err)
	}

	filename := fmt.Sprintf("%s.json", session.SessionID)
	path := filepath.Join(fs.basePath, filename)
	
	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("failed to write session file: %w", err)
	}

	return nil
}

// LoadSession loads a training session from disk
func (fs *FileSessionStorage) LoadSession(sessionID string) (*models.TrainingSession, error) {
	filename := fmt.Sprintf("%s.json", sessionID)
	path := filepath.Join(fs.basePath, filename)
	
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("session not found: %s", sessionID)
		}
		return nil, fmt.Errorf("failed to read session file: %w", err)
	}

	var session models.TrainingSession
	if err := json.Unmarshal(data, &session); err != nil {
		return nil, fmt.Errorf("failed to unmarshal session: %w", err)
	}

	return &session, nil
}

// ListSessions returns all sessions for a given user
func (fs *FileSessionStorage) ListSessions(userID string) ([]*models.TrainingSession, error) {
	if err := fs.ensureBasePath(); err != nil {
		return nil, fmt.Errorf("failed to ensure base path: %w", err)
	}

	files, err := os.ReadDir(fs.basePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read sessions directory: %w", err)
	}

	var sessions []*models.TrainingSession
	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".json" {
			sessionID := file.Name()[:len(file.Name())-5] // Remove .json extension
			session, err := fs.LoadSession(sessionID)
			if err != nil {
				// Log corrupted session files for troubleshooting
				log.Printf("Warning: Skipping corrupted session file %s: %v", file.Name(), err)
				continue
			}
			if session.UserID == userID {
				sessions = append(sessions, session)
			}
		}
	}

	return sessions, nil
}

// DeleteSession removes a session from disk
func (fs *FileSessionStorage) DeleteSession(sessionID string) error {
	filename := fmt.Sprintf("%s.json", sessionID)
	path := filepath.Join(fs.basePath, filename)
	
	if err := os.Remove(path); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to delete session file: %w", err)
	}

	return nil
}

// ensureBasePath creates the base directory if it doesn't exist
func (fs *FileSessionStorage) ensureBasePath() error {
	if err := os.MkdirAll(fs.basePath, 0755); err != nil {
		return fmt.Errorf("failed to create base path: %w", err)
	}
	return nil
}