// This file is based on Go's cmp standard library package:
// https://go.googlesource.com/go/+/refs/heads/master/src/cmp/cmp.go
//
// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package cmp provides types and functions related to comparing ordered values.
//
// It is a copy of [cmp] with some additions.
// It is a strict superset and can be used as a drop-in replacement for the standard library package everywhere.
package cmp

import (
	"cmp"
)

// Ordered is a constraint that permits any ordered type: any type
// that supports the operators < <= >= >.
// If future releases of Go add new ordered types,
// this constraint will be modified to include them.
//
// Note that floating-point types may contain NaN ("not-a-number") values.
// An operator such as == or < will always report false when
// comparing a NaN value with any other value, NaN or not.
// See the [Compare] function for a consistent way to compare NaN values.
//
// It is a type alias for [cmp.Ordered].
type Ordered = cmp.Ordered

// Equal reports whether x is equal to y.
// For floating-point types, a NaN is considered equal to a NaN,
// and -0.0 is equal to 0.0.
func Equal[T Ordered](x, y T) bool {
	// It is tempting to use comparable instead of Ordered,
	// but see https://github.com/golang/go/issues/70161

	return (isNaN(x) && isNaN(y)) || x == y
}

// Less reports whether x is less than y.
// For floating-point types, a NaN is considered less than any non-NaN,
// and -0.0 is not less than (is equal to) 0.0.
//
// It is the same as [cmp.Less].
func Less[T Ordered](x, y T) bool {
	return (isNaN(x) && !isNaN(y)) || x < y
}

// Greater reports whether x is greater than y.
// For floating-point types, any non-NaN is considered greater than a NaN,
// and -0.0 is not greater than (is equal to) 0.0.
func Greater[T Ordered](x, y T) bool {
	return (!isNaN(x) && isNaN(y)) || x > y
}

// Compare returns
//
//	-1 (int([OrderLess])) if x is less than y,
//	 0 (int([OrderEqual])) if x equals y,
//	+1 (int([OrderGreater])) if x is greater than y.
//
// For floating-point types, a NaN is considered less than any non-NaN,
// a NaN is considered equal to a NaN, and -0.0 is equal to 0.0.
//
// It is the same as [cmp.Compare].
func Compare[T Ordered](x, y T) int {
	xNaN := isNaN(x)
	yNaN := isNaN(y)
	if xNaN {
		if yNaN {
			return 0
		}
		return -1
	}
	if yNaN {
		return +1
	}
	if x < y {
		return -1
	}
	if x > y {
		return +1
	}
	return 0
}

// isNaN reports whether x is a NaN without requiring the math package.
// This will always return false if T is not floating-point.
//
// It is the same as [cmp.isNaN].
func isNaN[T Ordered](x T) bool {
	return x != x
}

// Or returns the first of its arguments that is not equal to the zero value.
// If no argument is non-zero, it returns the zero value.
//
// It is the same as [cmp.Or].
func Or[T comparable](vals ...T) T {
	var zero T
	for _, val := range vals {
		if val != zero {
			return val
		}
	}
	return zero
}
