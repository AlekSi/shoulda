package shoulda

import (
	"testing"
	"time"

	"github.com/AlekSi/shoulda/cmp"
)

func TestSatisfyf(t *testing.T) {
	t.Run("Inline", func(t *testing.T) {
		tt, lines := setup(t)
		Satisfyf(tt, 13, func(v int) bool { return v > 42 }, "extra message: %s, %d", "foo", 67)

		BeDeepEqual(t, lines(), []string{
			"actual is not satisfied by predicate:",
			"actual: 13 (int)",
			"extra message: foo, 67",
			"FAIL",
		})
	})

	t.Run("MethodValue", func(t *testing.T) {
		tt, lines := setup(t)
		actual := time.Date(2026, time.April, 9, 17, 32, 42, 123, time.UTC)
		Satisfyf(tt, actual, time.Now().Before, "extra message: %s, %d", "foo", 67)

		BeDeepEqual(t, lines(), []string{
			"actual is not satisfied by predicate:",
			"actual: time.Date(2026, 4, 9, 17, 32, 42, 123, time.UTC) (time.Time)",
			"extra message: foo, 67",
			"FAIL",
		})
	})

	t.Run("MethodExpression", func(t *testing.T) {
		tt, lines := setup(t)
		actual := time.Date(2026, time.April, 9, 17, 32, 42, 123, time.UTC)
		Satisfyf(tt, actual, time.Time.IsZero, "extra message: %s, %d", "foo", 67)

		BeDeepEqual(t, lines(), []string{
			"actual is not satisfied by predicate:",
			"actual: time.Date(2026, 4, 9, 17, 32, 42, 123, time.UTC) (time.Time)",
			"extra message: foo, 67",
			"FAIL",
		})
	})
}

func TestNotSatisfyf(t *testing.T) {
	t.Run("Inline", func(t *testing.T) {
		tt, lines := setup(t)
		NotSatisfyf(tt, 13, func(v int) bool { return v < 42 }, "extra message: %s, %d", "foo", 67)

		BeDeepEqual(t, lines(), []string{
			"actual is satisfied by predicate:",
			"actual: 13 (int)",
			"extra message: foo, 67",
			"FAIL",
		})
	})

	t.Run("MethodValue", func(t *testing.T) {
		tt, lines := setup(t)
		actual := time.Date(2026, time.April, 9, 17, 32, 42, 123, time.UTC)
		NotSatisfyf(tt, actual, time.Now().After, "extra message: %s, %d", "foo", 67)

		BeDeepEqual(t, lines(), []string{
			"actual is satisfied by predicate:",
			"actual: time.Date(2026, 4, 9, 17, 32, 42, 123, time.UTC) (time.Time)",
			"extra message: foo, 67",
			"FAIL",
		})
	})

	t.Run("MethodExpression", func(t *testing.T) {
		tt, lines := setup(t)
		actual := time.Time{}
		NotSatisfyf(tt, actual, time.Time.IsZero, "extra message: %s, %d", "foo", 67)

		BeDeepEqual(t, lines(), []string{
			"actual is satisfied by predicate:",
			"actual: time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC) (time.Time)",
			"extra message: foo, 67",
			"FAIL",
		})
	})
}

