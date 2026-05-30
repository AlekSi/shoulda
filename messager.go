package shoulda

import (
	"fmt"
)

// messager is an interface for lazy message construction.
type messager interface {
	Message() string
}

// msgString implements [messager] for a simple string message.
type msgString string

// Message implements [messager].
func (m msgString) Message() string { return string(m) }

// msgFunc implements [messager] for a function that constructs a message.
type msgFunc func() string

// Message implements [messager].
func (m msgFunc) Message() string { return m() }

// messagef constructs a [messager] from a message or format string, and arguments.
func messagef(msg string, args ...any) messager {
	if len(args) == 0 {
		return msgString(msg)
	}

	return msgFunc(func() string {
		return fmt.Sprintf(msg, args...)
	})
}

// dumpf constructs a [messager] from a format string and values.
func dumpf(tb TB, format string, vs ...any) messager {
	return msgFunc(func() string {
		tb.Helper()

		var args []any
		for _, v := range vs {
			args = append(args, v, Dump(tb, v))
		}

		return fmt.Sprintf(format, args...)
	})
}

// // msgDiff implements [messager] using [Diff].
// func msgDiff(tb TB, actualName string, actual []byte, expectedName string, expected []byte) messager {
// 	return msgFunc(func() string {
// 		tb.Helper()
// 		return string(Diff(tb, actualName, actual, expectedName, expected))
// 	})
// }.
