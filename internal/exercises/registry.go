package exercises

import "github.com/cmyers78/claude/internal/models"

// Registry holds all available exercises
type Registry struct {
	exercises map[string]models.Exercise
}

// NewRegistry creates a new exercise registry
func NewRegistry() *Registry {
	registry := &Registry{
		exercises: make(map[string]models.Exercise),
	}
	
	// Register all exercises
	registry.exercises["variables"] = GetVariablesExercise()
	registry.exercises["basic-types"] = GetBasicTypesExercise()
	registry.exercises["composite-types"] = GetCompositeTypesExercise()
	registry.exercises["functions"] = GetFunctionsExercise()
	registry.exercises["structs"] = GetStructsExercise()
	
	return registry
}

// GetAll returns all exercises in learning order
func (r *Registry) GetAll() []models.Exercise {
	// Return exercises in pedagogical order
	order := []string{"variables", "basic-types", "composite-types", "functions", "structs"}
	
	var exercises []models.Exercise
	for _, id := range order {
		if exercise, exists := r.exercises[id]; exists {
			exercises = append(exercises, exercise)
		}
	}
	
	return exercises
}

// GetByID returns a specific exercise by ID
func (r *Registry) GetByID(id string) (models.Exercise, bool) {
	exercise, exists := r.exercises[id]
	return exercise, exists
}

// GetByPrerequisites returns exercises that match the given prerequisites
func (r *Registry) GetByPrerequisites(completed []string) []models.Exercise {
	var available []models.Exercise
	
	for _, exercise := range r.exercises {
		if r.hasPrerequisites(exercise.Prerequisites, completed) {
			available = append(available, exercise)
		}
	}
	
	return available
}

// hasPrerequisites checks if all prerequisites are met
func (r *Registry) hasPrerequisites(required []string, completed []string) bool {
	for _, req := range required {
		found := false
		for _, comp := range completed {
			if req == comp {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}