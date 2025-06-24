# Go Trainer with Cognitive Load Theory

An interactive command-line trainer for learning Go programming fundamentals, built using proven learning science principles from Cognitive Load Theory (CLT).

## Features

- **Worked Examples First** - Study complete solutions before practicing
- **Progressive Disclosure** - Complex concepts broken into manageable chunks  
- **Multiple Practice Challenges** - Faded guidance from scaffolded to independent
- **Adaptive Feedback** - Context-aware hints and support
- **Session Management** - Pause and resume training sessions seamlessly
- **Learning Analytics** - Track time, attempts, and progress

## Cognitive Load Theory Principles Applied

### 🧠 Intrinsic Load Management
- Prerequisites clearly defined
- Single learning objective per exercise
- Concepts introduced in logical progression

### 🎯 Extraneous Load Reduction  
- Clean, consistent interface
- Clear learning goals stated upfront
- Minimal cognitive distractions

### 💡 Germane Load Optimization
- Worked examples demonstrate patterns
- Multiple contexts for same concept
- Elaborative feedback reinforces learning

## Topics Covered

1. **Variables and Types** - Master Go's variable declarations and type system
2. **Basic Data Types** - Understand numeric types, strings, constants, and type conversions
3. **Composite Types** - Work with arrays, slices, and maps effectively
4. **Functions** - Learn to create and use functions with parameters and return values  
5. **Structs and Methods** - Define custom types with methods and embedding

## Usage

### Start New Training Session
```bash
go run cmd/trainer/main.go
```

### Resume Paused Session
```bash
go run cmd/trainer/main.go resume
```

### View All Sessions
```bash
go run cmd/trainer/main.go list
```

### Delete Old Sessions
```bash
go run cmd/trainer/main.go delete
```

## Commands

During training challenges, use these commands:

- `help` - Show available commands
- `hint` - Get step-by-step guidance
- `skip` - Skip current challenge and see solution
- `pause` - Save progress and exit (resume later)
- `quit` - Exit without saving progress

## Session Management

Training sessions are automatically saved to `~/.claude-trainer/sessions/` and include:

- Current exercise position
- Progress and scores for completed exercises
- Time spent and attempts made
- Hints used and configuration settings

Sessions persist across application restarts, allowing you to pause training at any time and resume exactly where you left off.

## Testing

Run all tests:
```bash
go test
```

Run tests with coverage:
```bash
go test -cover
```

Run benchmarks:
```bash
go test -bench=.
```

Run specific test:
```bash
go test -run TestVariablesValidator
```

## Architecture

```
├── cmd/trainer/           # Application entry points
├── internal/              # Private application code
│   ├── models/           # Core data structures (Exercise, Trainer, Config)
│   ├── exercises/        # Exercise definitions and registry
│   ├── storage/          # Session persistence and storage
│   └── trainer/          # CLT-based training logic
└── tests/                # Test organization
    ├── unit/             # Unit tests
    ├── integration/      # Integration tests
    └── benchmark/        # Performance benchmarks
```

### Key Components

- **Models** - Domain entities with CLT-specific fields (cognitive level, exercise type, training sessions)
- **Exercises** - Learning modules with worked examples and progressive challenges  
- **Storage** - File-based session persistence with JSON serialization
- **Trainer** - CLT implementation with adaptive pacing, feedback, scoring, and session management
- **Tests** - Comprehensive validation including CLT principle adherence and session operations