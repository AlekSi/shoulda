package musta

import (
	"fmt"
	"mime"
	"strconv"
)

func ExampleNotFail() {
	actual := NotFail(strconv.Atoi("42"))(t)
	fmt.Println(actual)

	// Output:
	// 42
}

func ExampleNotFail2() {
	mediaType, params := NotFail2(mime.ParseMediaType("text/html; charset=utf-8"))(t)
	fmt.Printf("%s %v", mediaType, params)

	// Output:
	// text/html map[charset:utf-8]
}

func ExampleNotFail3() {
	value, multibyte, tail := NotFail3(strconv.UnquoteChar(`\u263a!`, 0))(t)
	fmt.Printf("%c %t %q\n", value, multibyte, tail)

	// Output:
	// ☺ true "!"
}
