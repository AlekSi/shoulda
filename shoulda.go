// Package shoulda provides an assertion library that is as simple as possible, but not simpler.
package shoulda

import (
	"fmt"
	"testing"

	"github.com/AlekSi/shoulda/internal"
)

// TB is a subset of [testing.TB] that is sufficient for assertions.
type TB interface {
	Helper()
	Log(args ...any)
	Fail()
}

// messager is an interface for lazy message construction.
type messager interface {
	message() string
}

// msgString implements [messager] for a simple string message.
type msgString string

// message implements [messager].
func (m msgString) message() string { return string(m) }

// msgFunc implements [messager] for a function that constructs a message.
type msgFunc func() string

// message implements [messager].
func (m msgFunc) message() string { return m() }

// msgFmt constructs a message from [fmt.Sprintf]-like arguments.
func msgFmt(msg string, args ...any) messager {
	return msgFunc(func() string {
		return fmt.Sprintf(msg, args...)
	})
}

// assert returns true if condition is true;
// otherwise it logs msg, fails test, and returns false.
func assert(tb TB, condition bool, msg messager) bool {
	tb.Helper()

	if condition {
		return true
	}

	tb.Log(msg.message())
	tb.Fail()
	return false
}

// check interfaces
var (
	_ TB = (testing.TB)(nil)
	_ TB = internal.TestTB{}
)
