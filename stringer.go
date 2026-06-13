package shoulda

import (
	"fmt"
	"strings"
)

// stringer implements [fmt.Stringer] for a function.
type stringer func() string

// String implements [fmt.Stringer].
func (s stringer) String() string { return s() }

// sprintf constructs a [fmt.Stringer] from a message or format string, and arguments.
func sprintf(format string, args ...any) fmt.Stringer {
	if len(args) == 0 {
		return stringer(func() string {
			return strings.TrimRight(format, "\n")
		})
	}

	return stringer(func() string {
		return strings.TrimRight(fmt.Sprintf(format, args...), "\n")
	})
}

// dumpf constructs a [fmt.Stringer] from a value, format string, and arguments.
// The value itself and its [Dump] result are prepended to args.
func dumpf(tb TB, value any, format string, args ...any) fmt.Stringer {
	tb.Helper()

	return stringer(func() string {
		tb.Helper()

		a := append([]any{value, Dump(tb, value)}, args...)

		return strings.TrimRight(fmt.Sprintf(format, a...), "\n")
	})
}

// msgDiff constructs a [fmt.Stringer] from a format string and values plus their diff.
func msgDiff(tb TB, format string, actual any, expected any) stringer {
	tb.Helper()

	return stringer(func() string {
		tb.Helper()

		args := []any{
			actual, Dump(tb, actual),
			expected, Dump(tb, expected),
			Diff(tb, "actual", actual, "expected", expected),
		}

		return strings.TrimRight(fmt.Sprintf(format, args...), "\n")
	})
}
