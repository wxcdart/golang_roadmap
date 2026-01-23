// Demonstration: variable scope, package/imports, main entrypoint, and naming
// Every Go file starts with a package declaration.
package main

// import brings in code from other packages.
import "fmt"

// Package-level (global) variables must use `var`.
// Naming conventions: uppercase = exported, lowercase = private.
var ExportedVar = "I'm exported (uppercase)"
var unexportedVar = "I'm unexported (lowercase)"
var GlobalCount = 42

// main package and main() are the program entry point.
func main() {
    // Variables can be declared using `var` inside functions...
    var localVar string = "local var (declared with var)"

    // ...or using `:=` (short declaration) — only allowed inside functions.
    shortVar := "short var (declared with :=)"

    fmt.Println("-- Basic file structure --")
    fmt.Println("Package: package main")
    fmt.Println("Imports: fmt is used to print")
    fmt.Println("Entry point: main()")

    fmt.Println("\n-- Variables and scope --")
    fmt.Println("Exported global:", ExportedVar)
    fmt.Println("Unexported global:", unexportedVar)
    fmt.Println("GlobalCount:", GlobalCount)
    fmt.Println("localVar:", localVar)
    fmt.Println("shortVar:", shortVar)

    // Block-level scope: `blockVar` only exists inside the if-block.
    if true {
        blockVar := "I live inside this block"
        fmt.Println("blockVar inside block:", blockVar)
    }
    // The following line would be a compile error if uncommented,
    // because blockVar is not visible here (it's out of scope):
    // fmt.Println(blockVar)

    // Shadowing: redeclare a local variable with the same name.
    GlobalCount := 100 // this shadows the package-level GlobalCount inside main
    fmt.Println("Shadowed GlobalCount in main:", GlobalCount)
    fmt.Println("Package-level GlobalCount accessed via helper:", packageLevelCount())

    // Naming conventions matter for exported vs private.
    fmt.Println("Exported function call:", ExportedFunction())
    fmt.Println("Unexported function call:", unexportedFunction())
}

// packageLevelCount returns the package-level GlobalCount value.
func packageLevelCount() int { return GlobalCount }

// ExportedFunction (uppercase) is visible to other packages.
func ExportedFunction() string { return "I'm an exported function" }

// unexportedFunction (lowercase) is private to this package.
func unexportedFunction() string { return "I'm an unexported function" }
