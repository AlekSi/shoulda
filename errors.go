package shoulda

import "errors"

// Error checks that actual is a non-nil error.
func Error(tb TB, actual error) bool {
	tb.Helper()

	return Errorf(tb, actual, "")
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

	return ErrorIsf(tb, actual, expected, "")
}

// ErrorIsf checks that actual matches expected using [errors.Is].
func ErrorIsf(tb TB, actual, expected error, format string, args ...any) bool {
	tb.Helper()

	s := sprintf(
		"actual error does not match expected:\nactual: %s\nexpected: %s\n%s",
		dumpf(tb, "%[1]v\n%[2]s", actual, ""),
		dumpf(tb, "%[1]v\n%[2]s\n", expected, ""),
		sprintf(format, args...),
	)

	return assert(tb, errors.Is(actual, expected), s)
}

// NoError checks that actual is a nil error.
func NoError(tb TB, actual error) bool {
	tb.Helper()

	return NoErrorf(tb, actual, "")
}

// NoErrorf checks that actual is a nil error.
func NoErrorf(tb TB, actual error, format string, args ...any) bool {
	tb.Helper()

	s := dumpf(tb, "actual is not nil error:\nactual: %[1]v\n%[2]s\n", actual, format, args...)

	return assert(tb, actual == nil, s)
}
