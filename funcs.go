package shoulda

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/AlekSi/shoulda/cmp"
)

// Satisfy checks that predicate returns true for actual.
func Satisfy[A any](tb TB, actual A, predicate func(_ A) bool) bool {
	tb.Helper()

	return Satisfyf(tb, actual, predicate, "")
}

// Satisfyf checks that predicate returns true for actual.
func Satisfyf[A any](tb TB, actual A, predicate func(_ A) bool, format string, args ...any) bool {
	tb.Helper()

	s := dumpf(tb, "actual is not satisfied by predicate:\nactual: %[2]s\n", actual, format, args...)

	return assert(tb, predicate(actual), s)
}

// SatisfyWith checks that predicate returns true for actual and expected.
func SatisfyWith[A, E any](tb TB, actual A, expected E, predicate func(_ A, _ E) bool) bool {
	tb.Helper()

	return SatisfyWithf(tb, actual, expected, predicate, "")
}

// SatisfyWithf checks that predicate returns true for actual and expected.
func SatisfyWithf[A, E any](tb TB, actual A, expected E, predicate func(_ A, _ E) bool, format string, args ...any) bool {
	tb.Helper()

	s := sprintf(
		"%s\n%s",
		msgDiff(
			tb,
			"actual and expected are not satisfied by predicate:\nactual: %[2]s\nexpected: %[4]s\n%[5]s",
			actual, expected,
		),
		sprintf(format, args...),
	)

	return assert(tb, predicate(actual, expected), s)
}

// CompareWith checks that compare(actual, expected) returns order.
func CompareWith[A, E any](tb TB, actual A, expected E, order cmp.Order, compare func(_ A, _ E) int) bool {
	tb.Helper()

	switch order {
	case cmp.OrderEqual:
		return CompareEqual(tb, actual, expected, compare)
	case cmp.OrderLess:
		return CompareLess(tb, actual, expected, compare)
	case cmp.OrderGreater:
		return CompareGreater(tb, actual, expected, compare)
	default:
		return assert(tb, false, sprintf("invalid cmp.%s", order))
	}
}

// CompareEqual checks that compare(actual, expected) returns 0 ([cmp.OrderEqual]).
func CompareEqual[A, E any](tb TB, actual A, expected E, compare func(_ A, _ E) int) bool {
	tb.Helper()

	res := compare(actual, expected)

	m := msgDiff(
		tb,
		"actual is not equal to expected, but "+cmp.Order(res).String()+":\nactual: %[2]s\nexpected: %[4]s\n%[5]s",
		actual, expected,
	)

	return assert(tb, res == 0, m)
}

// CompareLess checks that compare(actual, expected) returns -1 ([cmp.OrderLess]).
func CompareLess[A, E any](tb TB, actual A, expected E, compare func(_ A, _ E) int) bool {
	tb.Helper()

	res := compare(actual, expected)

	m := msgDiff(
		tb,
		"actual is not less than expected, but "+cmp.Order(res).String()+":\nactual: %[2]s\nexpected: %[4]s\n%[5]s",
		actual, expected,
	)

	return assert(tb, res == -1, m)
}

// CompareGreater checks that compare(actual, expected) returns 1 ([cmp.OrderGreater]).
func CompareGreater[A, E any](tb TB, actual A, expected E, compare func(_ A, _ E) int) bool {
	tb.Helper()

	res := compare(actual, expected)

	m := msgDiff(
		tb,
		"actual is not greater than expected, but "+cmp.Order(res).String()+":\nactual: %[2]s\nexpected: %[4]s\n%[5]s",
		actual, expected,
	)

	return assert(tb, res == +1, m)
}

// NotPanic checks that f does not panic.
func NotPanic(tb TB, f func()) (ok bool) {
	tb.Helper()

	return NotPanicf(tb, f, "")
}

// NotPanicf checks that f does not panic.
func NotPanicf(tb TB, f func(), format string, args ...any) (ok bool) {
	tb.Helper()

	defer func() {
		tb.Helper()

		r := recover()
		s := dumpf(tb, "function panicked:\nactual: %[2]s\n", r, format, args...)
		ok = assert(tb, r == nil, s)
	}()

	f()
	return
}

// PanicSatisfy checks that f panics with the value of type A.
// If predicate is not nil, it also checks that the panic value satisfies it.
func PanicSatisfy[A any](tb TB, predicate func(_ A) bool, f func()) (ok bool) {
	tb.Helper()

	return PanicSatisfyf(tb, predicate, f, "")
}

// PanicSatisfyf checks that f panics with the value of type A.
// If predicate is not nil, it also checks that the panic value satisfies it.
func PanicSatisfyf[A any](tb TB, predicate func(_ A) bool, f func(), format string, args ...any) (ok bool) {
	tb.Helper()

	defer func() {
		tb.Helper()

		r := recover()
		if r == nil {
			ok = assert(tb, false, sprintf("function did not panic\n"+format, args...))
			return
		}

		var actual A
		actual, ok = r.(A)
		s := stringer(func() string {
			tb.Helper()

			s := fmt.Sprintf(
				"actual panic value is not of type %s, but:\nactual: %s\n%s",
				reflect.TypeFor[A](), Dump(tb, r), fmt.Sprintf(format, args...),
			)
			return strings.TrimRight(s, "\n")
		})
		if !assert(tb, ok, s) {
			return
		}

		if predicate != nil {
			ok = Satisfyf(tb, actual, predicate, format, args...)
		}
	}()

	f()
	return
}
