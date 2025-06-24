# Claude Development Guidelines

## Git Workflow
- **ALWAYS create feature branches** - never work directly on main
- Use descriptive branch names: `feature/description`, `fix/issue-name`, `refactor/component`
- Branch naming format: `feature/pause-resume-training`, `fix/session-storage-bug`

## Development Standards
- Follow existing code patterns and conventions
- Run tests before committing (when available)
- Use meaningful commit messages
- Always add proper error handling

## Project Structure
- `cmd/trainer/` - CLI entry points
- `internal/models/` - Data models and types
- `internal/trainer/` - Core training logic
- `internal/storage/` - Persistence layer
- `internal/exercises/` - Exercise definitions

## Commands for Development
```bash
# Start new feature
git checkout -b feature/description

# Run training
go run cmd/trainer/main.go

# Resume training
go run cmd/trainer/main.go resume

# List sessions
go run cmd/trainer/main.go list

# Delete sessions
go run cmd/trainer/main.go delete
```

## Session Storage
- Location: `~/.claude-trainer/sessions/`
- Format: JSON files
- Contains: Full training state, progress, configuration