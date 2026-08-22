package shoulda

func ExampleBeNilf() {
	BeNilf(t, new(uint32(13)), "extra message: %s, %d", "foo", 42)

	// Output:
	// actual is not untyped nil, but:
	// actual: &13 (*uint32)
	// extra message: foo, 42
	// FAIL
}

func ExampleNotBeNilf() {
	NotBeNilf(t, nil, "extra message: %s, %d", "foo", 42)

	// Output:
	// actual is untyped nil
	// extra message: foo, 42
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

func ExampleNotBeZerof() {
	NotBeZerof(t, 0, "extra message: %s, %d", "foo", 42)

	// Output:
	// actual is zero
	// extra message: foo, 42
	// FAIL
}
