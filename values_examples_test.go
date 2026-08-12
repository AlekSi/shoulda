package shoulda

import "github.com/AlekSi/shoulda/internal"

// t is used by examples.
var t internal.TestTB

func ExampleBeDeepEqual() {
	BeDeepEqual(t, []int{13}, []int64{13})

	// Output:
	// actual is not deep equal to expected:
	// actual: []int{
	//   13,
	// } ([]int)
	// expected: []int64{
	//   13,
	// } ([]int64)
	// diff expected actual
	// --- expected
	// +++ actual
	// @@ -1,3 +1,3 @@
	// -[]int64{
	// +[]int{
	//    13,
	// -} ([]int64)
	// +} ([]int)
	// FAIL
}

func ExampleBeDeepEqualf() {
	BeDeepEqualf(t, []int{13}, []int64{13}, "extra message: %s, %d", "foo", 42)

	// Output:
	// actual is not deep equal to expected:
	// actual: []int{
	//   13,
	// } ([]int)
	// expected: []int64{
	//   13,
	// } ([]int64)
	// diff expected actual
	// --- expected
	// +++ actual
	// @@ -1,3 +1,3 @@
	// -[]int64{
	// +[]int{
	//    13,
	// -} ([]int64)
	// +} ([]int)
	// extra message: foo, 42
	// FAIL
}

// ExampleNotBeDeepEqual demonstrates NotBeDeepEqual.
func ExampleNotBeDeepEqual() {
	NotBeDeepEqual(t, []int{13}, []int{13})

	// Output:
	// actual is deep equal to expected:
	// actual: []int{
	//   13,
	// } ([]int)
	// expected: []int{
	//   13,
	// } ([]int)
	// FAIL
}
