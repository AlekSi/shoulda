package internal

import (
	"fmt"
)

// Messager is an interface for lazy message construction.
type Messager interface {
	Message() string
}

// MsgString implements [messager] for a simple string message.
type MsgString string

// Message implements [messager].
func (m MsgString) Message() string { return string(m) }

// msgFunc implements [messager] for a function that constructs a message.
type msgFunc func() string

// Message implements [messager].
func (m msgFunc) Message() string { return m() }

// MsgFmt constructs a message from [fmt.Sprintf]-like arguments.
func MsgFmt(msg string, args ...any) Messager {
	return msgFunc(func() string {
		return fmt.Sprintf(msg, args...)
	})
}

// check interfaces
var (
	_ Messager = MsgString("")
	_ Messager = MsgFmt("")
)
