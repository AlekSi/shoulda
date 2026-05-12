package musta

import (
	"github.com/AlekSi/shoulda"
)

// BeNil checks that actual is (untyped) nil.
func BeNil(tb TB, actual any) {
	tb.Helper()

	if !shoulda.BeNil(tb, actual) {
		tb.FailNow()
	}
}

// BeDeepEqual checks that actual and expected are deeply equal.
func BeDeepEqual(tb TB, actual, expected any) {
	tb.Helper()

	if !shoulda.BeDeepEqual(tb, actual, expected) {
		tb.FailNow()
	}
}
