//go:build go1.27

package musta

import (
	"errors"
	"testing"
)

func TestMustNotFail(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		tt, lines := setup(t)
		f := func() (int, error) { return 42, nil }
		actual := With(tt).NotFail(f())

		BeEqual(t, actual, 42)
		BeDeepEqual(t, lines(), []string{""})
	})

	t.Run("Failure", func(t *testing.T) {
		tt, lines := setup(t)
		With(tt).NotFail(42, errors.New("boom"))

		BeDeepEqual(t, lines(), []string{
			"actual is not nil error:",
			"actual: boom",
			"&errors.errorString{",
			`  s: "boom",`,
			"} (*errors.errorString)",
			"FAIL",
		})
	})
}

func TestAssertionsNotFail2(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		tt, lines := setup(t)
		f := func() (int, string, error) { return 42, "foo", nil }
		actual1, actual2 := With(tt).NotFail2(f())

		BeEqual(t, actual1, 42)
		BeEqual(t, actual2, "foo")
		BeDeepEqual(t, lines(), []string{""})
	})

	t.Run("Failure", func(t *testing.T) {
		tt, lines := setup(t)
		With(tt).NotFail2(42, "foo", errors.New("boom"))

		BeDeepEqual(t, lines(), []string{
			"actual is not nil error:",
			"actual: boom",
			"&errors.errorString{",
			`  s: "boom",`,
			"} (*errors.errorString)",
			"FAIL",
		})
	})
}

func TestAssertionsNotFail3(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		tt, lines := setup(t)
		f := func() (int, string, bool, error) { return 42, "foo", true, nil }
		actual1, actual2, actual3 := With(tt).NotFail3(f())

		BeEqual(t, actual1, 42)
		BeEqual(t, actual2, "foo")
		BeTrue(t, actual3)
		BeDeepEqual(t, lines(), []string{""})
	})

	t.Run("Failure", func(t *testing.T) {
		tt, lines := setup(t)
		With(tt).NotFail3(42, "foo", true, errors.New("boom"))

		BeDeepEqual(t, lines(), []string{
			"actual is not nil error:",
			"actual: boom",
			"&errors.errorString{",
			`  s: "boom",`,
			"} (*errors.errorString)",
			"FAIL",
		})
	})
}
