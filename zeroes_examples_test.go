package shoulda

import "errors"

func ExampleBeNil() {
	BeNil(t, new(uint32(13)))

	// Output:
	// actual is not untyped nil, but:
	// actual: &13 (*uint32)
	// FAIL
}

func ExampleBeNil_typedNil() {
	BeNil(t, (*uint32)(nil))

	// Output:
	// actual is not untyped nil, but:
	// actual: nil (*uint32)
	// FAIL
}

func ExampleBeNilf() {
	BeNilf(t, new(uint32(13)), "extra message: %s, %d", "foo", 42)

	// Output:
	// actual is not untyped nil, but:
	// actual: &13 (*uint32)
	// extra message: foo, 42
	// FAIL
}

func ExampleNotBeNil() {
	NotBeNil(t, nil)

	// Output:
	// actual is untyped nil
	// FAIL
}

func ExampleNotBeNilf() {
	NotBeNilf(t, nil, "extra message: %s, %d", "foo", 42)

	// Output:
	// actual is untyped nil
	// extra message: foo, 42
	// FAIL
}

func ExampleBeZero() {
	BeZero(t, 13)

	// Output:
	// actual is not zero, but:
	// actual: 13 (int)
	// FAIL
}

func ExampleBeZerof() {
	BeZerof(t, 13, "extra message: %s, %d", "foo", 42)

	// Output:
	// actual is not zero, but:
	// actual: 13 (int)
	// extra message: foo, 42
	// FAIL
}

func ExampleNotBeZero() {
	NotBeZero(t, 0)

	// Output:
	// actual is zero
	// FAIL
}

func ExampleNotBeZerof() {
	NotBeZerof(t, 0, "extra message: %s, %d", "foo", 42)

	// Output:
	// actual is zero
	// extra message: foo, 42
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

func ExampleErrorf() {
	Errorf(t, nil, "extra message: %s, %d", "foo", 42)

	// Output:
	// actual is nil error
	// extra message: foo, 42
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
