package musta

import (
	"bytes"
	"strings"
	"testing"

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

func TestBeDeepEqual(t *testing.T) {
	t.Run("Fail", func(t *testing.T) {
		tt, lines := setup(t)
		BeDeepEqual(tt, []int{13}, []int64{13})

		BeDeepEqual(t, lines(), []string{
			"Values are not deep equal:",
			"actual:   []int{13}",
			"expected: []int64{13}",
			"FAIL",
		})
	})
}
