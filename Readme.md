# Go Learning Notes

## Setup

- Install Go: `brew install go`
- Initialize a module: `go mod init <module-name>`
- Run all files: `go run .` (preferred over `go run main.go` when you have multiple files)

## Basics

### Packages
- Every Go file must declare a `package` at the top
- The `main` package is special — it defines an executable program
- All files in the same directory must belong to the same package

### Variables & Constants
- `const` for constants: `const apiName = "messages"`
- `var` for variables: `var count int` (zero-valued by default — `int` defaults to `0`)

### Functions
- Defined with `func`: `func GetName() string { return "John" }`
- **Exported** functions start with an uppercase letter (`GetName`) — accessible from other files in the same package and from external packages
- **Unexported** functions start with a lowercase letter (`getName`) — only accessible within the same file/package

### Running Code
- `go run main.go` — only compiles that single file
- `go run .` — compiles all `.go` files in the current directory (use this when you have multiple files)

## Packages & Imports

### One Directory = One Package
- All `.go` files in the same directory **must** declare the same `package` name
- e.g. `main.go`, `helpers.go`, `helpers2.go` in the root all use `package main`
- Mixing package names in the same folder (e.g. `package main` and `package x`) causes a compile error:
  `found packages main (a.go) and x (helpers2.go)`

### Same Package — No Import Needed
- Functions in the same package (same directory) are directly accessible without importing
- e.g. `GetName()` defined in `helpers.go` (package main) can be called directly in `main.go` (package main)

### Different Package — Must Import
- Functions in a different package (subdirectory) must be imported
- Import path = `<module-name>/<directory-name>` (e.g. `"GoLearning/utils"`)
- The **directory name** determines the import path, the **package name** inside the file determines how you call it:
  ```go
  // utils/helpers.go declares: package utils
  // Import it in main.go:
  import "GoLearning/utils"
  // Call it using the package name:
  utils.GetAge()
  ```

### Module Name
- Defined in `go.mod` via `go mod init <module-name>`
- Acts as the root prefix for all import paths in the project

## Structs

- Define a struct with `type`:
  ```go
  type User struct {
      Name   string
      Age    string
      Skills []string
  }
  ```
- **Creating a struct literal** — always need the type name before `{`:
  ```go
  user := User{Name: "Kanishk", Age: "25", Skills: []string{"React", "Go"}}
  ```
- You **cannot** write `user = { ... }` (unlike JS/JSON) — the type name is always required
- **Slices** use `[]string{"a", "b"}` not `["a", "b"]`
- **Trailing comma** is required when closing `}` is on a new line
- Two ways to declare and assign:
  ```go
  user := User{Name: "Kanishk"}           // short declaration + assignment
  var user User
  user = User{Name: "Kanishk"}            // separate declaration, then assignment
  ```

### Struct Tags (JSON)
- Add tags to control how fields are serialized:
  ```go
  type User struct {
      Name   string   `json:"name"`
      Age    string   `json:"age"`
      Skills []string `json:"skills"`
  }
  ```
- Without tags, JSON keys would be uppercase (`Name`, `Age`, `Skills`)

## JSON Marshal

- `json.Marshal(v)` converts a struct to JSON `[]byte`
- Import `"encoding/json"`
- Returns **two values**: `[]byte` and `error` (Go's error handling pattern)
- Must check `err != nil` — Go doesn't have try/catch
- Convert `[]byte` to string with `string(jsonData)`:
  ```go
  jsonData, err := json.Marshal(user)
  if err != nil {
      fmt.Println("Error:", err)
      return
  }
  fmt.Println(string(jsonData))
  // Output: {"name":"Kanishk","age":"25","skills":["React","Go"]}
  ```

## Pointers

### What is a Pointer?
- A pointer holds the **memory address** of a variable, not the value itself
- Think of it like a house address vs the house — the pointer is the address, the value is the house

### Two Key Operators

| Operator | Name        | What it does                        | Example          |
|----------|-------------|-------------------------------------|------------------|
| `&`      | Address-of  | Gets the memory address of a variable | `&y` → `0xABC`  |
| `*`      | Dereference | Gets the value at a memory address    | `*n` → `2`      |

### Declaring a Pointer
- `*int` means "a pointer to an int" (used in type declarations)
- `*n` means "the value that n points to" (used in expressions)
  ```go
  var y int = 2
  var p *int = &y    // p is a pointer to y, holds y's address
  fmt.Println(p)     // 0x1400000e0a8 (memory address)
  fmt.Println(*p)    // 2 (the value at that address)
  ```

### Why Pointers? — Pass by Value vs Pass by Reference

Go passes arguments **by value** (copies them). Without pointers, you can't modify the original variable.

```go
// WITHOUT pointer — works on a COPY, original unchanged
func Double(n int) int {
    n = n * 2
    return n
}

// WITH pointer — modifies the ORIGINAL value
func DoubleUsingPointer(n *int) {
    *n = *n * 2
}
```

```go
x := 2
y := 2

Double(x)                // x is still 2 (copy was doubled, not x)
DoubleUsingPointer(&y)   // y is now 4 (modified directly via address)
```

### Visual Breakdown

```
var y int = 2

  Variable y:
  ┌─────────┐
  │    2     │  ← value
  │ addr: A1 │  ← memory address
  └─────────┘

  &y  = A1              (& gives the address)

  DoubleUsingPointer(&y) — passes address A1

  Inside the function:
    n  = A1              (n holds the address)
    *n = 2               (dereference — read value at A1)
    *n = *n * 2          (write 4 to address A1)

  Back in main:
    y = 4                (y was modified directly)
```

### Common Mistakes
- `n = n * 2` inside a pointer function — this tries to multiply an **address** by 2 (won't compile)
- Must use `*n = *n * 2` — dereference first to get/set the actual value

## Project Structure Example

```
GoLearning/
├── go.mod              (module GoLearning)
├── main.go             (package main — entry point)
├── helpers.go          (package main — same package, no import needed)
├── helpers2.go         (package main — same package, no import needed)
└── utils/
    └── helpers.go      (package utils — different package, import as "GoLearning/utils")
```