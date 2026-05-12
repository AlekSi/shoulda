// Package shoulda provides an assertion library that is as simple as possible, but not simpler.
package shoulda

import (
	"testing"

	"github.com/AlekSi/shoulda/internal"
)

// TB is a subset of [testing.TB] that is sufficient for assertions.
type TB interface {
	Helper()
	Log(args ...any)
	Fail()
}

// assert returns true if condition is true;
// otherwise it logs msg, fails test, and returns false.
func assert(tb TB, condition bool, msg internal.Messager) bool {
	tb.Helper()

	if condition {
		return true
	}

	tb.Log(msg.Message())
	tb.Fail()
	return false
}

// check interfaces
var (
	_ TB = (testing.TB)(nil)
	_ TB = internal.TestTB{}
)