func TestSatisfyWithf(t *testing.T) {
	t.Run("Inline", func(t *testing.T) {
		tt, lines := setup(t)
		SatisfyWithf(tt, 13, 42, func(x, y int) bool { return x > y }, "extra message: %s, %d", "foo", 67)

		BeDeepEqual(t, lines(), []string{
			"actual and expected are not satisfied by predicate:",
			"actual: 13 (int)",
			"expected: 42 (int)",
			"diff expected actual",
			"--- expected",
			"+++ actual",
			"@@ -1,1 +1,1 @@",
			"-42 (int)",
			"+13 (int)",
			"extra message: foo, 67",
			"FAIL",
		})
	})

	t.Run("Function", func(t *testing.T) {
		tt, lines := setup(t)
		SatisfyWithf(tt, 13, 42, cmp.Greater, "extra message: %s, %d", "foo", 67)

		BeDeepEqual(t, lines(), []string{
			"actual and expected are not satisfied by predicate:",
			"actual: 13 (int)",
			"expected: 42 (int)",
			"diff expected actual",
			"--- expected",
			"+++ actual",
			"@@ -1,1 +1,1 @@",
			"-42 (int)",
			"+13 (int)",
			"extra message: foo, 67",
			"FAIL",
		})
	})

	t.Run("MethodExpression", func(t *testing.T) {
		tt, lines := setup(t)
		actual := time.Date(2026, time.April, 9, 17, 32, 42, 123, time.UTC)
		expected := time.Date(2026, time.April, 9, 17, 32, 42, 123, time.FixedZone("My", 4*int(time.Hour.Seconds())))
		SatisfyWithf(tt, actual, expected, time.Time.Before, "extra message: %s, %d", "foo", 67)

		BeDeepEqual(t, lines(), []string{
			"actual and expected are not satisfied by predicate:",
			"actual: time.Date(2026, 4, 9, 17, 32, 42, 123, time.UTC) (time.Time)",
			"expected: time.Date(2026, 4, 9, 13, 32, 42, 123, time.UTC) (time.Time)",
			"diff expected actual",
			"--- expected",
			"+++ actual",
			"@@ -1,1 +1,1 @@",
			"-time.Date(2026, 4, 9, 13, 32, 42, 123, time.UTC) (time.Time)",
			"+time.Date(2026, 4, 9, 17, 32, 42, 123, time.UTC) (time.Time)",
			"extra message: foo, 67",
			"FAIL",
		})
	})
}

func TestNotSatisfyWithf(t *testing.T) {
	t.Run("Inline", func(t *testing.T) {
		tt, lines := setup(t)
		NotSatisfyWithf(tt, 42, 13, func(x, y int) bool { return x > y }, "extra message: %s, %d", "foo", 67)

		BeDeepEqual(t, lines(), []string{
			"actual and expected are satisfied by predicate:",
			"actual: 42 (int)",
			"expected: 13 (int)",
			"diff expected actual",
			"--- expected",
			"+++ actual",
			"@@ -1,1 +1,1 @@",
			"-13 (int)",
			"+42 (int)",
			"extra message: foo, 67",
			"FAIL",
		})
	})

	t.Run("Function", func(t *testing.T) {
		tt, lines := setup(t)
		NotSatisfyWithf(tt, 42, 13, cmp.Greater, "extra message: %s, %d", "foo", 67)

		BeDeepEqual(t, lines(), []string{
			"actual and expected are satisfied by predicate:",
			"actual: 42 (int)",
			"expected: 13 (int)",
			"diff expected actual",
			"--- expected",
			"+++ actual",
			"@@ -1,1 +1,1 @@",
			"-13 (int)",
			"+42 (int)",
			"extra message: foo, 67",
			"FAIL",
		})
	})

	t.Run("MethodExpression", func(t *testing.T) {
		tt, lines := setup(t)
		actual := time.Date(2026, time.April, 9, 13, 32, 42, 123, time.UTC)
		expected := time.Date(2026, time.April, 9, 17, 32, 42, 123, time.UTC)
		NotSatisfyWithf(tt, actual, expected, time.Time.Before, "extra message: %s, %d", "foo", 67)

		BeDeepEqual(t, lines(), []string{
			"actual and expected are satisfied by predicate:",
			"actual: time.Date(2026, 4, 9, 13, 32, 42, 123, time.UTC) (time.Time)",
			"expected: time.Date(2026, 4, 9, 17, 32, 42, 123, time.UTC) (time.Time)",
			"diff expected actual",
			"--- expected",
			"+++ actual",
			"@@ -1,1 +1,1 @@",
			"-time.Date(2026, 4, 9, 17, 32, 42, 123, time.UTC) (time.Time)",
			"+time.Date(2026, 4, 9, 13, 32, 42, 123, time.UTC) (time.Time)",
			"extra message: foo, 67",
			"FAIL",
		})
	})
}

