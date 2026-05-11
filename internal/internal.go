// Package internal contains unstable code.
package internal

import (
	"cmp"
	"fmt"
	"io"
	"os"
)

// TestTB implements [shoulda.TB] and [musta.TB] for testing and examples.
type TestTB struct {
	W io.Writer
}

// Helper implements [shoulda.TB.Helper] and [musta.TB.Helper] by doing nothing.
func (t TestTB) Helper() {}

// FailNow implements [shoulda.TB.Helper] and [musta.TB.Helper] by doing nothing.
func (t TestTB) FailNow() {}

// Log implements [shoulda.TB.Helper] and [musta.TB.Helper].
func (t TestTB) Log(args ...any) {
	fmt.Fprintln(cmp.Or(t.W, io.Writer(os.Stdout)), args...)
}

// Fail implements [shoulda.TB.Helper] and [musta.TB.Helper].
func (t TestTB) Fail() {
	fmt.Fprintln(cmp.Or(t.W, io.Writer(os.Stdout)), "FAIL")
}
