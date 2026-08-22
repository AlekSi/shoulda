package musta

import (
	"mime"
	"strconv"
	"testing"
)

func TestNotFail(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		tt, lines := setup(t)
		actual := NotFail(strconv.Atoi("42"))(tt)

		BeEqual(t, actual, 42)
		BeDeepEqual(t, lines(), []string{""})
	})

	t.Run("Failure", func(t *testing.T) {
		tt, lines := setup(t)
		NotFail(strconv.Atoi("foo"))(tt)

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

func TestNotFail2(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		tt, lines := setup(t)
		mediaType, params := NotFail2(mime.ParseMediaType("text/html; charset=utf-8"))(tt)

		BeEqual(t, mediaType, "text/html")
		BeDeepEqual(t, params, map[string]string{"charset": "utf-8"})
		BeDeepEqual(t, lines(), []string{""})
	})

	t.Run("Failure", func(t *testing.T) {
		tt, lines := setup(t)
		NotFail2(mime.ParseMediaType(""))(tt)

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

func TestNotFail3(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		tt, lines := setup(t)
		value, multibyte, tail := NotFail3(strconv.UnquoteChar(`\u263a!`, 0))(tt)

		BeEqual(t, value, '☺')
		BeTrue(t, multibyte)
		BeEqual(t, tail, "!")
		BeDeepEqual(t, lines(), []string{""})
	})

	t.Run("Failure", func(t *testing.T) {
		tt, lines := setup(t)
		NotFail3(strconv.UnquoteChar("", 0))(tt)

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