func TestCompareEqualf(t *testing.T) {
	t.Run("MethodExpression", func(t *testing.T) {
		tt, lines := setup(t)
		actual := time.Date(2026, time.April, 9, 17, 32, 42, 123, time.UTC)
		expected := time.Date(2026, time.April, 9, 17, 32, 42, 123, time.FixedZone("My", 4*int(time.Hour.Seconds())))
		CompareEqualf(tt, actual, expected, time.Time.Compare, "extra message: %s, %d", "foo", 67)

		BeDeepEqual(t, lines(), []string{
			"actual is not equal to expected, but greater:",
			"actual: time.Date(2026, 4, 9, 17, 32, 42, 123, time.UTC) (time.Time)",
			"expected: time.Date(2026, 4, 9, 13, 32, 42, 123, time.UTC) (time.Time)",
			"diff expected actual",
			"--- expected",
			"+++ actual",
			"@@ -1,1 +1,1 @@",
			"-time.Date(2026, 4, 9, 13, 32, 42, 123, time.UTC) (time.Time)",
			"+time.Date(2026, 4, 9, 17, 32, 42, 123, time.UTC) (time.Time)",
			"extra message: foo, 67",
			"FAIL",
		})
	})

	t.Run("Format", func(t *testing.T) {
		tt, lines := setup(t)
		CompareEqualf(tt, 42, 13, cmp.Compare[int], "extra message: %s, %d", "foo", 67)

		BeDeepEqual(t, lines(), []string{
			"actual is not equal to expected, but greater:",
			"actual: 42 (int)",
			"expected: 13 (int)",
			"diff expected actual",
			"--- expected",
			"+++ actual",
			"@@ -1,1 +1,1 @@",
			"-13 (int)",
			"+42 (int)",
			"extra message: foo, 67",
			"FAIL",
		})
	})
}

func TestCompareLessf(t *testing.T) {
	t.Run("MethodExpression", func(t *testing.T) {
		tt, lines := setup(t)
		actual := time.Date(2026, time.April, 9, 17, 32, 42, 123, time.UTC)
		expected := time.Date(2026, time.April, 9, 17, 32, 42, 123, time.FixedZone("My", 4*int(time.Hour.Seconds())))
		CompareLessf(tt, actual, expected, time.Time.Compare, "extra message: %s, %d", "foo", 67)

		BeDeepEqual(t, lines(), []string{
			"actual is not less than expected, but greater:",
			"actual: time.Date(2026, 4, 9, 17, 32, 42, 123, time.UTC) (time.Time)",
			"expected: time.Date(2026, 4, 9, 13, 32, 42, 123, time.UTC) (time.Time)",
			"diff expected actual",
			"--- expected",
			"+++ actual",
			"@@ -1,1 +1,1 @@",
			"-time.Date(2026, 4, 9, 13, 32, 42, 123, time.UTC) (time.Time)",
			"+time.Date(2026, 4, 9, 17, 32, 42, 123, time.UTC) (time.Time)",
			"extra message: foo, 67",
			"FAIL",
		})
	})

	t.Run("Format", func(t *testing.T) {
		tt, lines := setup(t)
		CompareLessf(tt, 42, 13, cmp.Compare[int], "extra message: %s, %d", "foo", 67)

		BeDeepEqual(t, lines(), []string{
			"actual is not less than expected, but greater:",
			"actual: 42 (int)",
			"expected: 13 (int)",
			"diff expected actual",
			"--- expected",
			"+++ actual",
			"@@ -1,1 +1,1 @@",
			"-13 (int)",
			"+42 (int)",
			"extra message: foo, 67",
			"FAIL",
		})
	})
}

