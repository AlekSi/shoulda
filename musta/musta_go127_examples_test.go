//go:build go1.27

package musta

import (
	"fmt"
	"mime"
	"strconv"
)

func ExampleMust_NotFail() {
	actual := T(t).NotFail(strconv.Atoi("42"))
	fmt.Println(actual)

	// Output:
	// 42
}

func ExampleMust_NotFail2() {
	mediaType, params := T(t).NotFail2(mime.ParseMediaType("text/html; charset=utf-8"))
	fmt.Printf("%s %v", mediaType, params)

	// Output:
	// text/html map[charset:utf-8]
}

func ExampleMust_NotFail3() {
	value, multibyte, tail := T(t).NotFail3(strconv.UnquoteChar(`\u263a!`, 0))
	fmt.Printf("%c %t %q\n", value, multibyte, tail)

	// Output:
	// ☺ true "!"
}
