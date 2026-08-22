package shoulda

import (
	"reflect"

	"github.com/AlekSi/shoulda/cmp"
)

// BeFalse checks that actual is false.
func BeFalse(tb TB, actual bool) bool {
	tb.Helper()

	return BeFalsef(tb, actual, "")
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

	return BeTruef(tb, actual, "")
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

	return BeDeepEqualf(tb, actual, expected, "")
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

	return NotBeDeepEqualf(tb, actual, expected, "")
}

// NotBeDeepEqualf checks that actual and expected are not equal according to [reflect.DeepEqual].
func NotBeDeepEqualf(tb TB, actual, expected any, format string, args ...any) bool {
	tb.Helper()

	s := sprintf(
		"%s\n%s",
		msgDiff(tb, "actual is deep equal to expected:\nactual: %[2]s\nexpected: %[4]s\n%[5]s", actual, expected),
		sprintf(format, args...),
	)

	return assert(tb, !reflect.DeepEqual(actual, expected), s)
}

// BeEqual checks that actual and expected are equal according to [cmp.Equal].
func BeEqual[T cmp.Ordered](tb TB, actual, expected T) bool {
	tb.Helper()

	return BeEqualf(tb, actual, expected, "")
}

// BeEqualf checks that actual and expected are equal according to [cmp.Equal].
func BeEqualf[T cmp.Ordered](tb TB, actual, expected T, format string, args ...any) bool {
	tb.Helper()

	s := sprintf(
		"%s\n%s",
		msgDiff(tb, "actual is not equal to expected:\nactual: %[2]s\nexpected: %[4]s\n%[5]s", actual, expected),
		sprintf(format, args...),
	)

	return assert(tb, cmp.Equal(actual, expected), s)
}

// NotBeEqual checks that actual and expected are not equal according to [cmp.Equal].
func NotBeEqual[T cmp.Ordered](tb TB, actual, expected T) bool {
	tb.Helper()

	return NotBeEqualf(tb, actual, expected, "")
}

// NotBeEqualf checks that actual and expected are not equal according to [cmp.Equal].
func NotBeEqualf[T cmp.Ordered](tb TB, actual, expected T, format string, args ...any) bool {
	tb.Helper()

	s := sprintf(
		"%s\n%s",
		msgDiff(tb, "actual is equal to expected:\nactual: %[2]s\nexpected: %[4]s\n%[5]s", actual, expected),
		sprintf(format, args...),
	)

	return assert(tb, !cmp.Equal(actual, expected), s)
}

// BeLess checks that actual is less than expected according to [cmp.Less].
func BeLess[T cmp.Ordered](tb TB, actual, expected T) bool {
	tb.Helper()

	return BeLessf(tb, actual, expected, "")
}

// BeLessf checks that actual is less than expected according to [cmp.Less].
func BeLessf[T cmp.Ordered](tb TB, actual, expected T, format string, args ...any) bool {
	tb.Helper()

	s := sprintf(
		"%s\n%s",
		msgDiff(tb, "actual is not less than expected:\nactual: %[2]s\nexpected: %[4]s\n%[5]s", actual, expected),
		sprintf(format, args...),
	)

	return assert(tb, cmp.Less(actual, expected), s)
}

// BeGreater checks that actual is greater than expected according to [cmp.Greater].
func BeGreater[T cmp.Ordered](tb TB, actual, expected T) bool {
	tb.Helper()

	return BeGreaterf(tb, actual, expected, "")
}

// BeGreaterf checks that actual is greater than expected according to [cmp.Greater].
func BeGreaterf[T cmp.Ordered](tb TB, actual, expected T, format string, args ...any) bool {
	tb.Helper()

	s := sprintf(
		"%s\n%s",
		msgDiff(tb, "actual is not greater than expected:\nactual: %[2]s\nexpected: %[4]s\n%[5]s", actual, expected),
		sprintf(format, args...),
	)

	return assert(tb, cmp.Greater(actual, expected), s)
}
