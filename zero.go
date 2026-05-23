package shoulda

import (
	"github.com/AlekSi/shoulda/internal"
)

// BeNilf checks that actual is untyped nil.
func BeNilf(tb TB, actual any, msg string, args ...any) bool {
	tb.Helper()

	return assert(tb, actual == nil, internal.F(msg, args...))
}

// BeNil checks that actual is untyped nil.
func BeNil(tb TB, actual any) bool {
	tb.Helper()

	return BeNilf(tb, actual, "actual: %[1]v (%[1]T)\nis not nil", actual)
}

// NotBeNilf checks that actual is not (untyped) nil.
func NotBeNilf(tb TB, actual any, msg string, args ...any) bool {
	tb.Helper()

	return assert(tb, actual != nil, internal.F(msg, args...))
}

// NotBeNil checks that actual is not (untyped) nil.
func NotBeNil(tb TB, actual any) bool {
	tb.Helper()

	return NotBeNilf(tb, actual, "is nil")
}

// BeZerof checks that actual is the zero value of its type.
func BeZerof[T comparable](tb TB, actual T, msg string, args ...any) bool {
	tb.Helper()

	var zero T
	return assert(tb, actual == zero, internal.F(msg, args...))
}

// BeZero checks that actual is the zero value of its type.
func BeZero[T comparable](tb TB, actual T) bool {
	tb.Helper()

	return BeZerof(tb, actual, "actual: %v\nis not zero", Dump(tb, actual))
}

// NotBeZerof checks that actual is not the zero value of its type.
func NotBeZerof[T comparable](tb TB, actual T, msg string, args ...any) bool {
	tb.Helper()

	var zero T
	return assert(tb, actual != zero, internal.F(msg, args...))
}

// NotBeZero checks that actual is not the zero value of its type.
func NotBeZero[T comparable](tb TB, actual T) bool {
	tb.Helper()

	return NotBeZerof(tb, actual, "is zero")
}
