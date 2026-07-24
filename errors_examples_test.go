package shoulda

import "errors"

func ExampleError() {
	Error(t, nil)

	// Output:
	// actual is nil error
	// FAIL
}

func ExampleErrorf() {
	Errorf(t, nil, "extra message: %s, %d", "foo", 42)

	// Output:
	// actual is nil error
	// extra message: foo, 42
	// FAIL
}

func ExampleErrorIs() {
	ErrorIs(t, errors.New("actual"), errors.New("expected"))

	// Output:
	// actual does not match expected:
	// actual: &errors.errorString{
	//   s: "actual",
	// } (*errors.errorString)
	// expected: &errors.errorString{
	//   s: "expected",
	// } (*errors.errorString)
	// FAIL
}

func ExampleNoError() {
	NoError(t, errors.New("boom"))

	// Output:
	// actual is not nil error, but "boom":
	// actual: &errors.errorString{
	//   s: "boom",
	// } (*errors.errorString)
	// FAIL
}

func ExampleNoErrorf() {
	NoErrorf(t, errors.New("boom"), "extra message: %s, %d", "foo", 42)

	// Output:
	// actual is not nil error, but "boom":
	// actual: &errors.errorString{
	//   s: "boom",
	// } (*errors.errorString)
	// extra message: foo, 42
	// FAIL
}
