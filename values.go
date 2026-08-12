package shoulda

import (
	"reflect"

	"github.com/AlekSi/shoulda/cmp"
)

// BeFalse checks that actual is false.
func BeFalse(tb TB, actual bool) bool {
	tb.Helper()

	s := sprintf("actual is not false")

	return assert(tb, !actual, s)
}

// BeFalsef checks that actual is false.
func BeFalsef(tb TB, actual bool, format string, args ...any) bool {
	tb.Helper()

	s := sprintf("actual is not false\n"+format, args...)

	return assert(tb, !actual, s)
}

// BeTrue checks that actual is true.
func BeTrue(tb TB, actual bool) bool {
	tb.Helper()

	s := sprintf("actual is not true")

	return assert(tb, actual, s)
}

// BeTruef checks that actual is true.
func BeTruef(tb TB, actual bool, format string, args ...any) bool {
	tb.Helper()

	s := sprintf("actual is not true\n"+format, args...)

	return assert(tb, actual, s)
}

// BeDeepEqual checks that actual and expected are equal according to [reflect.DeepEqual].
func BeDeepEqual(tb TB, actual, expected any) bool {
	tb.Helper()

	m := msgDiff(tb, "actual is not deep equal to expected:\nactual: %[2]s\nexpected: %[4]s\n%[5]s", actual, expected)

	return assert(tb, reflect.DeepEqual(actual, expected), m)
}

// BeDeepEqualf checks that actual and expected are equal according to [reflect.DeepEqual].
func BeDeepEqualf(tb TB, actual, expected any, format string, args ...any) bool {
	tb.Helper()

	s := sprintf(
		"%s\n%s",
		msgDiff(tb, "actual is not deep equal to expected:\nactual: %[2]s\nexpected: %[4]s\n%[5]s", actual, expected),
		sprintf(format, args...),
	)

	return assert(tb, reflect.DeepEqual(actual, expected), s)
}

// NotBeDeepEqual checks that actual and expected are not equal according to [reflect.DeepEqual].
func NotBeDeepEqual(tb TB, actual, expected any) bool {
	tb.Helper()

	m := msgDiff(tb, "actual is deep equal to expected:\nactual: %[2]s\nexpected: %[4]s\n%[5]s", actual, expected)

	return assert(tb, !reflect.DeepEqual(actual, expected), m)
}

// BeEqual checks that actual and expected are equal according to [cmp.Equal].
func BeEqual[T cmp.Ordered](tb TB, actual, expected T) bool {
	tb.Helper()

	m := msgDiff(tb, "actual is not equal to expected:\nactual: %[2]s\nexpected: %[4]s\n%[5]s", actual, expected)

	return assert(tb, cmp.Equal(actual, expected), m)
}

// NotBeEqual checks that actual and expected are not equal according to [cmp.Equal].
func NotBeEqual[T cmp.Ordered](tb TB, actual, expected T) bool {
	tb.Helper()

	m := msgDiff(tb, "actual is equal to expected:\nactual: %[2]s\nexpected: %[4]s\n%[5]s", actual, expected)

	return assert(tb, !cmp.Equal(actual, expected), m)
}

// BeLess checks that actual is less than expected according to [cmp.Less].
func BeLess[T cmp.Ordered](tb TB, actual, expected T) bool {
	tb.Helper()

	m := msgDiff(tb, "actual is not less than expected:\nactual: %[2]s\nexpected: %[4]s\n%[5]s", actual, expected)

	return assert(tb, cmp.Less(actual, expected), m)
}

// BeGreater checks that actual is greater than expected according to [cmp.Greater].
func BeGreater[T cmp.Ordered](tb TB, actual, expected T) bool {
	tb.Helper()

	m := msgDiff(tb, "actual is not greater than expected:\nactual: %[2]s\nexpected: %[4]s\n%[5]s", actual, expected)

	return assert(tb, cmp.Greater(actual, expected), m)
}
