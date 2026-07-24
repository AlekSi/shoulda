package shoulda

import (
	"errors"
	"testing"
)

func TestError(t *testing.T) {
	t.Run("Simple", func(t *testing.T) {
		tt, actual := setup(t)
		Error(tt, errors.New("boom"))

		BeDeepEqual(t, actual(), []string{""})
	})

	t.Run("Nil", func(t *testing.T) {
		tt, actual := setup(t)
		Error(tt, nil)

		BeDeepEqual(t, actual(), []string{
			"actual is nil error",
			"FAIL",
		})
	})
}

func TestErrorf(t *testing.T) {
	tt, actual := setup(t)
	Errorf(tt, nil, "extra message: %s, %d", "foo", 42)

	BeDeepEqual(t, actual(), []string{
		"actual is nil error",
		"extra message: foo, 42",
		"FAIL",
	})
}

func TestErrorIs(t *testing.T) {
	expected := errors.New("expected")

	t.Run("Same", func(t *testing.T) {
		tt, actual := setup(t)
		ErrorIs(tt, expected, expected)

		BeDeepEqual(t, actual(), []string{""})
	})

	t.Run("Wrapped", func(t *testing.T) {
		tt, actual := setup(t)
		ErrorIs(tt, errors.Join(errors.New("other"), expected), expected)

		BeDeepEqual(t, actual(), []string{""})
	})

	t.Run("Different", func(t *testing.T) {
		tt, actual := setup(t)
		ErrorIs(tt, errors.New("actual"), expected)

		BeDeepEqual(t, actual(), []string{
			"actual does not match expected:",
			"actual: &errors.errorString{",
			`  s: "actual",`,
			"} (*errors.errorString)",
			"expected: &errors.errorString{",
			`  s: "expected",`,
			"} (*errors.errorString)",
			"FAIL",
		})
	})
}

func TestNoError(t *testing.T) {
	t.Run("Simple", func(t *testing.T) {
		tt, actual := setup(t)
		NoError(tt, errors.New("boom"))

		BeDeepEqual(t, actual(), []string{
			`actual is not nil error, but "boom":`,
			`actual: &errors.errorString{`,
			`  s: "boom",`,
			`} (*errors.errorString)`,
			"FAIL",
		})
	})

	t.Run("Nil", func(t *testing.T) {
		tt, actual := setup(t)
		NoError(tt, nil)

		BeDeepEqual(t, actual(), []string{""})
	})
}

func TestNoErrorf(t *testing.T) {
	tt, actual := setup(t)
	NoErrorf(tt, errors.New("boom"), "extra message: %s, %d", "foo", 42)

	BeDeepEqual(t, actual(), []string{
		`actual is not nil error, but "boom":`,
		`actual: &errors.errorString{`,
		`  s: "boom",`,
		`} (*errors.errorString)`,
		"extra message: foo, 42",
		"FAIL",
	})
}
