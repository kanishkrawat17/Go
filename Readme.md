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