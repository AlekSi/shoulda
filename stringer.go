package shoulda

import (
	"fmt"
	"strings"
)

// stringer implements [fmt.Stringer] for a function.
type stringer func() string

// String implements [fmt.Stringer].
func (s stringer) String() string { return s() }

// sprintf constructs a [fmt.Stringer] from a format string and arguments.
func sprintf(format string, args ...any) fmt.Stringer {
	return stringer(func() string {
		return strings.TrimRight(fmt.Sprintf(format, args...), "\n")
	})
}

// dumpf constructs a [fmt.Stringer] from format strings, value, and arguments.
// The value itself and its [Dump] result act as arguments for formatValue.
func dumpf(tb TB, formatValue string, value any, format string, args ...any) fmt.Stringer {
	tb.Helper()

	return stringer(func() string {
		tb.Helper()

		v := fmt.Sprintf(formatValue, value, Dump(tb, value))
		a := fmt.Sprintf(format, args...)

		return strings.TrimRight(v+a, "\n")
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
