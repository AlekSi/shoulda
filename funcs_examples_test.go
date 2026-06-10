package shoulda

import (
	"time"

	"github.com/AlekSi/shoulda/cmp"
)

func ExampleSatisfy_inline() {
	Satisfy(t, 13, func(v int) bool { return v > 42 })

	// Output:
	// actual is not satisfied by predicate:
	// actual: 13 (int)
	// FAIL
}

func ExampleSatisfy_methodValue() {
	actual := time.Date(2026, time.April, 9, 17, 32, 42, 123, time.UTC)
	Satisfy(t, actual, time.Now().Before)

	// Output:
	// actual is not satisfied by predicate:
	// actual: time.Date(2026, 4, 9, 17, 32, 42, 123, time.UTC) (time.Time)
	// FAIL
}

func ExampleSatisfy_methodExpression() {
	actual := time.Date(2026, time.April, 9, 17, 32, 42, 123, time.UTC)
	Satisfy(t, actual, time.Time.IsZero)

	// Output:
	// actual is not satisfied by predicate:
	// actual: time.Date(2026, 4, 9, 17, 32, 42, 123, time.UTC) (time.Time)
	// FAIL
}

func ExampleSatisfyf_inline() {
	Satisfyf(t, 13, func(v int) bool { return v > 42 }, "extra message: %s", "foo")

	// Output:
	// actual is not satisfied by predicate:
	// actual: 13 (int)
	// extra message: foo
	// FAIL
}

func ExampleSatisfyWith_inline() {
	SatisfyWith(t, 13, 42, func(x, y int) bool { return x > y })

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
	// FAIL
}

func ExampleSatisfyWith_function() {
	SatisfyWith(t, 13, 42, cmp.Greater)

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
	// FAIL
}

func ExampleSatisfyWith_methodExpression() {
	actual := time.Date(2026, time.April, 9, 17, 32, 42, 123, time.UTC)
	expected := time.Date(2026, time.April, 9, 17, 32, 42, 123, time.FixedZone("My", 4*int(time.Hour.Seconds())))
	SatisfyWith(t, actual, expected, time.Time.Before)

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
	// FAIL
}

func ExampleCompareEqual_methodExpression() {
	actual := time.Date(2026, time.April, 9, 17, 32, 42, 123, time.UTC)
	expected := time.Date(2026, time.April, 9, 17, 32, 42, 123, time.FixedZone("My", 4*int(time.Hour.Seconds())))
	CompareEqual(t, actual, expected, time.Time.Compare)

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
