package musta

import "github.com/AlekSi/shoulda"

//go:generate go run ../internal/mustagen

// TB is a subset of [testing.TB] that is used by this package.
type TB = shoulda.TB

// NotFail returns a function that checks that err is nil and returns actual.
func NotFail[T any](actual T, err error) func(TB) T {
	return func(tb TB) T {
		tb.Helper()

		NoError(tb, err)
		return actual
	}
}

// NotFail2 returns a function that checks that err is nil and returns (actual1, actual2).
func NotFail2[T1 any, T2 any](actual1 T1, actual2 T2, err error) func(TB) (T1, T2) {
	return func(tb TB) (T1, T2) {
		tb.Helper()

		NoError(tb, err)
		return actual1, actual2
	}
}

// NotFail3 returns a function that checks that err is nil and returns (actual1, actual2, actual3).
func NotFail3[T1 any, T2 any, T3 any](actual1 T1, actual2 T2, actual3 T3, err error) func(TB) (T1, T2, T3) {
	return func(tb TB) (T1, T2, T3) {
		tb.Helper()

		NoError(tb, err)
		return actual1, actual2, actual3
	}
}
