# Project Architecture

## Overview
`PROJECT_NAME` is a terminal-based tracker built with Go, following a layered architecture with clear separation of concerns. CLI (Cobra) delegates to controllers for business logic, which use data sources for persistence and models for domain entities.

## Directory Structure

```
PROJECT_ROOT/
├── cmd/                          # CLI command handlers (Cobra)
│   ├── root.go                   # Root command, global flags, help
│   ├── add.go                    # Add entry command
│   ├── delete.go                 # Delete entry/collection command
│   ├── list.go                   # List collections/entries command
│   ├── update.go                 # Update entry command
│   └── constant.go               # Shared CLI constants
├── core/                         # Core domain
│   ├── controller/              # Business logic
│   │   ├── main_controller.go   # Main controller
│   │   ├── identifier.go        # Identifier type
│   │   └── *_test.go                  # Controller tests
│   └── data/
│       ├── model/                # Domain entities
│       │   ├── collection.go     # Collection { Id, Title, Entries[]Entry }
│       │   ├── entry.go          # Entry { Id, Text }
│       │   ├── collection_extension.go  # GetNextCollectionId
│       │   ├── entry_extension.go   # GetNextId, MaxIdWidth
│       │   └── *_test.go
│       └── data-source/         # Repository Pattern
│           ├── data_source.go       # DataSource interface
│           ├── project_data_source.go  # JSON file I/O
│           └── path/
│               ├── project_path.go    # ProjectPath interface
│               ├── local_project_path.go  # Production path
│               ├── test_project_path.go   # Test path
│               └── *_test.go
├── assert/                       # Custom assertions
│   ├── assert.go                 # Assert(truth, msg)
│   └── assert_test.go
├── test-utils/                   # Test utilities
│   ├── test_utils.go            # RequireExit
│   └── test_utils_test.go
├── main.go                       # Entry point
└── go.mod                        # Go 1.23.3, cobra v1.10.2
```

## Layered Architecture

```
┌────────────────────────────────────────┐
│  CLI Layer (cmd/)                       │
│  Cobra parsing, flags, fmt.Println/Scanln │
└──────────────────┬─────────────────────┘
                   │ delegates
                   ▼
┌────────────────────────────────────────┐
│  Controller Layer (core/controller/) │
│  Business logic, validation, orchestration │
│  MainController methods:           │
│    * ListCollections(id *Identifier)     │
│    * AddToCollections(id, entryText)          │
│    * DeleteFromCollections(id, entryId)       │
│    * UpdateEntryInCollection(id, entryId, text) │
└──────────────────┬─────────────────────┘
                   │ uses
                   ▼
┌────────────────────────────────────────┐
│  Data Layer                              │
│  ├── Model: Collection, Entry (pure structs)  │
│  └── DataSource: Load/Save interface     │
│       ProjectDataSource (JSON file)   │
│       ProjectPath (file path abstraction) │
└────────────────────────────────────────┘
```

## Key Design Patterns

### Repository Pattern
- `DataSource` interface: `Load() []Collection`, `Save([]Collection)`
- `ProjectDataSource`: JSON file persistence
- Controllers depend on interface → testable via `TestProjectPath`

### Constructor Pattern
```go
func NewMainController(path projectpath.ProjectPath) *MainController
func NewProjectDataSource(path string) *ProjectDataSource
```

### Identifier Pattern
`Identifier` resolves collections by Id or Title:
```go
type Identifier struct { Id int; Title string }
func (id Identifier) IsValid() bool  // Id >= 0 OR Title != ""
func (id Identifier) Validate() *Identifier  // nil if invalid
```

### Dependency Injection for Paths
`ProjectPath` interface allows implementation swapping:
- `LocalProjectPath` → `$HOME/.PROJECT_ROOT/data.json`
- `TestProjectPath` → `./tmp/PROJECT_ROOT/test.json`

## Data Flow: Add Entry

1. `add.go` parses `--collection-id`, `--collection`, entry text
2. Creates `Identifier` from flags
3. `NewMainController(projectPath)` → controller
4. `controller.AddToCollections(id, entryText)`
5. Controller: validate text, load collections, find/create collection, append entry, save