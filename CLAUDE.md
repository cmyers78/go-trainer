# Claude Development Guidelines

## Git Workflow
- **ALWAYS create feature branches** - never work directly on main
- Use descriptive branch names: `feature/description`, `fix/issue-name`, `refactor/component`
- Branch naming format: `feature/pause-resume-training`, `fix/session-storage-bug`

## Documentation Requirements for ALL PRs
**MANDATORY**: Before pushing any feature branch or creating a PR, ALWAYS update:

1. **README.md** - Add new features, usage examples, commands, and architecture changes
2. **CHANGELOG.md** - Document all changes in the [Unreleased] section following Keep a Changelog format
3. **Code Documentation** - Update inline comments and function documentation as needed

**No PR should be created without proper documentation updates.** This ensures:
- Users understand new functionality
- Changes are properly tracked and versioned
- Future contributors can understand the evolution of the codebase

## Development Standards
- Follow existing code patterns and conventions
- Run tests before committing (when available)
- Use meaningful commit messages
- Always add proper error handling
- Write comprehensive tests for all new functionality

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