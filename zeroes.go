package shoulda

// BeNil checks that actual is untyped nil.
func BeNil(tb TB, actual any) bool {
	tb.Helper()

	m := dumpf(tb, "actual is not untyped nil, but %T:\n%s", actual)

	return assert(tb, actual == nil, m)
}

// NotBeNil checks that actual is not (untyped) nil.
func NotBeNil(tb TB, actual any) bool {
	tb.Helper()

	m := messagef("actual is untyped nil")

	return assert(tb, actual != nil, m)
}

// BeZero checks that actual is the zero value of its type.
func BeZero[T comparable](tb TB, actual T) bool {
	tb.Helper()

	m := dumpf(tb, "actual is not zero, but %T:\n%s", actual)

	var zero T
	return assert(tb, actual == zero, m)
}

// NotBeZero checks that actual is not the zero value of its type.
func NotBeZero[T comparable](tb TB, actual T) bool {
	tb.Helper()

	m := messagef("actual is zero")

	var zero T
	return assert(tb, actual != zero, m)
}
