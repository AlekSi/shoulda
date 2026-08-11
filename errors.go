package shoulda

import "errors"

// Error checks that actual is a non-nil error.
func Error(tb TB, actual error) bool {
	tb.Helper()

	s := sprintf("actual is nil error")

	return assert(tb, actual != nil, s)
}

// Errorf checks that actual is a non-nil error.
func Errorf(tb TB, actual error, format string, args ...any) bool {
	tb.Helper()

	s := sprintf("actual is nil error\n"+format, args...)

	return assert(tb, actual != nil, s)
}

// ErrorIs checks that actual matches expected using [errors.Is].
func ErrorIs(tb TB, actual, expected error) bool {
	tb.Helper()

	s := sprintf(
		"actual does not match expected:\nactual: %s\nexpected: %s",
		dumpf(tb, "%[2]s", actual, ""),
		dumpf(tb, "%[2]s", expected, ""),
	)

	return assert(tb, errors.Is(actual, expected), s)
}

// NoError checks that actual is a nil error.
func NoError(tb TB, actual error) bool {
	tb.Helper()

	s := dumpf(tb, "actual is not nil error, but %[1]q:\nactual: %[2]s", actual, "")

	return assert(tb, actual == nil, s)
}

// NoErrorf checks that actual is a nil error.
func NoErrorf(tb TB, actual error, format string, args ...any) bool {
	tb.Helper()

	s := dumpf(tb, "actual is not nil error, but %[1]q:\nactual: %[2]s\n", actual, format, args...)

	return assert(tb, actual == nil, s)
}
