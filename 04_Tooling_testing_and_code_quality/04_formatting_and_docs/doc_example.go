package docexample

// Package docexample demonstrates how to write package and symbol comments
// that are picked up by godoc and other documentation tools.
//
// The short package description should start right after the package clause.

// PublicFunc returns a greeting message for the given name.
func PublicFunc(name string) string {
	// This comment will appear in godoc as well.
	return "hello, " + name
}
