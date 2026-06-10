package shoulda

import (
	"reflect"

	"github.com/AlekSi/shoulda/cmp"
)

// BeFalse checks that actual is false.
func BeFalse(tb TB, actual bool) bool {
	tb.Helper()

	m := msgf("actual is not false")

	return assert(tb, !actual, m)
}

// BeTrue checks that actual is true.
func BeTrue(tb TB, actual bool) bool {
	tb.Helper()

	m := msgf("actual is not true")

	return assert(tb, actual, m)
}

// BeDeepEqual checks that actual and expected are equal according to [reflect.DeepEqual].
func BeDeepEqual(tb TB, actual, expected any) bool {
	tb.Helper()

	args := []any{
		Dump(tb, actual),
		Dump(tb, expected),
		Diff(tb, "actual", actual, "expected", expected),
	}

	m := msgf("actual is not deep equal to expected:\nactual: %s\nexpected: %s\n%s", args...)

	return assert(tb, reflect.DeepEqual(actual, expected), m)
}

// NotBeDeepEqual checks that actual and expected are not equal according to [reflect.DeepEqual].
func NotBeDeepEqual(tb TB, actual, expected any) bool {
	tb.Helper()

	args := []any{
		Dump(tb, actual),
		Dump(tb, expected),
		Diff(tb, "actual", actual, "expected", expected),
	}

	m := msgf("actual is deep equal to expected:\nactual: %s\nexpected: %s\n%s", args...)

	return assert(tb, !reflect.DeepEqual(actual, expected), m)
}

// BeEqual checks that actual and expected are equal according to [cmp.Equal].
func BeEqual[T cmp.Ordered](tb TB, actual, expected T) bool {
	tb.Helper()

	args := []any{
		Dump(tb, actual),
		Dump(tb, expected),
		Diff(tb, "actual", actual, "expected", expected),
	}

	m := msgf("actual is not equal to expected:\nactual: %s\nexpected: %s\n%s", args...)

	return assert(tb, cmp.Equal(actual, expected), m)
}

// NotBeEqual checks that actual and expected are not equal according to [cmp.Equal].
func NotBeEqual[T cmp.Ordered](tb TB, actual, expected T) bool {
	tb.Helper()

	args := []any{
		Dump(tb, actual),
		Dump(tb, expected),
		Diff(tb, "actual", actual, "expected", expected),
	}

	m := msgf("actual is equal to expected:\nactual: %s\nexpected: %s\n%s", args...)

	return assert(tb, !cmp.Equal(actual, expected), m)
}

// BeLess checks that actual is less than expected according to [cmp.Less].
func BeLess[T cmp.Ordered](tb TB, actual, expected T) bool {
	tb.Helper()

	args := []any{
		Dump(tb, actual),
		Dump(tb, expected),
		Diff(tb, "actual", actual, "expected", expected),
	}

	m := msgf("actual is not less than expected:\nactual: %s\nexpected: %s\n%s", args...)

	return assert(tb, cmp.Less(actual, expected), m)
}

// BeGreater checks that actual is greater than expected according to [cmp.Greater].
func BeGreater[T cmp.Ordered](tb TB, actual, expected T) bool {
	tb.Helper()

	args := []any{
		Dump(tb, actual),
		Dump(tb, expected),
		Diff(tb, "actual", actual, "expected", expected),
	}

	m := msgf("actual is not greater than expected:\nactual: %s\nexpected: %s\n%s", args...)

	return assert(tb, cmp.Greater(actual, expected), m)
}
