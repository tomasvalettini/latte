# Coding Standards

## Tech Stack & Versions

| Component        | Version     | Purpose                          |
|------------------|-------------|----------------------------------|
| Go               | 1.23.3      | Core language                    |
| cobra            | v1.10.2     | CLI command framework            |
| pflag            | v1.0.9      | Flag parsing (cobra dependency)  |
| mousetrap         | v1.1.0      | Windows console (transitive)    |
| Standard library | (stdlib)    | `encoding/json`, `os`, `log`, `fmt`, `strconv`, `strings`, `path/filepath` |

Module path: `github.com/tomasvalettini/latte`

No external testing framework. Uses Go stdlib `testing` package + custom `assert` package.

## Formatting & Linting

- **Format**: `go fmt ./...` (tabs for indentation, standard gofmt layout)
- **Vet**: `go vet ./...` before commit
- **Build**: `go build -o latte`

## Naming Conventions

### Packages
- Single-word, lowercase: `cmd`, `controller`, `datasource`, `datamodel`, `carafepath`, `assert`, `testutils`
- Import with explicit aliases for nested paths:
  ```go
  datasource "github.com/tomasvalettini/latte/coffeeshop/data/data-source"
  carafepath "github.com/tomasvalettini/latte/coffeeshop/data/data-source/path"
  datamodel "github.com/tomasvalettini/latte/coffeeshop/data/model"
  testutils "github.com/tomasvalettini/latte/test-utils"
  ```

### Types & Structs
- **PascalCase** for exported types: `CoffeeShopController`, `BlendIdentifier`, `Blend`, `Drip`, `LocalCarafePath`, `TestCarafePath`
- **Short receiver names** (1-3 chars):
  - `csc` for `*CoffeeShopController`
  - `bi` for `BlendIdentifier`
  - `coffeeShop` for `*CoffeeShopDataSource`
  - `ltp` for `*LocalCarafePath`
  - `ttp` for `*TestCarafePath`

### Constants
- **UPPER_SNAKE_CASE** for module-level constants: `HOUSE_BLEND_TITLE`, `HOUSE_BLEND_ID`, `LATTE_HOME_DIRECTORY`, `CARAFE_FILE_NAME`, `TMP`, `TEST_CARAFE_FILE`, `BE_CRASHER`
- File-scope constants use same convention: `margin_width = 6`

### Variables & Functions
- **PascalCase** for exported functions: `ListBlends`, `AddToBlends`, `GetNextId`, `NewCoffeeShopController`
- **camelCase** for unexported: `getBlendFromIdentifier`, `printBlends`, `getOrCreateBlendFromIdentifier`, `addBlendToBlendList`, `performTestChecks`, `findBlendByTitle`
- Struct field names PascalCase with snake-case JSON tags:
  ```go
  type Blend struct {
      Id    int    `json:"id"`
      Title string `json:"text"`   // Note: "text" not "title"
      Drips []Drip `json:"drips"`
  }
  ```
- Test helper functions camelCase at bottom of test file: `getTestCoffeeShopController`, `getTestBlends`, `performTest`

### Test Types & Cases
- Test case struct prefix with `Test`: `TestBlendIdentifier{ bi, expected }`
- Test functions prefixed with `Test`: `TestIsValid`, `TestAddToBlends_WithNilIdentifier`
- Subtests via `t.Run(name, func(t *testing.T){...})` with descriptive `fmt.Sprintf` names
- Naming pattern: `TestXxx` for basic, `TestXxx_YyyZzz` for scenarios

## Code Style

### Error Handling
- Controller errors → `fmt.Println` to user (not panic, not os.Exit)
- Fatal/programmer errors → `assert.Assert(truth, msg)` (calls `log.Fatalln`)
- Missing file in `Load()` → returns empty slice (graceful degradation)
- JSON parse errors → `assert.Assert(err == nil, ...)`

### Struct Initialization
- Prefer named fields in struct literals:
  ```go
  BlendIdentifier{Id: -1, Title: "Espresso"}
  datamodel.Drip{Id: 0, Text: "test drip 1"}
  ```
- Constructors return pointer to struct: `&CoffeeShopController{...}`

### Validation Methods
- Type-specific validation methods on the type:
  ```go
  func (bi BlendIdentifier) IsValid() bool
  func (bi BlendIdentifier) IsIdValid() bool
  func (bi BlendIdentifier) IsTitleValid() bool
  func (bi BlendIdentifier) Validate() *BlendIdentifier
  ```
- Methods use value receivers for immutable operations

### Slices & Returns
- Return empty slices, not nil: `return []datamodel.Blend{}`
- Use `append` for additions; iterate with `for i, x := range slice` when index needed
- Use `for _, x := range slice` when only values needed

### Constants at Top of File
Group module constants at file top with comments:
```go
// default blend (collection) that is `latte` flavoured
const HOUSE_BLEND_TITLE = "House Blend"
const HOUSE_BLEND_ID = 0

const margin_width = 6
```

### JSON Serialization
- `encoding/json` with `MarshalIndent` (2 spaces) for persistence
- `Unmarshal` for loading
- Trailing newline appended on save: `data = append(data, '\n')`
- File permissions: `0o644` for file, `0o755` for directories

## Package Conventions

### Layout per File
1. `package` declaration
2. `import` block (stdlib first, then third-party, then local)
3. Constants (if any)
4. Type definitions
5. Constructors (`New*`)
6. Methods on types
7. Unexported helper functions at bottom

### Exporting
- Only export what the CLI or other packages need
- Controller exports `ListBlends`, `AddToBlends`, `DeleteFromBlends`, `UpdateDripInBlend`
- Model exports `Blend`, `Drip`, `GetNextBlendId`, `GetNextId`, `MaxIdWidth`, `FindIndexFromId`
- Unexported helpers stay in same package: `getBlendFromIdentifier`, `printBlends`

## Common Pitfalls

- **JSON tag mismatch**: `Blend.Title` serializes as `json:"text"` not `json:"title"`. Do not change without updating all consumers.
- **House Blend special case**: `Id == 0` is the default/house blend. `GetNextBlendId` starts at 0 and returns `max+1`.
- **Drip ID starts at 0**: `GetNextId` starts at -1 and returns `max+1`. First drip has `Id == 0`.
- **Invalid identifier**: `Id < 0 AND Title == ""` → invalid. Use `Validate()` to get nil or pointer.
- **Import aliases**: Paths with hyphens (`data-source`) or reserved words (`model`) require aliases.