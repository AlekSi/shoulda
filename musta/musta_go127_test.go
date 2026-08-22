//go:build go1.27

package musta

import (
	"mime"
	"strconv"
	"testing"
)

func TestMustNotFail(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		tt, lines := setup(t)
		actual := T(tt).NotFail(strconv.Atoi("42"))

		BeEqual(t, actual, 42)
		BeDeepEqual(t, lines(), []string{""})
	})

	t.Run("Failure", func(t *testing.T) {
		tt, lines := setup(t)
		T(tt).NotFail(strconv.Atoi("foo"))

		BeDeepEqual(t, lines(), []string{
			"actual is not nil error:",
			`actual: strconv.Atoi: parsing "foo": invalid syntax`,
			"&strconv.NumError{",
			`  Func: "Atoi",`,
			`  Num: "foo",`,
			"  Err: &errors.errorString{",
			`    s: "invalid syntax",`,
			"  },",
			"} (*strconv.NumError)",
			"FAIL",
		})
	})
}

func TestMustNotFail2(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		tt, lines := setup(t)
		mediaType, params := T(tt).NotFail2(mime.ParseMediaType("text/html; charset=utf-8"))

		BeEqual(t, mediaType, "text/html")
		BeDeepEqual(t, params, map[string]string{"charset": "utf-8"})
		BeDeepEqual(t, lines(), []string{""})
	})

	t.Run("Failure", func(t *testing.T) {
		tt, lines := setup(t)
		T(tt).NotFail2(mime.ParseMediaType(""))

		BeDeepEqual(t, lines(), []string{
			"actual is not nil error:",
			"actual: mime: no media type",
			"&errors.errorString{",
			`  s: "mime: no media type",`,
			"} (*errors.errorString)",
			"FAIL",
		})
	})
}

func TestMustNotFail3(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		tt, lines := setup(t)
		value, multibyte, tail := T(tt).NotFail3(strconv.UnquoteChar(`\u263a!`, 0))

		BeEqual(t, value, '☺')
		BeTrue(t, multibyte)
		BeEqual(t, tail, "!")
		BeDeepEqual(t, lines(), []string{""})
	})

	t.Run("Failure", func(t *testing.T) {
		tt, lines := setup(t)
		T(tt).NotFail3(strconv.UnquoteChar("", 0))

		BeDeepEqual(t, lines(), []string{
			"actual is not nil error:",
			"actual: invalid syntax",
			"&errors.errorString{",
			`  s: "invalid syntax",`,
			"} (*errors.errorString)",
			"FAIL",
		})
	})
}
