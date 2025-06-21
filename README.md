# Go Trainer

An interactive command-line trainer for learning Go programming fundamentals.

## Features

- 5 progressive exercises covering Go basics
- Interactive examples before each exercise
- Hint system and solution validation
- Progress tracking and scoring

## Topics Covered

1. **Variables and Types** - Learn Go's type system and variable declarations
2. **Functions** - Create functions with parameters and return values
3. **Slices** - Work with Go's dynamic arrays
4. **Structs** - Define custom data types
5. **Methods** - Add behavior to your structs

## Usage

```bash
go run .
```

## Commands

- `help` - Show available commands
- `hint` - Show solution for current exercise
- `skip` - Skip current exercise
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

- `main.go` - Application entry point
- `models.go` - Core data structures
- `exercises.go` - Exercise definitions
- `trainer.go` - Training logic and UI
- `*_test.go` - Test files with unit tests and benchmarks