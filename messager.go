package shoulda

import (
	"fmt"
	"strings"
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

// msgf constructs a [messager] from a message or format string, and arguments.
func msgf(format string, args ...any) messager {
	if len(args) == 0 {
		return msgString(strings.TrimRight(format, "\n"))
	}

	return msgFunc(func() string {
		return strings.TrimRight(fmt.Sprintf(format, args...), "\n")
	})
}
