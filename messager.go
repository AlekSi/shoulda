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

// messagef constructs a [messager] from a message string and arguments.
func messagef(msg string, args ...any) messager {
	if len(args) == 0 {
		return msgString(msg)
	}

	return msgFunc(func() string {
		return fmt.Sprintf(msg, args...)
	})
}

// dumpf implements [messager] from a message string and arguments using [Dump].
// args is prepended with the value and its dump string.
func dumpf(tb TB, msg string, v any, args ...any) messager {
	return msgFunc(func() string {
		tb.Helper()
		s := Dump(tb, v)
		args = append([]any{v, s}, args...)
		return fmt.Sprintf(msg, args...)
	})
}

// // msgDiff implements [messager] using [Diff].
// func msgDiff(tb TB, actualName string, actual []byte, expectedName string, expected []byte) messager {
// 	return msgFunc(func() string {
// 		tb.Helper()
// 		return string(Diff(tb, actualName, actual, expectedName, expected))
// 	})
// }.