func TestCompareGreaterf(t *testing.T) {
	t.Run("MethodExpression", func(t *testing.T) {
		tt, lines := setup(t)
		actual := time.Date(2026, time.April, 9, 17, 32, 42, 123, time.FixedZone("My", 4*int(time.Hour.Seconds())))
		expected := time.Date(2026, time.April, 9, 17, 32, 42, 123, time.UTC)
		CompareGreaterf(tt, actual, expected, time.Time.Compare, "extra message: %s, %d", "foo", 67)

		BeDeepEqual(t, lines(), []string{
			"actual is not greater than expected, but less:",
			"actual: time.Date(2026, 4, 9, 13, 32, 42, 123, time.UTC) (time.Time)",
			"expected: time.Date(2026, 4, 9, 17, 32, 42, 123, time.UTC) (time.Time)",
			"diff expected actual",
			"--- expected",
			"+++ actual",
			"@@ -1,1 +1,1 @@",
			"-time.Date(2026, 4, 9, 17, 32, 42, 123, time.UTC) (time.Time)",
			"+time.Date(2026, 4, 9, 13, 32, 42, 123, time.UTC) (time.Time)",
			"extra message: foo, 67",
			"FAIL",
		})
	})

	t.Run("Format", func(t *testing.T) {
		tt, lines := setup(t)
		CompareGreaterf(tt, 13, 42, cmp.Compare[int], "extra message: %s, %d", "foo", 67)

		BeDeepEqual(t, lines(), []string{
			"actual is not greater than expected, but less:",
			"actual: 13 (int)",
			"expected: 42 (int)",
			"diff expected actual",
			"--- expected",
			"+++ actual",
			"@@ -1,1 +1,1 @@",
			"-42 (int)",
			"+13 (int)",
			"extra message: foo, 67",
			"FAIL",
		})
	})
}

func TestNotPanicf(t *testing.T) {
	t.Run("NoPanic", func(t *testing.T) {
		tt, lines := setup(t)
		NotPanicf(tt, func() {}, "extra message: %s, %d", "foo", 67)

		BeDeepEqual(t, lines(), []string{""})
	})

	t.Run("Panic", func(t *testing.T) {
		tt, lines := setup(t)
		NotPanicf(tt, func() { panic("boom") }, "extra message: %s, %d", "foo", 67)

		BeDeepEqual(t, lines(), []string{
			"function panicked:",
			`actual: "boom" (string)`,
			"extra message: foo, 67",
			"FAIL",
		})
	})

	t.Run("FormatNoPanic", func(t *testing.T) {
		tt, lines := setup(t)
		NotPanicf(tt, func() {}, "extra message: %s, %d", "foo", 67)

		BeDeepEqual(t, lines(), []string{""})
	})

	t.Run("FormatPanic", func(t *testing.T) {
		tt, lines := setup(t)
		NotPanicf(tt, func() { panic("boom") }, "extra message: %s, %d", "foo", 67)

		BeDeepEqual(t, lines(), []string{
			"function panicked:",
			`actual: "boom" (string)`,
			"extra message: foo, 67",
			"FAIL",
		})
	})
}

