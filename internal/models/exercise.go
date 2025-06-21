package models

// CognitiveLevel represents the complexity level based on Cognitive Load Theory
type CognitiveLevel int

const (
	Beginner CognitiveLevel = iota
	Intermediate
	Advanced
)

// ExerciseType categorizes the type of learning exercise
type ExerciseType int

const (
	Concept ExerciseType = iota  // Understanding concepts
	Application                  // Applying knowledge
	Synthesis                   // Combining concepts
)

// Example represents a single learning example
type Example struct {
	Title       string
	Code        string
	Explanation string
	Output      string
}

// Challenge represents a practice challenge
type Challenge struct {
	Description string
	Template    string
	Solution    string
	Hints       []string
	Validator   func(string) bool
}

// Exercise represents a complete learning module
type Exercise struct {
	ID              string
	Title           string
	Description     string
	CognitiveLevel  CognitiveLevel
	ExerciseType    ExerciseType
	Prerequisites   []string
	LearningGoals   []string
	Examples        []Example
	Challenges      []Challenge
	EstimatedTime   int // minutes
}