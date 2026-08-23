package shoulda

import (
	"errors"
	"testing"
)

func TestErrorf(t *testing.T) {
	t.Run("Simple", func(t *testing.T) {
		tt, actual := setup(t)
		Errorf(tt, errors.New("boom"), "extra message: %s, %d", "foo", 67)

		BeDeepEqual(t, actual(), []string{""})
	})

	t.Run("Nil", func(t *testing.T) {
		tt, actual := setup(t)
		Errorf(tt, nil, "extra message: %s, %d", "foo", 67)

		BeDeepEqual(t, actual(), []string{
			"actual is nil error",
			"extra message: foo, 67",
			"FAIL",
		})
	})

	t.Run("Format", func(t *testing.T) {
		tt, actual := setup(t)
		Errorf(tt, nil, "extra message: %s, %d", "foo", 67)

		BeDeepEqual(t, actual(), []string{
			"actual is nil error",
			"extra message: foo, 67",
			"FAIL",
		})
	})
}

func TestErrorIsf(t *testing.T) {
	expected := errors.New("expected")

	t.Run("Same", func(t *testing.T) {
		tt, actual := setup(t)
		ErrorIsf(tt, expected, expected, "extra message: %s, %d", "foo", 67)

		BeDeepEqual(t, actual(), []string{""})
	})

	t.Run("Wrapped", func(t *testing.T) {
		tt, actual := setup(t)
		ErrorIsf(tt, errors.Join(errors.New("other"), expected), expected, "extra message: %s, %d", "foo", 67)

		BeDeepEqual(t, actual(), []string{""})
	})

	t.Run("Different", func(t *testing.T) {
		tt, actual := setup(t)
		ErrorIsf(tt, errors.New("boom"), expected, "extra message: %s, %d", "foo", 67)

		BeDeepEqual(t, actual(), []string{
			"actual error does not match expected:",
			"actual: boom",
			"&errors.errorString{",
			`  s: "boom",`,
			"} (*errors.errorString)",
			"expected: expected",
			"&errors.errorString{",
			`  s: "expected",`,
			"} (*errors.errorString)",
			"extra message: foo, 67",
			"FAIL",
		})
	})

	t.Run("Nil", func(t *testing.T) {
		tt, actual := setup(t)
		ErrorIsf(tt, nil, expected, "extra message: %s, %d", "foo", 67)

		BeDeepEqual(t, actual(), []string{
			"actual error does not match expected:",
			"actual: <nil>",
			"nil (<nil>)",
			"expected: expected",
			"&errors.errorString{",
			`  s: "expected",`,
			"} (*errors.errorString)",
			"extra message: foo, 67",
			"FAIL",
		})
	})

	t.Run("Format", func(t *testing.T) {
		tt, actual := setup(t)
		ErrorIsf(tt, errors.New("boom"), errors.New("target"), "extra message: %s, %d", "foo", 67)

		BeDeepEqual(t, actual(), []string{
			"actual error does not match expected:",
			"actual: boom",
			"&errors.errorString{",
			`  s: "boom",`,
			"} (*errors.errorString)",
			"expected: target",
			"&errors.errorString{",
			`  s: "target",`,
			"} (*errors.errorString)",
			"extra message: foo, 67",
			"FAIL",
		})
	})
}

func TestNoErrorf(t *testing.T) {
	t.Run("Simple", func(t *testing.T) {
		tt, actual := setup(t)
		NoErrorf(tt, errors.New("boom"), "extra message: %s, %d", "foo", 67)

		BeDeepEqual(t, actual(), []string{
			"actual is not nil error:",
			"actual: boom",
			"&errors.errorString{",
			`  s: "boom",`,
			`} (*errors.errorString)`,
			"extra message: foo, 67",
			"FAIL",
		})
	})

	t.Run("Nil", func(t *testing.T) {
		tt, actual := setup(t)
		NoErrorf(tt, nil, "extra message: %s, %d", "foo", 67)

		BeDeepEqual(t, actual(), []string{""})
	})

	t.Run("Format", func(t *testing.T) {
		tt, actual := setup(t)
		NoErrorf(tt, errors.New("boom"), "extra message: %s, %d", "foo", 67)

		BeDeepEqual(t, actual(), []string{
			"actual is not nil error:",
			"actual: boom",
			"&errors.errorString{",
			`  s: "boom",`,
			`} (*errors.errorString)`,
			"extra message: foo, 67",
			"FAIL",
		})
	})
}
