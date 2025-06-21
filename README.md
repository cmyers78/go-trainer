# Go Trainer with Cognitive Load Theory

An interactive command-line trainer for learning Go programming fundamentals, built using proven learning science principles from Cognitive Load Theory (CLT).

## Features

- **Worked Examples First** - Study complete solutions before practicing
- **Progressive Disclosure** - Complex concepts broken into manageable chunks  
- **Multiple Practice Challenges** - Faded guidance from scaffolded to independent
- **Adaptive Feedback** - Context-aware hints and support
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
2. **Functions** - Learn to create and use functions effectively

## Usage

```bash
go run cmd/trainer/main.go
```

## Commands

- `help` - Show available commands
- `hint` - Get step-by-step guidance
- `skip` - Skip current challenge and see solution
- `quit` - Exit the trainer

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
│   ├── trainer/          # CLT-based training logic
│   └── ui/               # User interface components
├── pkg/clt/              # Cognitive Load Theory utilities
└── tests/                # Test organization
    ├── unit/             # Unit tests
    ├── integration/      # Integration tests
    └── benchmark/        # Performance benchmarks
```

### Key Components

- **Models** - Domain entities with CLT-specific fields (cognitive level, exercise type)
- **Exercises** - Learning modules with worked examples and progressive challenges
- **Trainer** - CLT implementation with adaptive pacing and feedback
- **CLT Package** - Reusable cognitive load theory utilities