# Coding Standards

## Tech Stack & Versions

| Component        | Version     | Purpose                          |
|------------------|-------------|----------------------------------|
| Go               | 1.23.3      | Core language                    |
| cobra            | v1.10.2     | CLI command framework            |
| pflag            | v1.0.9      | Flag parsing (cobra dependency)  |
| mousetrap        |  v1.1.0     |  Windows console (transitive)    |
| Standard library | (stdlib)    | `encoding/json`, `os`, `log`, `fmt`, `strconv`, `strings`, `path/filepath` |

Module path: `github.com/<owner>/<project-name>`

No external testing framework. Uses Go stdlib `testing` package + custom `assert` package.

## Formatting & Linting

- **Format**: `go fmt ./...` (tabs for indentation, standard gofmt layout)
- **Vet**: `go vet ./...` before commit
- **Build**: `go build -o PROJECT_NAME`

## Naming Conventions

### Packages
- Single-word, lowercase: `cmd`, `controller`, `datasource`, `datamodel`, `projectpath`, `assert`, `testutils`
- Import with explicit aliases for nested paths:
  ```go
  datasource "github.com/<owner>/<project-name>/core/data/data-source"
  projectpath "github.com/<owner>/<project-name>/core/data/data-source/path"
  datamodel "github.com/<owner>/<project-name>/core/data/model"
  testutils "github.com/<owner>/<project-name>/test-utils"
  ```

### Types & Structs
- **PascalCase** for exported types: `MainController`, `Identifier`, `Collection`, `Entry`, `LocalProjectPath`, `TestProjectPath`
- **Short receiver names** (1-3 chars):
  - `mc` for `*MainController`
  - `id` for `Identifier`
  - `projectData` for `*ProjectDataSource`
  - `lpp` for `*LocalProjectPath`
  - `tpp` for `*TestProjectPath`

### Constants
- **UPPER_SNAKE_CASE** for module-level constants: `DEFAULT_COLLECTION_TITLE`, `DEFAULT_COLLECTION_ID`, `PROJECT_HOME_DIRECTORY`, `DATA_FILE_NAME`, `TMP`, `TEST_DATA_FILE`, `BE_CRASHER`
- File-scope constants use same convention: `margin_width = 6`

### Variables & Functions
- **PascalCase** for exported functions: `ListCollections`, `AddToCollections`, `GetNextId`, `NewMainController`
- **camelCase** for unexported: `getCollectionFromIdentifier`, `printCollections`, `getOrCreateCollectionFromIdentifier`, `addCollectionToCollectionList`, `performTestChecks`, `findCollectionByTitle`
- Struct field names PascalCase with snake-case JSON tags:
    ```go
  type Collection struct {
      Id    int    `json:"id"`
      Title string `json:"text"`   // Note: "text" not "title"
      Entries []Entry `json:"entries"`
  }
  ```
- Test helper functions camelCase at bottom of test file: `getTestMainController`, `getTestCollections`, `performTest`

### Test Types & Cases
- Test case struct prefix with `Test`: `TestIdentifier{ id, expected }`
- Test functions prefixed with `Test`: `TestIsValid`, `TestAddToCollections_WithNilIdentifier`
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
  Identifier{Id: -1, Title: "Example"}
  datamodel.Entry{Id: 0, Text: "test entry 1"}
  ```
- Constructors return pointer to struct: `&MainController{...}`

### Validation Methods
- Type-specific validation methods on the type:
    ```go
  func (id Identifier) IsValid() bool
  func (id Identifier) IsIdValid() bool
  func (id Identifier) IsTitleValid() bool
  func (id Identifier) Validate() *Identifier
  ```
- Methods use value receivers for immutable operations

### Slices & Returns
- Return empty slices, not nil: `return []datamodel.Collection{}`
- Use `append` for additions; iterate with `for i, x := range slice` when index needed
- Use `for _, x := range slice` when only values needed

### Constants at Top of File
Group module constants at file top with comments:
```go
// default collection that is PROJECT_NAME flavoured
const DEFAULT_COLLECTION_TITLE = "Default Collection"
const DEFAULT_COLLECTION_ID = 0

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
- Controller exports `ListCollections`, `AddToCollections`, `DeleteFromCollections`, `UpdateEntryInCollection`
- Model exports `Collection`, `Entry`, `GetNextCollectionId`, `GetNextId`, `MaxIdWidth`, `FindIndexFromId`
- Unexported helpers stay in same package: `getCollectionFromIdentifier`, `printCollections`

## Common Pitfalls

- **JSON tag mismatch**: `Collection.Title` serializes as `json:"text"` not `json:"title"`. Do not change without updating all consumers.
- **Default collection special case**: `Id == 0` is the default collection. `GetNextCollectionId` starts at 0 and returns `max+1`.
- **Entry ID starts at 0**: `GetNextId` starts at -1 and returns `max+1`. First entry has `Id == 0`.
- **Invalid identifier**: `Id < 0 AND Title == ""` → invalid. Use `Validate()` to get nil or pointer.
- **Import aliases**: Paths with hyphens (`data-source`) or reserved words (`model`) require aliases.
