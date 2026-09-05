# Testing & Validation

## Tooling

| Command                              | Purpose                                  |
|--------------------------------------|------------------------------------------|
| `go test ./...`                      | Run all tests across all packages        |
| `go test ./coffeeshop/controller`    | Run tests in a specific package          |
| `go test -run TestName ./...`        | Run a single test by name                |
| `go test -v ./...`                   | Verbose output (show each test)          |
| `go test -cover ./...`              | Coverage report                          |
| `go test -race ./...`                | Race condition detection                 |

## Custom Assertion Library (`assert/`)

```go
func Assert(truth bool, msg string) {
    if !truth {
        log.Fatalln(msg)
    }
}
```

**Note**: `log.Fatalln` is fatal — tests must isolate fatal paths using subprocess pattern.

## Fatal Test Pattern (`test-utils/RequireExit`)

For testing code paths that call `assert.Assert(false)` or `log.Fatalln`:

```go
func RequireExit(t *testing.T, testName string, testFunction func()) {
    if os.Getenv("BE_CRASHER") == "1" {
        testFunction()
        return
    }
    cmd := exec.Command(os.Args[0], "-test.run="+testName)
    cmd.Env = append(os.Environ(), "BE_CRASHER=1")
    err := cmd.Run()
    if e, ok := err.(*exec.ExitError); ok && !e.Success() {
        return // non-zero exit = success
    }
}
```

### When to use
- Testing `Validate()` with invalid input
- Testing `FindIndexFromId` with non-existent ID
- Testing `Load()` on a path pointing to a directory (causes read error)

## Test Structure

### Table-Driven Tests
Used for pure function testing (e.g. `BlendIdentifier` validation):

```go
type TestBlendIdentifier struct {
    bi       BlendIdentifier
    expected bool
}

func TestIsValid(t *testing.T) {
    testCases := []TestBlendIdentifier{
        {BlendIdentifier{Id: 0, Title: "Coffee"}, true},
        {BlendIdentifier{Id: -1, Title: ""}, false},
    }
    for _, tc := range testCases {
        name := fmt.Sprintf("Id=%d,Title=%s", tc.bi.Id, tc.bi.Title)
        t.Run(name, func(t *testing.T) {
            result := tc.bi.IsValid()
            msg := fmt.Sprintf("%s,result=%t,expected=%t", name, result, tc.expected)
            assert.Assert(result == tc.expected, msg)
        })
    }
}
```

### Integration Tests (Controller)
Tests controller methods end-to-end:

```go
func TestAddToBlends_WithNilIdentifier(t *testing.T) {
    tc := getTestCoffeeShopController()
    tc.AddToBlends(nil, "test drip 1")

    blends := tc.dataSource.Load()
    houseBlend := findBlendByTitle(blends, HOUSE_BLEND_TITLE)

    performTestChecks(map[string]bool{
        "House blend not found": houseBlend == nil,
        fmt.Sprintf("Expected 3 drips, got %d", len): dripsLen != 3,
    })

    t.Cleanup(func() {
        os.RemoveAll(carafepath.TMP)
    })
}
```

### Shared Test Helpers (bottom of test files)

```go
func getTestCoffeeShopController() *CoffeeShopController {
    tp := carafepath.GetTestingCarafePath()
    return NewCoffeeShopController(tp)
}

func performTestChecks(checks map[string]bool) {
    for check, passed := range checks {
        assert.Assert(!passed, check)
    }
}

func findBlendByTitle(blends []datamodel.Blend, title string) *datamodel.Blend {
    for i := range blends {
        if blends[i].Title == title {
            return &blends[i]
        }
    }
    return nil
}
```

## Test File Organization

- **Colocated**: `foo.go` and `foo_test.go` in same directory, same package
- **Package choice**:
  - White-box: `package controller` (can test unexported)
  - Black-box: `package datasource_test` (only exported)
## Test Naming Convention

- `TestXxx` — basic function/method test
- `TestXxx_WithYyy` — scenario-based test
- `TestXxx_WithYyyZzz` — more specific scenario

### Examples from codebase:
- `TestIsValid`, `TestIsIdValid`, `TestIsTitleValid`, `TestValidate`
- `TestAddToBlends_WithNilIdentifier`, `TestAddToBlends_WithEmptyDripText`
- `TestAddToBlends_CreatesNewBlend`, `TestAddToBlends_AddsToExistingBlendByTitle`
- `TestDeleteFromBlends_DeleteDripById`, `TestDeleteFromBlends_DeleteBlendWithConfirmation`
- `TestUpdateDripInBlend_UpdateFirstDrip`, `TestUpdateDripInBlend_UpdateLastDrip`, `TestUpdateDripInBlend_UpdateByBlendId`
- `TestCoffeeShopDataSourceLogic`, `TestCoffeeShopDataSourceFailingFile`

## Test Data Cleanup Pattern

Always use `t.Cleanup()` for filesystem cleanup:

```go
t.Cleanup(func() {
    os.RemoveAll(carafepath.TMP)  // removes tmp/ directory
})
```

- `TestCarafePath` → test path `tmp/latte/test.json`
- `GetTestingCarafePath()` returns the `CarafePath` interface
- Every controller test that writes data must call cleanup

## Validation Assertions

Standard pattern:

```go
assert.Assert(result == tc.expected,
    fmt.Sprintf("%s,result=%t,expected=%t", name, result, tc.expected))
```

### Pointer Validation (Validate returns *BlendIdentifier)

```go
result := tc.bi.Validate()
isNonNil := result != nil
msg := fmt.Sprintf("%s,result=%v,expected=%t", name, result, tc.expected)
assert.Assert(isNonNil == tc.expected, msg)

if result != nil {
    assert.Assert(result.Id == tc.bi.Id && result.Title == tc.bi.Title,
        fmt.Sprintf("%s,returned pointer has incorrect values", name))
}
```

## Coverage

- Target: >80% controller, >90% model/validation
- `coverage.out` is generated and checked into repo root

```bash
go test -cover ./...
go test -coverprofile=coverage.out ./...
go tool cover -func=coverage.out
go tool cover -html=coverage.out
```

## Pre-commit Checklist

1. `go fmt ./...` — no formatting diffs
2. `go vet ./...` — no warnings
3. `go test ./...` — all pass
4. `go test -race ./...` — no race conditions
5. `go test -cover ./...` — coverage acceptable
- **Helpers**: Table-driven case structs use `Test` prefix (`TestBlendIdentifier`)