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
	ErrorIs(t, errors.New("boom"), errors.New("target"))

	// Output:
	// actual error does not match expected:
	// actual: boom
	// &errors.errorString{
	//   s: "boom",
	// } (*errors.errorString)
	// expected: target
	// &errors.errorString{
	//   s: "target",
	// } (*errors.errorString)
	// FAIL
}

func ExampleErrorIsf() {
	ErrorIsf(t, errors.New("boom"), errors.New("target"), "extra message: %s, %d", "foo", 42)

	// Output:
	// actual error does not match expected:
	// actual: boom
	// &errors.errorString{
	//   s: "boom",
	// } (*errors.errorString)
	// expected: target
	// &errors.errorString{
	//   s: "target",
	// } (*errors.errorString)
	// extra message: foo, 42
	// FAIL
}

func ExampleNoError() {
	NoError(t, errors.New("boom"))

	// Output:
	// actual is not nil error:
	// actual: boom
	// &errors.errorString{
	//   s: "boom",
	// } (*errors.errorString)
	// FAIL
}

func ExampleNoErrorf() {
	NoErrorf(t, errors.New("boom"), "extra message: %s, %d", "foo", 42)

	// Output:
	// actual is not nil error:
	// actual: boom
	// &errors.errorString{
	//   s: "boom",
	// } (*errors.errorString)
	// extra message: foo, 42
	// FAIL
}
