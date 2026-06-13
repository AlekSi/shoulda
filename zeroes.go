package shoulda

// BeNil checks that actual is untyped nil.
//
// It is recommended to use [NoError] for errors and [BeZero] where possible otherwise.
func BeNil(tb TB, actual any) bool {
	tb.Helper()

	s := dumpf(tb, actual, "actual is not untyped nil, but:\nactual: %[2]s")

	return assert(tb, actual == nil, s)
}

// BeNilf checks that actual is untyped nil.
//
// It is recommended to use [NoError] for errors and [BeZero] where possible otherwise.
func BeNilf(tb TB, actual any, format string, args ...any) bool {
	tb.Helper()

	s := dumpf(tb, actual, "actual is not untyped nil, but:\nactual: %[2]s\n"+format, args...)

	return assert(tb, actual == nil, s)
}

// NotBeNil checks that actual is not (untyped) nil.
//
// It is recommended to use [Error] for errors and [NotBeZero] where possible otherwise.
func NotBeNil(tb TB, actual any) bool {
	tb.Helper()

	s := sprintf("actual is untyped nil")

	return assert(tb, actual != nil, s)
}

// NotBeNilf checks that actual is not (untyped) nil.
//
// It is recommended to use [Error] for errors and [NotBeZero] where possible otherwise.
func NotBeNilf(tb TB, actual any, format string, args ...any) bool {
	tb.Helper()

	s := sprintf("actual is untyped nil\n"+format, args...)

	return assert(tb, actual != nil, s)
}

// BeZero checks that actual is the zero value of its type.
func BeZero[T comparable](tb TB, actual T) bool {
	tb.Helper()

	s := dumpf(tb, actual, "actual is not zero, but:\nactual: %[2]s")

	var zero T
	return assert(tb, actual == zero, s)
}

// BeZerof checks that actual is the zero value of its type.
func BeZerof[T comparable](tb TB, actual T, format string, args ...any) bool {
	tb.Helper()

	s := dumpf(tb, actual, "actual is not zero, but:\nactual: %[2]s\n"+format, args...)

	var zero T
	return assert(tb, actual == zero, s)
}

// NotBeZero checks that actual is not the zero value of its type.
func NotBeZero[T comparable](tb TB, actual T) bool {
	tb.Helper()

	s := sprintf("actual is zero")

	var zero T
	return assert(tb, actual != zero, s)
}

// NotBeZerof checks that actual is not the zero value of its type.
func NotBeZerof[T comparable](tb TB, actual T, format string, args ...any) bool {
	tb.Helper()

	s := sprintf("actual is zero\n"+format, args...)

	var zero T
	return assert(tb, actual != zero, s)
}

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

// NoError checks that actual is a nil error.
func NoError(tb TB, actual error) bool {
	tb.Helper()

	s := dumpf(tb, actual, "actual is not nil error, but %[1]q:\nactual: %[2]s")

	return assert(tb, actual == nil, s)
}

// NoErrorf checks that actual is a nil error.
func NoErrorf(tb TB, actual error, format string, args ...any) bool {
	tb.Helper()

	s := dumpf(tb, actual, "actual is not nil error, but %[1]q:\nactual: %[2]s\n"+format, args...)

	return assert(tb, actual == nil, s)
}
