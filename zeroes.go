package shoulda

// BeZero checks that actual is the zero value of its type.
func BeZero[T comparable](tb TB, actual T) bool {
	tb.Helper()

	return BeZerof(tb, actual, "")
}

// BeZerof checks that actual is the zero value of its type.
func BeZerof[T comparable](tb TB, actual T, format string, args ...any) bool {
	tb.Helper()

	s := dumpf(tb, "actual is not zero, but:\nactual: %[2]s\n", actual, format, args...)

	var zero T
	return assert(tb, actual == zero, s)
}

// NotBeZero checks that actual is not the zero value of its type.
func NotBeZero[T comparable](tb TB, actual T) bool {
	tb.Helper()

	return NotBeZerof(tb, actual, "")
}

// NotBeZerof checks that actual is not the zero value of its type.
func NotBeZerof[T comparable](tb TB, actual T, format string, args ...any) bool {
	tb.Helper()

	s := sprintf("actual is zero\n"+format, args...)

	var zero T
	return assert(tb, actual != zero, s)
}
