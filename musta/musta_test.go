package musta

import (
	"errors"
	"testing"
)

func TestNotFail(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		tt, lines := setup(t)
		f := func() (int, error) { return 42, nil }
		actual := NotFail(f())(tt)

		BeEqual(t, actual, 42)
		BeDeepEqual(t, lines(), []string{""})
	})

	t.Run("Failure", func(t *testing.T) {
		tt, lines := setup(t)
		NotFail(42, errors.New("boom"))(tt)

		BeDeepEqual(t, lines(), []string{
			`actual is not nil error, but "boom":`,
			"actual: &errors.errorString{",
			`  s: "boom",`,
			"} (*errors.errorString)",
			"FAIL",
		})
	})
}

func TestNotFail2(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		tt, lines := setup(t)
		f := func() (int, string, error) { return 42, "foo", nil }
		actual1, actual2 := NotFail2(f())(tt)

		BeEqual(t, actual1, 42)
		BeEqual(t, actual2, "foo")
		BeDeepEqual(t, lines(), []string{""})
	})

	t.Run("Failure", func(t *testing.T) {
		tt, lines := setup(t)
		NotFail2(42, "foo", errors.New("boom"))(tt)

		BeDeepEqual(t, lines(), []string{
			`actual is not nil error, but "boom":`,
			"actual: &errors.errorString{",
			`  s: "boom",`,
			"} (*errors.errorString)",
			"FAIL",
		})
	})
}

func TestNotFail3(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		tt, lines := setup(t)
		f := func() (int, string, bool, error) { return 42, "foo", true, nil }
		actual1, actual2, actual3 := NotFail3(f())(tt)

		BeEqual(t, actual1, 42)
		BeEqual(t, actual2, "foo")
		BeTrue(t, actual3)
		BeDeepEqual(t, lines(), []string{""})
	})

	t.Run("Failure", func(t *testing.T) {
		tt, lines := setup(t)
		NotFail3(42, "foo", true, errors.New("boom"))(tt)

		BeDeepEqual(t, lines(), []string{
			`actual is not nil error, but "boom":`,
			"actual: &errors.errorString{",
			`  s: "boom",`,
			"} (*errors.errorString)",
			"FAIL",
		})
	})
}
