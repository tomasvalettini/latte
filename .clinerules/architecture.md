# Project Architecture

## Overview
`latte` is a terminal-based drip tracker built with Go, following a layered architecture with clear separation of concerns. CLI (Cobra) delegates to controllers for business logic, which use data sources for persistence and models for domain entities.

## Directory Structure

```
latte/
├── cmd/                          # CLI command handlers (Cobra)
│   ├── root.go                   # Root command, global flags, help
│   ├── add.go                    # Add drip command
│   ├── delete.go                 # Delete drip/blend command
│   ├── list.go                   # List blends/drips command
│   ├── update.go                 # Update drip command
│   └── constant.go               # Shared CLI constants
├── coffeeshop/                   # Core domain
│   ├── controller/              # Business logic
│   │   ├── coffeeshop_controller.go  # Main controller
│   │   ├── blend_identifier.go       # BlendIdentifier type
│   │   └── *_test.go                  # Controller tests
│   └── data/
│       ├── model/                # Domain entities
│       │   ├── blend.go         # Blend { Id, Title, Drips[]Drip }
│       │   ├── drip.go          # Drip { Id, Text }
│       │   ├── blend_extension.go  # GetNextBlendId
│       │   ├── drip_extension.go   # GetNextId, MaxIdWidth
│       │   └── *_test.go
│       └── data-source/         # Repository Pattern
│           ├── data_source.go       # DataSource interface
│           ├── coffeeshop_data_source.go  # JSON file I/O
│           └── path/
│               ├── carafe_path.go     # CarafePath interface
│               ├── local_carafe_path.go  # Production path
│               ├── test_carafe_path.go   # Test path
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
│  Controller Layer (coffeeshop/controller/) │
│  Business logic, validation, orchestration │
│  CoffeeShopController methods:           │
│    * ListBlends(bi *BlendIdentifier)     │
│    * AddToBlends(bi, dripText)          │
│    * DeleteFromBlends(bi, dripId)       │
│    * UpdateDripInBlend(bi, dripId, text) │
└──────────────────┬─────────────────────┘
                   │ uses
                   ▼
┌────────────────────────────────────────┐
│  Data Layer                              │
│  ├── Model: Blend, Drip (pure structs)  │
│  └── DataSource: Load/Save interface     │
│       CoffeeShopDataSource (JSON file)   │
│       CarafePath (file path abstraction) │
└────────────────────────────────────────┘
```

## Key Design Patterns

### Repository Pattern
- `DataSource` interface: `Load() []Blend`, `Save([]Blend)`
- `CoffeeShopDataSource`: JSON file persistence
- Controllers depend on interface → testable via `TestCarafePath`

### Constructor Pattern
```go
func NewCoffeeShopController(path carafepath.CarafePath) *CoffeeShopController
func NewCoffeeShopDataSource(path string) *CoffeeShopDataSource
```

### Identifier Pattern
`BlendIdentifier` resolves blends by Id or Title:
```go
type BlendIdentifier struct { Id int; Title string }
func (bi BlendIdentifier) IsValid() bool  // Id >= 0 OR Title != ""
func (bi BlendIdentifier) Validate() *BlendIdentifier  // nil if invalid
```

### Dependency Injection for Paths
`CarafePath` interface allows implementation swapping:
- `LocalCarafePath` → `$HOME/.latte/carafes.json`
- `TestCarafePath` → `./tmp/latte/test.json`

## Data Flow: Add Drip

1. `add.go` parses `--blend-id`, `--blend`, drip text
2. Creates `BlendIdentifier` from flags
3. `NewCoffeeShopController(carafePath)` → controller
4. `controller.AddToBlends(bi, dripText)`
5. Controller: validate text, load blends, find/create blend, append drip, save