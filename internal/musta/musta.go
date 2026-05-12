// Package musta provides assertions that fail the test immediately.
package musta

import (
	"testing"

	"github.com/AlekSi/shoulda"
	"github.com/AlekSi/shoulda/internal"
)

//go:generate go run gen.go

// TB is a subset of [testing.TB] that is sufficient for assertions.
type TB interface {
	shoulda.TB
	FailNow()
}

// check interfaces
var (
	_ TB = (testing.TB)(nil)
	_ TB = internal.TestTB{}
)
