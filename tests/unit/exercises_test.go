package unit

import (
	"testing"

	"github.com/cmyers78/go-trainer/internal/exercises"
	"github.com/cmyers78/go-trainer/internal/models"
)

func TestExerciseRegistry(t *testing.T) {
	registry := exercises.NewRegistry()
	
	// Test GetAll
	allExercises := registry.GetAll()
	if len(allExercises) == 0 {
		t.Error("Expected exercises to be populated")
	}
	
	// Test GetByID
	variablesEx, exists := registry.GetByID("variables")
	if !exists {
		t.Error("Expected variables exercise to exist")
	}
	
	if variablesEx.Title != "Variables and Types" {
		t.Errorf("Expected 'Variables and Types', got '%s'", variablesEx.Title)
	}
}

func TestVariablesExercise(t *testing.T) {
	exercise := exercises.GetVariablesExercise()
	
	// Test basic structure
	if exercise.ID != "variables" {
		t.Errorf("Expected ID 'variables', got '%s'", exercise.ID)
	}
	
	if exercise.CognitiveLevel != models.Beginner {
		t.Error("Variables should be beginner level")
	}
	
	if len(exercise.Examples) == 0 {
		t.Error("Expected examples to be provided")
	}
	
	if len(exercise.Challenges) == 0 {
		t.Error("Expected challenges to be provided")
	}
	
	if len(exercise.LearningGoals) == 0 {
		t.Error("Expected learning goals to be defined")
	}
}

func TestFunctionsExercise(t *testing.T) {
	exercise := exercises.GetFunctionsExercise()
	
	// Test prerequisites
	if len(exercise.Prerequisites) == 0 {
		t.Error("Functions should have prerequisites")
	}
	
	if exercise.Prerequisites[0] != "variables" {
		t.Error("Functions should require variables as prerequisite")
	}
	
	// Test progressive difficulty
	if len(exercise.Challenges) < 2 {
		t.Error("Should have multiple challenges for progressive learning")
	}
}

func TestChallengeValidators(t *testing.T) {
	exercise := exercises.GetVariablesExercise()
	
	// Test first challenge (explicit type)
	validator := exercise.Challenges[0].Validator
	
	testCases := []struct {
		input    string
		expected bool
	}{
		{`var name string = "John"`, true},
		{`var age int = 25`, false},     // wrong type
		{`name := "John"`, false},       // no var keyword
		{`var name string`, false},      // no assignment
	}
	
	for _, tc := range testCases {
		result := validator(tc.input)
		if result != tc.expected {
			t.Errorf("Input %q: expected %v, got %v", tc.input, tc.expected, result)
		}
	}
}

func TestCLTPrinciplesInExerciseDesign(t *testing.T) {
	exercise := exercises.GetVariablesExercise()
	
	// Test worked example effect - examples before challenges
	if len(exercise.Examples) == 0 {
		t.Error("Should provide worked examples before challenges")
	}
	
	// Test progressive disclosure - multiple examples
	if len(exercise.Examples) < 3 {
		t.Error("Should provide multiple examples for variability effect")
	}
	
	// Test completion effect - challenges have scaffolding
	for _, challenge := range exercise.Challenges {
		if challenge.Template == "" {
			t.Error("Challenges should provide templates for completion effect")
		}
		
		if len(challenge.Hints) == 0 {
			t.Error("Challenges should provide hints for guidance")
		}
	}
	
	// Test faded guidance - multiple challenges with decreasing support
	if len(exercise.Challenges) < 2 {
		t.Error("Should provide multiple challenges for faded guidance")
	}
}