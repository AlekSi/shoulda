package shoulda

import (
	"time"

	"github.com/AlekSi/shoulda/cmp"
)

func ExampleSatisfyf_inline() {
	Satisfyf(t, 13, func(v int) bool { return v > 42 }, "extra message: %s, %d", "foo", 42)

	// Output:
	// actual is not satisfied by predicate:
	// actual: 13 (int)
	// extra message: foo, 42
	// FAIL
}

func ExampleSatisfyf_methodValue() {
	actual := time.Date(2026, time.April, 9, 17, 32, 42, 123, time.UTC)
	Satisfyf(t, actual, time.Now().Before, "extra message: %s, %d", "foo", 42)

	// Output:
	// actual is not satisfied by predicate:
	// actual: time.Date(2026, 4, 9, 17, 32, 42, 123, time.UTC) (time.Time)
	// extra message: foo, 42
	// FAIL
}

func ExampleSatisfyf_methodExpression() {
	actual := time.Date(2026, time.April, 9, 17, 32, 42, 123, time.UTC)
	Satisfyf(t, actual, time.Time.IsZero, "extra message: %s, %d", "foo", 42)

	// Output:
	// actual is not satisfied by predicate:
	// actual: time.Date(2026, 4, 9, 17, 32, 42, 123, time.UTC) (time.Time)
	// extra message: foo, 42
	// FAIL
}

func ExampleSatisfyWithf_inline() {
	SatisfyWithf(t, 13, 42, func(x, y int) bool { return x > y }, "extra message: %s, %d", "foo", 42)

	// Output:
	// actual and expected are not satisfied by predicate:
	// actual: 13 (int)
	// expected: 42 (int)
	// diff expected actual
	// --- expected
	// +++ actual
	// @@ -1,1 +1,1 @@
	// -42 (int)
	// +13 (int)
	// extra message: foo, 42
	// FAIL
}

func ExampleSatisfyWithf_function() {
	SatisfyWithf(t, 13, 42, cmp.Greater, "extra message: %s, %d", "foo", 42)

	// Output:
	// actual and expected are not satisfied by predicate:
	// actual: 13 (int)
	// expected: 42 (int)
	// diff expected actual
	// --- expected
	// +++ actual
	// @@ -1,1 +1,1 @@
	// -42 (int)
	// +13 (int)
	// extra message: foo, 42
	// FAIL
}

func ExampleSatisfyWithf_methodExpression() {
	actual := time.Date(2026, time.April, 9, 17, 32, 42, 123, time.UTC)
	expected := time.Date(2026, time.April, 9, 17, 32, 42, 123, time.FixedZone("My", 4*int(time.Hour.Seconds())))
	SatisfyWithf(t, actual, expected, time.Time.Before, "extra message: %s, %d", "foo", 42)

	// Output:
	// actual and expected are not satisfied by predicate:
	// actual: time.Date(2026, 4, 9, 17, 32, 42, 123, time.UTC) (time.Time)
	// expected: time.Date(2026, 4, 9, 13, 32, 42, 123, time.UTC) (time.Time)
	// diff expected actual
	// --- expected
	// +++ actual
	// @@ -1,1 +1,1 @@
	// -time.Date(2026, 4, 9, 13, 32, 42, 123, time.UTC) (time.Time)
	// +time.Date(2026, 4, 9, 17, 32, 42, 123, time.UTC) (time.Time)
	// extra message: foo, 42
	// FAIL
}

func ExampleCompareEqualf_methodExpression() {
	actual := time.Date(2026, time.April, 9, 17, 32, 42, 123, time.UTC)
	expected := time.Date(2026, time.April, 9, 17, 32, 42, 123, time.FixedZone("My", 4*int(time.Hour.Seconds())))
	CompareEqualf(t, actual, expected, time.Time.Compare, "extra message: %s, %d", "foo", 42)

	// Output:
	// actual is not equal to expected, but greater:
	// actual: time.Date(2026, 4, 9, 17, 32, 42, 123, time.UTC) (time.Time)
	// expected: time.Date(2026, 4, 9, 13, 32, 42, 123, time.UTC) (time.Time)
	// diff expected actual
	// --- expected
	// +++ actual
	// @@ -1,1 +1,1 @@
	// -time.Date(2026, 4, 9, 13, 32, 42, 123, time.UTC) (time.Time)
	// +time.Date(2026, 4, 9, 17, 32, 42, 123, time.UTC) (time.Time)
	// extra message: foo, 42
	// FAIL
}

func ExampleNotPanicf() {
	NotPanicf(t, func() { panic("boom") }, "extra message: %s, %d", "foo", 42)

	// Output:
	// function panicked:
	// actual: "boom" (string)
	// extra message: foo, 42
	// FAIL
}

func ExamplePanicSatisfyf_methodExpression() {
	PanicSatisfyf(t, time.Time.IsZero, func() {
		panic(time.Date(2026, time.April, 9, 17, 32, 42, 123, time.UTC))
	}, "extra message: %s, %d", "foo", 42)

	// Output:
	// actual is not satisfied by predicate:
	// actual: time.Date(2026, 4, 9, 17, 32, 42, 123, time.UTC) (time.Time)
	// extra message: foo, 42
	// FAIL
}

func ExamplePanicSatisfyf_assertion() {
	PanicSatisfyf(t, func(actual time.Time) bool {
		expected := time.Date(2026, time.April, 9, 17, 32, 42, 123, time.FixedZone("My", 4*int(time.Hour.Seconds())))
		CompareEqual(t, actual, expected, time.Time.Compare)
		return true
	}, func() {
		panic(time.Date(2026, time.April, 9, 17, 32, 42, 123, time.UTC))
	}, "extra message: %s, %d", "foo", 42)

	// Output:
	// actual is not equal to expected, but greater:
	// actual: time.Date(2026, 4, 9, 17, 32, 42, 123, time.UTC) (time.Time)
	// expected: time.Date(2026, 4, 9, 13, 32, 42, 123, time.UTC) (time.Time)
	// diff expected actual
	// --- expected
	// +++ actual
	// @@ -1,1 +1,1 @@
	// -time.Date(2026, 4, 9, 13, 32, 42, 123, time.UTC) (time.Time)
	// +time.Date(2026, 4, 9, 17, 32, 42, 123, time.UTC) (time.Time)
	// FAIL
}
