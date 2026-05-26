package shoulda

import "errors"

func ExampleBeNil_pointer() {
	BeNil(t, new(uint32(13)))

	// Output:
	// actual is not untyped nil, but *uint32:
	// &13
	// FAIL
}

func ExampleBeNil_typedNil() {
	BeNil(t, (*uint32)(nil))

	// Output:
	// actual is not untyped nil, but *uint32:
	// nil
	// FAIL
}

func ExampleBeNil_error() {
	BeNil(t, errors.New("boom"))

	// Output:
	// actual is not untyped nil, but *errors.errorString:
	// &errors.errorString{
	//   s: "boom",
	// }
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
	// actual is not zero, but int:
	// 13
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
