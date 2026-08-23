package shoulda

func ExampleBeZerof() {
	BeZerof(t, 13, "extra message: %s, %d", "foo", 67)

	// Output:
	// actual is not zero, but:
	// actual: 13 (int)
	// extra message: foo, 67
	// FAIL
}

func ExampleNotBeZerof() {
	NotBeZerof(t, 0, "extra message: %s, %d", "foo", 67)

	// Output:
	// actual is zero
	// extra message: foo, 67
	// FAIL
}
