package shoulda

// BeNil checks that actual is untyped nil.
//
// It is recommended to use [NoError] for errors and [BeZero] where possible otherwise.
func BeNil(tb TB, actual any) bool {
	tb.Helper()

	args := []any{Dump(tb, actual)}
	m := msgf("actual is not untyped nil, but:\nactual: %s", args...)

	return assert(tb, actual == nil, m)
}

// BeNilf checks that actual is untyped nil.
//
// It is recommended to use [NoError] for errors and [BeZero] where possible otherwise.
func BeNilf(tb TB, actual any, format string, args ...any) bool {
	tb.Helper()

	args = append([]any{Dump(tb, actual)}, args...)
	m := msgf("actual is not untyped nil, but:\nactual: %s\n"+format, args...)

	return assert(tb, actual == nil, m)
}

// NotBeNil checks that actual is not (untyped) nil.
//
// It is recommended to use [Error] for errors and [NotBeZero] where possible otherwise.
func NotBeNil(tb TB, actual any) bool {
	tb.Helper()

	m := msgf("actual is untyped nil")

	return assert(tb, actual != nil, m)
}

// NotBeNilf checks that actual is not (untyped) nil.
//
// It is recommended to use [Error] for errors and [NotBeZero] where possible otherwise.
func NotBeNilf(tb TB, actual any, format string, args ...any) bool {
	tb.Helper()

	m := msgf("actual is untyped nil\n"+format, args...)

	return assert(tb, actual != nil, m)
}

// BeZero checks that actual is the zero value of its type.
func BeZero[T comparable](tb TB, actual T) bool {
	tb.Helper()

	args := []any{Dump(tb, actual)}
	m := msgf("actual is not zero, but:\nactual: %s", args...)

	var zero T
	return assert(tb, actual == zero, m)
}

// BeZerof checks that actual is the zero value of its type.
func BeZerof[T comparable](tb TB, actual T, format string, args ...any) bool {
	tb.Helper()

	args = append([]any{Dump(tb, actual)}, args...)
	m := msgf("actual is not zero, but:\nactual: %s\n"+format, args...)

	var zero T
	return assert(tb, actual == zero, m)
}

// NotBeZero checks that actual is not the zero value of its type.
func NotBeZero[T comparable](tb TB, actual T) bool {
	tb.Helper()

	m := msgf("actual is zero")

	var zero T
	return assert(tb, actual != zero, m)
}

// NotBeZerof checks that actual is not the zero value of its type.
func NotBeZerof[T comparable](tb TB, actual T, format string, args ...any) bool {
	tb.Helper()

	m := msgf("actual is zero\n"+format, args...)

	var zero T
	return assert(tb, actual != zero, m)
}

// Error checks that actual is a non-nil error.
func Error(tb TB, actual error) bool {
	tb.Helper()

	m := msgf("actual is nil error")

	return assert(tb, actual != nil, m)
}

// Errorf checks that actual is a non-nil error.
func Errorf(tb TB, actual error, format string, args ...any) bool {
	tb.Helper()

	m := msgf("actual is nil error\n"+format, args...)

	return assert(tb, actual != nil, m)
}

// NoError checks that actual is a nil error.
func NoError(tb TB, actual error) bool {
	tb.Helper()

	args := []any{actual, Dump(tb, actual)}
	m := msgf("actual is not nil error, but %q:\nactual: %s", args...)

	return assert(tb, actual == nil, m)
}

// NoErrorf checks that actual is a nil error.
func NoErrorf(tb TB, actual error, format string, args ...any) bool {
	tb.Helper()

	args = append([]any{actual, Dump(tb, actual)}, args...)
	m := msgf("actual is not nil error, but %q:\nactual: %s\n"+format, args...)

	return assert(tb, actual == nil, m)
}
