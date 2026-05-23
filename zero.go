package shoulda

// BeNilf checks that actual is untyped nil.
func BeNilf(tb TB, actual any, msg string, args ...any) bool {
	tb.Helper()

	m := messagef(msg, args...)

	return assert(tb, actual == nil, m)
}

// BeNil checks that actual is untyped nil.
func BeNil(tb TB, actual any) bool {
	tb.Helper()

	m := dumpf(tb, "actual is not nil, but %T:\n%s", actual)

	return assert(tb, actual == nil, m)
}

// NotBeNilf checks that actual is not (untyped) nil.
func NotBeNilf(tb TB, actual any, msg string, args ...any) bool {
	tb.Helper()

	m := messagef(msg, args...)

	return assert(tb, actual != nil, m)
}

// NotBeNil checks that actual is not (untyped) nil.
func NotBeNil(tb TB, actual any) bool {
	tb.Helper()

	m := messagef("actual is nil")

	return assert(tb, actual != nil, m)
}

// BeZerof checks that actual is the zero value of its type.
func BeZerof[T comparable](tb TB, actual T, msg string, args ...any) bool {
	tb.Helper()

	m := messagef(msg, args...)

	var zero T
	return assert(tb, actual == zero, m)
}

// BeZero checks that actual is the zero value of its type.
func BeZero[T comparable](tb TB, actual T) bool {
	tb.Helper()

	m := dumpf(tb, "actual is not zero, but %T:\n%s", actual)

	var zero T
	return assert(tb, actual == zero, m)
}

// NotBeZerof checks that actual is not the zero value of its type.
func NotBeZerof[T comparable](tb TB, actual T, msg string, args ...any) bool {
	tb.Helper()

	m := messagef(msg, args...)

	var zero T
	return assert(tb, actual != zero, m)
}

// NotBeZero checks that actual is not the zero value of its type.
func NotBeZero[T comparable](tb TB, actual T) bool {
	tb.Helper()

	m := messagef("actual is zero")

	var zero T
	return assert(tb, actual != zero, m)
}
