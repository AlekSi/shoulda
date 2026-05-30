package shoulda

import (
	"testing"
	"time"

	"github.com/AlekSi/shoulda/cmp"
)

func TestSatisfy(t *testing.T) {
	t.Run("Inline", func(t *testing.T) {
		tt, lines := setup(t)
		Satisfy(tt, 13, func(v int) bool { return v > 42 })

		BeDeepEqual(t, lines(), []string{
			"actual is not satisfied by predicate:",
			"actual: 13 (int)",
			"FAIL",
		})
	})

	t.Run("MethodValue", func(t *testing.T) {
		tt, lines := setup(t)
		actual := time.Date(2026, time.April, 9, 17, 32, 42, 123, time.UTC)
		Satisfy(tt, actual, time.Now().Before)

		BeDeepEqual(t, lines(), []string{
			"actual is not satisfied by predicate:",
			"actual: time.Date(2026, 4, 9, 17, 32, 42, 123, time.UTC) (time.Time)",
			"FAIL",
		})
	})

	t.Run("MethodExpression", func(t *testing.T) {
		tt, lines := setup(t)
		actual := time.Date(2026, time.April, 9, 17, 32, 42, 123, time.UTC)
		Satisfy(tt, actual, time.Time.IsZero)

		BeDeepEqual(t, lines(), []string{
			"actual is not satisfied by predicate:",
			"actual: time.Date(2026, 4, 9, 17, 32, 42, 123, time.UTC) (time.Time)",
			"FAIL",
		})
	})
}

func TestSatisfyWith(t *testing.T) {
	t.Run("Inline", func(t *testing.T) {
		tt, lines := setup(t)
		SatisfyWith(tt, 13, 42, func(x, y int) bool { return x > y })

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
			"FAIL",
		})
	})

	t.Run("Function", func(t *testing.T) {
		tt, lines := setup(t)
		SatisfyWith(tt, 13, 42, cmp.Greater)

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
			"FAIL",
		})
	})

	t.Run("MethodExpression", func(t *testing.T) {
		tt, lines := setup(t)
		actual := time.Date(2026, time.April, 9, 17, 32, 42, 123, time.UTC)
		expected := time.Date(2026, time.April, 9, 17, 32, 42, 123, time.FixedZone("My", 4*int(time.Hour.Seconds())))
		SatisfyWith(tt, actual, expected, time.Time.Before)

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
			"FAIL",
		})
	})
}

func TestCompareWith(t *testing.T) {
	t.Run("Function", func(t *testing.T) {
		tt, lines := setup(t)
		CompareWith(tt, 42, 13, cmp.OrderLess, cmp.Compare[int])

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
			"FAIL",
		})
	})
}

func TestCompareEqual(t *testing.T) {
	t.Run("MethodExpression", func(t *testing.T) {
		tt, lines := setup(t)
		actual := time.Date(2026, time.April, 9, 17, 32, 42, 123, time.UTC)
		expected := time.Date(2026, time.April, 9, 17, 32, 42, 123, time.FixedZone("My", 4*int(time.Hour.Seconds())))
		CompareEqual(tt, actual, expected, time.Time.Compare)

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
			"FAIL",
		})
	})
}

func TestCompareLess(t *testing.T) {
	t.Run("MethodExpression", func(t *testing.T) {
		tt, lines := setup(t)
		actual := time.Date(2026, time.April, 9, 17, 32, 42, 123, time.UTC)
		expected := time.Date(2026, time.April, 9, 17, 32, 42, 123, time.FixedZone("My", 4*int(time.Hour.Seconds())))
		CompareLess(tt, actual, expected, time.Time.Compare)

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
			"FAIL",
		})
	})
}

func TestCompareGreater(t *testing.T) {
	t.Run("MethodExpression", func(t *testing.T) {
		tt, lines := setup(t)
		actual := time.Date(2026, time.April, 9, 17, 32, 42, 123, time.FixedZone("My", 4*int(time.Hour.Seconds())))
		expected := time.Date(2026, time.April, 9, 17, 32, 42, 123, time.UTC)
		CompareGreater(tt, actual, expected, time.Time.Compare)

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
			"FAIL",
		})
	})
}
