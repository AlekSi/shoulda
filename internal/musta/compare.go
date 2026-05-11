package musta

import "github.com/AlekSi/shoulda"

// BeDeepEqual checks that actual and expected are deeply equal.
func BeDeepEqual(tb TB, actual, expected any) {
	tb.Helper()

	if !shoulda.BeDeepEqual(tb, actual, expected) {
		tb.FailNow()
	}
}
