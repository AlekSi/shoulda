//go:build go1.27

package musta

// Must provides assertions bound to a [TB].
type Must struct {
	tb TB
}

// T returns assertions bound to tb.
func T(tb TB) Must {
	return Must{tb: tb}
}

// NotFail checks that err is nil and returns actual.
func (m Must) NotFail[T any](actual T, err error) T {
	m.tb.Helper()

	return NotFail(actual, err)(m.tb)
}

// NotFail2 checks that err is nil and returns (actual1, actual2).
func (m Must) NotFail2[T1 any, T2 any](actual1 T1, actual2 T2, err error) (T1, T2) {
	m.tb.Helper()

	return NotFail2(actual1, actual2, err)(m.tb)
}

// NotFail3 checks that err is nil and returns (actual1, actual2, actual3).
func (m Must) NotFail3[T1 any, T2 any, T3 any](actual1 T1, actual2 T2, actual3 T3, err error) (T1, T2, T3) {
	m.tb.Helper()

	return NotFail3(actual1, actual2, actual3, err)(m.tb)
}
