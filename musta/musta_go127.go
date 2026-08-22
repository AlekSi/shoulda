//go:build go1.27

package musta

// Must provides assertions bound to a [TB].
type Must struct {
	tb TB
}

// With returns assertions bound to tb.
func With(tb TB) Must {
	return Must{tb: tb}
}

// NotFail checks that err is nil and returns actual.
func (m Must) NotFail[T any](actual T, err error) T {
	m.tb.Helper()

	NoError(m.tb, err)
	return actual
}

// NotFail2 checks that err is nil and returns (actual1, actual2).
func (m Must) NotFail2[T1 any, T2 any](actual1 T1, actual2 T2, err error) (T1, T2) {
	m.tb.Helper()

	NoError(m.tb, err)
	return actual1, actual2
}

// NotFail3 checks that err is nil and returns (actual1, actual2, actual3).
func (m Must) NotFail3[T1 any, T2 any, T3 any](actual1 T1, actual2 T2, actual3 T3, err error) (T1, T2, T3) {
	m.tb.Helper()

	NoError(m.tb, err)
	return actual1, actual2, actual3
}
