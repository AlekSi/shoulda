package shoulda

// BeNil checks that actual is untyped nil.
//
// It is recommended to use [NoError] for errors and [BeZero] where possible otherwise.
func BeNil(tb TB, actual any) bool {
	tb.Helper()

	m := msgDumpf(tb, "actual is not untyped nil, but:\nactual: %[2]s", actual)

	return assert(tb, actual == nil, m)
}

// NotBeNil checks that actual is not (untyped) nil.
//
// It is recommended to use [Error] for errors and [NotBeZero] where possible otherwise.
func NotBeNil(tb TB, actual any) bool {
	tb.Helper()

	m := msg("actual is untyped nil")

	return assert(tb, actual != nil, m)
}

// BeZero checks that actual is the zero value of its type.
func BeZero[T comparable](tb TB, actual T) bool {
	tb.Helper()

	m := msgDumpf(tb, "actual is not zero, but:\nactual: %[2]s", actual)

	var zero T
	return assert(tb, actual == zero, m)
}

// NotBeZero checks that actual is not the zero value of its type.
func NotBeZero[T comparable](tb TB, actual T) bool {
	tb.Helper()

	m := msg("actual is zero")

	var zero T
	return assert(tb, actual != zero, m)
}

// Error checks that actual is a non-nil error.
func Error(tb TB, actual error) bool {
	tb.Helper()

	m := msg("actual is nil error")

	return assert(tb, actual != nil, m)
}

// NoError checks that actual is a nil error.
func NoError(tb TB, actual error) bool {
	tb.Helper()

	m := msgDumpf(tb, "actual is not nil error, but %[1]q:\nactual: %[2]s", actual)

	return assert(tb, actual == nil, m)
}
