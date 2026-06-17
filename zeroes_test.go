package shoulda

import (
	"errors"
	"testing"
)

func TestBeNil(t *testing.T) {
	t.Run("Value", func(t *testing.T) {
		tt, actual := setup(t)
		BeNil(tt, uint32(13))

		BeDeepEqual(t, actual(), []string{
			"actual is not untyped nil, but:",
			"actual: 13 (uint32)",
			"FAIL",
		})
	})

	t.Run("Pointer", func(t *testing.T) {
		tt, actual := setup(t)
		BeNil(tt, new(uint32(13)))

		BeDeepEqual(t, actual(), []string{
			"actual is not untyped nil, but:",
			"actual: &13 (*uint32)",
			"FAIL",
		})
	})

	t.Run("UntypedNil", func(t *testing.T) {
		tt, actual := setup(t)
		BeNil(tt, nil)

		BeDeepEqual(t, actual(), []string{""})
	})

	t.Run("TypedNil", func(t *testing.T) {
		tt, actual := setup(t)
		BeNil(tt, (*uint32)(nil))

		BeDeepEqual(t, actual(), []string{
			"actual is not untyped nil, but:",
			"actual: nil (*uint32)",
			"FAIL",
		})
	})

	t.Run("Error", func(t *testing.T) {
		tt, actual := setup(t)
		BeNil(tt, errors.New("boom"))

		BeDeepEqual(t, actual(), []string{
			`actual is not untyped nil, but:`,
			`actual: &errors.errorString{`,
			`  s: "boom",`,
			`} (*errors.errorString)`,
			"FAIL",
		})
	})
}

func TestBeNilf(t *testing.T) {
	tt, actual := setup(t)
	BeNilf(tt, uint32(13), "extra message: %s, %d", "foo", 42)

	BeDeepEqual(t, actual(), []string{
		"actual is not untyped nil, but:",
		"actual: 13 (uint32)",
		"extra message: foo, 42",
		"FAIL",
	})
}

func TestNotBeNil(t *testing.T) {
	t.Run("Value", func(t *testing.T) {
		tt, actual := setup(t)
		NotBeNil(tt, uint32(13))

		BeDeepEqual(t, actual(), []string{""})
	})

	t.Run("Pointer", func(t *testing.T) {
		tt, actual := setup(t)
		NotBeNil(tt, new(uint32(13)))

		BeDeepEqual(t, actual(), []string{""})
	})

	t.Run("UntypedNil", func(t *testing.T) {
		tt, actual := setup(t)
		NotBeNil(tt, nil)

		BeDeepEqual(t, actual(), []string{
			"actual is untyped nil",
			"FAIL",
		})
	})

	t.Run("TypedNil", func(t *testing.T) {
		tt, actual := setup(t)
		NotBeNil(tt, (*uint32)(nil))

		BeDeepEqual(t, actual(), []string{""})
	})

	t.Run("Error", func(t *testing.T) {
		tt, actual := setup(t)
		NotBeNil(tt, errors.New("boom"))

		BeDeepEqual(t, actual(), []string{""})
	})
}

func TestNotBeNilf(t *testing.T) {
	tt, actual := setup(t)
	NotBeNilf(tt, nil, "extra message: %s, %d", "foo", 42)

	BeDeepEqual(t, actual(), []string{
		"actual is untyped nil",
		"extra message: foo, 42",
		"FAIL",
	})
}

func TestBeZero(t *testing.T) {
	t.Run("Simple", func(t *testing.T) {
		tt, actual := setup(t)
		BeZero(tt, 13)

		BeDeepEqual(t, actual(), []string{
			"actual is not zero, but:",
			"actual: 13 (int)",
			"FAIL",
		})
	})

	t.Run("Nil", func(t *testing.T) {
		tt, actual := setup(t)
		BeZero(tt, (*int)(nil))

		BeDeepEqual(t, actual(), []string{""})
	})
}

func TestBeZerof(t *testing.T) {
	tt, actual := setup(t)
	BeZerof(tt, 13, "extra message: %s, %d", "foo", 42)

	BeDeepEqual(t, actual(), []string{
		"actual is not zero, but:",
		"actual: 13 (int)",
		"extra message: foo, 42",
		"FAIL",
	})
}

func TestNotBeZero(t *testing.T) {
	t.Run("Simple", func(t *testing.T) {
		tt, actual := setup(t)
		NotBeZero(tt, 0)

		BeDeepEqual(t, actual(), []string{
			"actual is zero",
			"FAIL",
		})
	})

	t.Run("Nil", func(t *testing.T) {
		tt, actual := setup(t)
		NotBeZero(tt, (*int)(nil))

		BeDeepEqual(t, actual(), []string{
			"actual is zero",
			"FAIL",
		})
	})
}

func TestNotBeZerof(t *testing.T) {
	tt, actual := setup(t)
	NotBeZerof(tt, 0, "extra message: %s, %d", "foo", 42)

	BeDeepEqual(t, actual(), []string{
		"actual is zero",
		"extra message: foo, 42",
		"FAIL",
	})
}

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
