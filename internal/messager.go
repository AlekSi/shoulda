package internal

import (
	"fmt"
)

// Messager is an interface for lazy message construction.
type Messager interface {
	Message() string
}

// F constructs a Messager from a format string and arguments.
func F(msg string, args ...any) Messager {
	if len(args) == 0 {
		return msgString(msg)
	}

	return msgFunc(func() string {
		return fmt.Sprintf(msg, args...)
	})
}

// msgString implements [Messager] for a simple string message.
type msgString string

// Message implements [Messager].
func (m msgString) Message() string { return string(m) }

// msgFunc implements [Messager] for a function that constructs a message.
type msgFunc func() string

// Message implements [Messager].
func (m msgFunc) Message() string { return m() }
