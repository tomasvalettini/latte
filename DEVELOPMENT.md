# Latte Development Guide

**latte** — Little Actions Toward Tangible Éndings

A terminal-based drip tracker built with Go and the Unix philosophy.

---

## Project Structure

```
cmd/              - CLI commands (Cobra-based)
coffeeshop/       - Coffeeshop domain
  ├── controller/ - Business logic and orchestration
  ├── data/
  │   ├── data-source/  - Persistence layer (Load/Save)
  │   │   └── path/     - File path handling abstractions
  │   └── model/        - Data models and domain entities
assert/           - Custom assertion helpers
test-utils/       - Shared test utilities (e.g., RequireExit)
```

### Directory Responsibilities

- **cmd/** - Cobra command handlers that parse flags and delegate to controllers
- **controller/** - Contains business logic; orchestrates data sources and models
- **data-source/** - Handles I/O (file reading/writing, persistence)
- **model/** - Pure data structures (Blend, Drip) with JSON serialization
- **path/** - Abstracts file path logic for portability and testing

---

## Building & Running

```bash
# Build the binary
go build -o latte

# Run a command
./latte add "my drip"
./latte list
./latte delete
```

Dependencies are managed in `go.mod`. Currently uses:
- `github.com/spf13/cobra` — CLI framework

---

## Architecture & Design Patterns

### Layered Architecture

```
Cmd (CLI Layer)
    ↓
Controller (Business Logic)
    ↓
DataSource (Persistence) + Model (Domain Entities)
```

**Flow:**
1. CLI command receives user input via Cobra flags/args
2. Command instantiates a controller
3. Controller implements business logic (validation, orchestration)
4. Controller delegates to DataSource for persistence
5. DataSource loads/saves model entities to JSON files

### Key Patterns

**Repository Pattern**
- `CoffeeShopDataSource` handles all persistence via `Load()` and `Save()`
- Controllers never access the file system directly
- Models are passed between layers, not raw data

**Constructor Pattern**
```go
func NewCoffeeShopController(path carafepath.CarafePath) *CoffeeShopController
```
All major types use `New*` constructors for initialization.

**Identifier Pattern**
```go
type BlendIdentifier struct {
    Id    int
    Title string
}

func (bi BlendIdentifier) IsValid() bool { ... }
```
Identifiers can be resolved by either `Id` (≥0) or `Title` (non-empty). This provides flexibility for CLI usage.

**Validation**
- Validations are methods on types (e.g., `BlendIdentifier.IsValid()`)
- Controllers validate before performing operations
- Errors are communicated to the user via `fmt.Println` (not panic or os.Exit)

---

## Naming Conventions

### Constants
- UPPER_SNAKE_CASE for module-level constants
- Often grouped at the top of files
- Example: `HOUSE_BLEND_TITLE`, `HOUSE_BLEND_ID`, `margin_width`

### Types & Structs
- PascalCase: `CoffeeShopController`, `BlendIdentifier`, `Blend`, `Drip`
- Receiver names are short abbreviations:
  - `csc` for `*CoffeeShopController`
  - `bi` for `BlendIdentifier`
  - `ds` for `*DataSource`

### Variables & Functions
- camelCase: `foundBlend`, `dripId`, `dripIndex`, `dripText`
- Exported functions: PascalCase (`ListBlends`, `AddToBlends`)
- Unexported helpers: camelCase (`getBlendFromIdentifier`, `printBlend`)

### Packages
- Single-word, lowercase: `cmd`, `controller`, `model`, `datasource`
- Import with aliases for clarity:
  ```go
  datasource "github.com/tomasvalettini/latte/coffeeshop/data/data-source"
  carafepath "github.com/tomasvalettini/latte/coffeeshop/data/data-source/path"
  datamodel "github.com/tomasvalettini/latte/coffeeshop/data/model"
  ```

---

## Data Model

### Core Entities

**Blend**
```go
type Blend struct {
    Id    int    `json:"id"`
    Title string `json:"text"`      // Note: serialized as "text"
    Drips []Drip `json:"drips"`
}
```
- Represents a collection of drips (tasks)
- `Id` is unique, auto-generated
- Drips are ordered; each has a unique `Id` within the blend

**Drip**
```go
type Drip struct {
    Id   int    `json:"id"`
    Text string `json:"text"`
}
```
- Represents a single action/task
- `Id` is unique within its parent blend

### Serialization
- JSON tags define file format
- Files are stored as JSON arrays of blends
- Use `encoding/json` with `Marshal` and `Unmarshal`

---

## Testing

### Test File Conventions

- Test files: `*_test.go` in the same package
- Test functions: `TestFeatureName` or `TestMethod_Behavior`
- Example: `TestDeleteFromBlends_DeleteDripById`

### Common Test Setup

```go
func TestSomeFeature(t *testing.T) {
    // Setup
    tc := getTestCoffeeShopController()
    
    // Act
    tc.AddToBlends(&BlendIdentifier{Id: -1, Title: "Test"}, "text")
    
    // Assert & Verify
    blends := tc.dataSource.Load()
    // ... assertions
    
    // Cleanup
    t.Cleanup(func() {
        os.RemoveAll(carafepath.TMP)
    })
}
```

### Mocking stdin

For interactive prompts (e.g., `fmt.Scanln`), mock stdin using `os.Pipe()`:

```go
// Mock stdin for the confirmation prompt
oldStdin := os.Stdin
r, w, _ := os.Pipe()
os.Stdin = r
w.WriteString("yes\n")
w.Close()

// Call function that reads from stdin
tc.DeleteFromBlends(identifier, dripId)

// Restore stdin
os.Stdin = oldStdin
```

### Assertions

Use the `assert/` package for custom assertions:
```go
assert.Assert(condition, "error message")
```

For standard Go testing, use `testing.T`:
```go
if len(testBlend.Drips) != expectedCount {
    t.Errorf("Expected %d drips, got %d", expectedCount, len(testBlend.Drips))
}
```

### Test Utilities

The `test-utils/` package provides:
- `RequireExit(t *testing.T, testName string, testFunction func())` — Tests functions that call `os.Exit()`

---

## Adding a New Feature

### Adding a Command

1. **Create command file** in `cmd/yourcommand.go`:
   ```go
   package cmd
   
   import "github.com/spf13/cobra"
   
   var yourCmd = &cobra.Command{
       Use:   "yourcommand",
       Short: "Short description",
       Run: func(cmd *cobra.Command, args []string) {
           // Instantiate controller and call methods
       },
   }
   
   func init() {
       rootCmd.AddCommand(yourCmd)
   }
   ```

2. **Register in `cmd/root.go`** — The `init()` function automatically adds it via `rootCmd.AddCommand()`

3. **Add controller method** if business logic is needed (in `coffeeshop/controller/`)

4. **Write tests** alongside the command or controller

5. **Update README.md** with the new command

### Adding a Model Field

1. Update the struct in `coffeeshop/data/model/`
2. Add JSON tags for serialization
3. Update any constructors or helper functions
4. Write tests to verify serialization/deserialization
5. Update controller logic if needed

---

## Code Quality

### Formatting & Linting

```bash
# Format code
go fmt ./...

# Run linter
go vet ./...

# Run tests
go test ./...

# Run tests with coverage
go test -cover ./...
```

### Style Guidelines

- Keep functions focused and small
- Use meaningful variable names (prefer `foundBlend` over `b`)
- Use type-specific validation methods
- Prefer composition over inheritance (use embedding if needed)
- Avoid panics in business logic; use error returns or user-friendly messages
- Use constants for magic strings/numbers

### Error Handling

- Controllers communicate errors to users via `fmt.Println`
- Use `assert.Assert()` for fatal validation errors
- Log non-fatal issues and allow graceful degradation
- No panics in normal operation

---

## File Persistence

### Data Storage

- Blends are stored as JSON in a single file (path determined by `CarafePath`)
- File format: JSON array of blend objects
- Empty file system returns an empty slice (no error)
- `Load()` and `Save()` are the only I/O entry points

### CarafePath Abstraction

The `path/` layer provides multiple implementations:
- `LocalCarafePath` — Uses actual user home directory
- `TestCarafePath` — Uses temporary directories for testing
- `path.Interface` — Allows dependency injection

This enables tests to avoid interfering with real user data.

---

## Git Workflow

- Commit messages should be clear and descriptive
- Feature branches are recommended for larger changes
- Tests should pass before committing
- Keep commits atomic (one logical change per commit)

---

## Philosophy & Principles

1. **Unix Philosophy** — Do one thing well; compose small tools
2. **Simplicity** — Prefer straightforward code over clever abstractions
3. **Testability** — Design for easy testing (inject dependencies, avoid global state)
4. **User-Friendly** — Clear messages, sensible defaults, no surprises
5. **Separation of Concerns** — Keep layers distinct; avoid coupling

---

## Common Tasks

### Run all tests
```bash
go test ./...
```

### Run tests for a specific package
```bash
go test ./coffeeshop/controller
```

### Run a single test
```bash
go test -run TestDeleteFromBlends_DeleteDripById ./coffeeshop/controller
```

### Check coverage
```bash
go test -cover ./...
```

### Build for distribution
```bash
go build -o latte
```

### Install on Linux/Macos
```bash
go install
```

---

## Troubleshooting

**Test fails due to temp files not cleaning up**
- Ensure `t.Cleanup()` is called to remove `carafepath.TMP`

**stdin mocking not working**
- Confirm you're capturing `os.Stdin` before creating the pipe
- Make sure to close the write end and restore original stdin after

**Import path conflicts**
- Use import aliases to disambiguate (e.g., `datasource "github.com/..."`)

---

## Resources

- [Go Documentation](https://golang.org/doc)
- [Cobra CLI Framework](https://cobra.dev)
- [Clean Architecture in Go](https://github.com/golang-standards/project-layout)