func TestPanicSatisfyf(t *testing.T) {
	t.Run("NoPanic", func(t *testing.T) {
		tt, lines := setup(t)
		PanicSatisfyf[any](tt, nil, func() {}, "extra message: %s, %d", "foo", 67)

		BeDeepEqual(t, lines(), []string{
			"function did not panic",
			"extra message: foo, 67",
			"FAIL",
		})
	})

	t.Run("Panic", func(t *testing.T) {
		tt, lines := setup(t)
		PanicSatisfyf[any](tt, nil, func() { panic("boom") }, "extra message: %s, %d", "foo", 67)

		BeDeepEqual(t, lines(), []string{""})
	})

	t.Run("Predicate", func(t *testing.T) {
		tt, lines := setup(t)
		var actual string
		PanicSatisfyf(tt, func(r string) bool {
			actual = r
			return true
		}, func() { panic("boom") }, "extra message: %s, %d", "foo", 67)

		BeEqual(t, actual, "boom")
		BeDeepEqual(t, lines(), []string{""})
	})

	t.Run("PredicateFail", func(t *testing.T) {
		tt, lines := setup(t)
		var actual string
		PanicSatisfyf(tt, func(r string) bool {
			actual = r
			return false
		}, func() { panic("boom") }, "extra message: %s, %d", "foo", 67)

		BeEqual(t, actual, "boom")
		BeDeepEqual(t, lines(), []string{
			"actual is not satisfied by predicate:",
			`actual: "boom" (string)`,
			"extra message: foo, 67",
			"FAIL",
		})
	})

	t.Run("Assertion", func(t *testing.T) {
		tt, lines := setup(t)
		var called bool
		PanicSatisfyf(tt, func(r string) bool {
			called = true
			BeEqual(tt, r, "boom") // no return to work with musta
			return true
		}, func() { panic("boom") }, "extra message: %s, %d", "foo", 67)

		BeTrue(t, called)
		BeDeepEqual(t, lines(), []string{""})
	})

	t.Run("AssertionFail", func(t *testing.T) {
		tt, lines := setup(t)
		var called bool
		PanicSatisfyf(tt, func(r string) bool {
			called = true
			NotBeEqual(tt, r, "boom") // no return to work with musta
			return true
		}, func() { panic("boom") }, "extra message: %s, %d", "foo", 67)

		BeTrue(t, called)
		BeDeepEqual(t, lines(), []string{
			"actual is equal to expected:",
			`actual: "boom" (string)`,
			`expected: "boom" (string)`,
			"FAIL",
		})
	})

	t.Run("WrongType", func(t *testing.T) {
		tt, lines := setup(t)
		var called bool
		PanicSatisfyf(tt, func(r int) bool {
			called = true
			return true
		}, func() { panic("boom") }, "extra message: %s, %d", "foo", 67)

		BeFalse(t, called)
		BeDeepEqual(t, lines(), []string{
			"actual panic value is not of type int, but:",
			`actual: "boom" (string)`,
			"extra message: foo, 67",
			"FAIL",
		})
	})

	t.Run("WrongInterfaceType", func(t *testing.T) {
		tt, lines := setup(t)
		var called bool
		PanicSatisfyf(tt, func(r error) bool {
			called = true
			return true
		}, func() { panic("boom") }, "extra message: %s, %d", "foo", 67)

		BeFalse(t, called)
		BeDeepEqual(t, lines(), []string{
			"actual panic value is not of type error, but:",
			`actual: "boom" (string)`,
			"extra message: foo, 67",
			"FAIL",
		})
	})

	t.Run("NoPanic", func(t *testing.T) {
		tt, lines := setup(t)
		PanicSatisfyf[any](tt, nil, func() {}, "extra message: %s, %d", "foo", 67)

		BeDeepEqual(t, lines(), []string{
			"function did not panic",
			"extra message: foo, 67",
			"FAIL",
		})
	})

	t.Run("Panic", func(t *testing.T) {
		tt, lines := setup(t)
		PanicSatisfyf[any](tt, nil, func() { panic("boom") }, "extra message: %s, %d", "foo", 67)

		BeDeepEqual(t, lines(), []string{""})
	})

	t.Run("PredicateFail", func(t *testing.T) {
		tt, lines := setup(t)
		PanicSatisfyf(
			tt, func(string) bool { return false }, func() { panic("boom") },
			"extra message: %s, %d", "foo", 67,
		)

		BeDeepEqual(t, lines(), []string{
			"actual is not satisfied by predicate:",
			`actual: "boom" (string)`,
			"extra message: foo, 67",
			"FAIL",
		})
	})

	t.Run("WrongType", func(t *testing.T) {
		tt, lines := setup(t)
		PanicSatisfyf(
			tt, func(int) bool { return true }, func() { panic("boom") },
			"extra message: %s, %d", "foo", 67,
		)

		BeDeepEqual(t, lines(), []string{
			"actual panic value is not of type int, but:",
			`actual: "boom" (string)`,
			"extra message: foo, 67",
			"FAIL",
		})
	})
}
