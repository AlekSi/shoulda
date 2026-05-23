// Package shoulda provides an assertion library that is as simple as possible, but not simpler.
//
// # Comparison
//
// Helpers ending with `f` accept [fmt.Printf]-like format string and arguments.
// They never [Dump] or [Diff] values.
// Those functions should be used by the caller if desired.
package shoulda
