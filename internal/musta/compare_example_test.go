package musta

import (
	"github.com/AlekSi/shoulda/internal"
)

// t is used by examples.
var t internal.TestTB

func ExampleBeDeepEqual() {
	actual := []int{13}
	expected := []int{42}
	BeDeepEqual(t, actual, expected)

	// Output:
	// Values are not deep equal:
	// actual:   []int{13}
	// expected: []int{42}
	// FAIL
}
