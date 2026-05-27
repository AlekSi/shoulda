package shoulda

import "errors"

func ExampleBeNil() {
	BeNil(t, new(uint32(13)))

	// Output:
	// actual is not untyped nil, but &13 (*uint32)
	// FAIL
}

func ExampleBeNil_typedNil() {
	BeNil(t, (*uint32)(nil))

	// Output:
	// actual is not untyped nil, but nil (*uint32)
	// FAIL
}

func ExampleNotBeNil() {
	NotBeNil(t, nil)

	// Output:
	// actual is untyped nil
	// FAIL
}

func ExampleBeZero() {
	BeZero(t, 13)

	// Output:
	// actual is not zero, but 13 (int)
	// FAIL
}

func ExampleNotBeZero() {
	NotBeZero(t, 0)

	// Output:
	// actual is zero
	// FAIL
}

func ExampleNotBeZero_pointer() {
	NotBeZero(t, (*int)(nil))

	// Output:
	// actual is zero
	// FAIL
}

func ExampleError() {
	Error(t, nil)

	// Output:
	// actual is nil error
	// FAIL
}

func ExampleNoError() {
	NoError(t, errors.New("boom"))

	// Output:
	// actual is not nil error, but "boom"
	// &errors.errorString{
	//   s: "boom",
	// } (*errors.errorString)
	// FAIL
}
