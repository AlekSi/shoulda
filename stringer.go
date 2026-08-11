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

// dumpf constructs a [fmt.Stringer] for value followed by an optional suffix.
//
// valueFormat is formatted with these operands:
//
//	%[1] = value
//	%[2] = Dump(tb, value)
//
// valueFormat must use explicit operand indexes.
//
// format is formatted separately with args, so its operand numbering starts at %[1] again.
// Pass an empty format when no suffix is needed.
//
// The two results are concatenated without a separator, then trailing newlines are removed.
func dumpf(tb TB, valueFormat string, value any, format string, args ...any) fmt.Stringer {
	tb.Helper()

	return stringer(func() string {
		tb.Helper()

		v := fmt.Sprintf(valueFormat, value, Dump(tb, value))
		a := fmt.Sprintf(format, args...)

		return strings.TrimRight(v+a, "\n")
	})
}

// msgDiff constructs a [fmt.Stringer] comparing actual with expected.
//
// format is formatted with these operands:
//
//	%[1] = actual
//	%[2] = Dump(tb, actual)
//	%[3] = expected
//	%[4] = Dump(tb, expected)
//	%[5] = Diff(tb, "actual", actual, "expected", expected)
//
// format must use explicit operand indexes.
//
// Trailing newlines are removed from the formatted message.
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
