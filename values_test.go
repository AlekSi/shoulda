package shoulda

import (
	"bytes"
	"math"
	"strings"
	"testing"
	"time"

	"github.com/AlekSi/shoulda/internal"
)

// setup returns an [internal.TestTB] and a function to get the output lines
// for the given test.
func setup(t *testing.T) (internal.TestTB, func() []string) {
	t.Helper()

	var buf bytes.Buffer
	tt := internal.TestTB{
		W: &buf,
	}
	f := func() []string {
		s := strings.TrimRight(buf.String(), "\n")
		return strings.Split(s, "\n")
	}
	return tt, f
}

func TestBeFalse(t *testing.T) {
	t.Run("Simple", func(t *testing.T) {
		tt, lines := setup(t)
		actual := time.Date(2026, time.April, 9, 17, 32, 42, 123, time.UTC)
		expected := time.Date(2026, time.April, 9, 21, 32, 42, 123, time.FixedZone("My", 4*int(time.Hour.Seconds())))
		BeFalse(tt, time.Time.Equal(actual, expected))

		BeDeepEqual(t, lines(), []string{
			"actual is not false",
			"FAIL",
		})
	})
}

func TestBeTrue(t *testing.T) {
	t.Run("Simple", func(t *testing.T) {
		tt, lines := setup(t)
		actual := time.Date(2026, time.April, 9, 17, 32, 42, 123, time.UTC)
		expected := time.Date(2026, time.April, 9, 17, 32, 42, 123, time.FixedZone("My", 4*int(time.Hour.Seconds())))
		BeTrue(tt, time.Time.Equal(actual, expected))

		BeDeepEqual(t, lines(), []string{
			"actual is not true",
			"FAIL",
		})
	})
}

func TestBeDeepEqual(t *testing.T) {
	t.Run("Simple", func(t *testing.T) {
		tt, lines := setup(t)
		BeDeepEqual(tt, []int{13}, []int64{13})

		BeDeepEqual(t, lines(), []string{
			"actual is not deep equal to expected:",
			"actual: []int{13}",
			"expected: []int64{13}",
			"FAIL",
		})
	})

	t.Run("NaN", func(t *testing.T) {
		tt, lines := setup(t)
		BeDeepEqual(tt, []float64{math.NaN()}, []float64{math.NaN()})

		BeDeepEqual(t, lines(), []string{
			"actual is not deep equal to expected:",
			"actual: []float64{NaN}",
			"expected: []float64{NaN}",
			"FAIL",
		})
	})
}

func TestNotBeDeepEqual(t *testing.T) {
	t.Run("Simple", func(t *testing.T) {
		tt, lines := setup(t)
		NotBeDeepEqual(tt, []int{13}, []int{13})

		BeDeepEqual(t, lines(), []string{
			"Values are deep equal:",
			"actual:   []int{13}",
			"expected: []int{13}",
			"FAIL",
		})
	})

	t.Run("NaN", func(t *testing.T) {
		tt, lines := setup(t)
		NotBeDeepEqual(tt, []float64{math.NaN()}, []float64{math.NaN()})

		BeDeepEqual(t, lines(), []string{""})
	})
}

func TestBeEqual(t *testing.T) {
	t.Run("Simple", func(t *testing.T) {
		tt, lines := setup(t)
		BeEqual(tt, 13, 42)

		BeDeepEqual(t, lines(), []string{
			"Values are not equal:",
			"actual:   13",
			"expected: 42",
			"FAIL",
		})
	})

	t.Run("NaN", func(t *testing.T) {
		tt, lines := setup(t)
		BeEqual(tt, math.NaN(), math.NaN())

		BeDeepEqual(t, lines(), []string{""})
	})
}

func TestBeLess(t *testing.T) {
	t.Run("Simple", func(t *testing.T) {
		tt, lines := setup(t)
		BeLess(tt, 42, 13)

		BeDeepEqual(t, lines(), []string{
			"actual:   42",
			"is not less than",
			"expected: 13",
			"FAIL",
		})
	})
}

func TestBeGreater(t *testing.T) {
	t.Run("Simple", func(t *testing.T) {
		tt, lines := setup(t)
		BeGreater(tt, 13, 42)

		BeDeepEqual(t, lines(), []string{
			"actual:   13",
			"is not greater than",
			"expected: 42",
			"FAIL",
		})
	})
}
