//go:build list_return

package myerrors

import (
	"fmt"
)

func ExampleListReturn() {
	var errs []error

	var i int
	i, err := NewInt(true)
	errs = Filter(errs, err)
	fmt.Printf("int(\"%d\"): with error\n", i)

	i, err = NewInt(false)
	errs = Filter(errs, err)
	fmt.Printf("int(\"%d\")\n", i)

	var s string
	s, err = NewString(true)
	errs = Filter(errs, err)
	fmt.Printf("string(\"%s\"): with error\n", s)

	s, err = NewString(false)
	errs = Filter(errs, err)
	fmt.Printf("string(\"%s\")\n", s)

	fmt.Println(errs)

	// Output: int("0"): with error
	// int("1")
	// string(""): with error
	// string("string")
	// [`int err` `string err`]
}

func ExampleListCallback() {
	var errs []error

	var i int
	i, errs = Check(NewInt(true))(errs)
	fmt.Printf("int(\"%d\"): with error\n", i)

	i, errs = Check(NewInt(false))(errs)
	fmt.Printf("int(\"%d\")\n", i)

	var s string
	s, errs = Check(NewString(true))(errs)
	fmt.Printf("string(\"%s\"): with error\n", s)

	s, errs = Check(NewString(false))(errs)
	fmt.Printf("string(\"%s\")\n", s)

	fmt.Println(errs)

	// Output: int("0"): with error
	// int("1")
	// string(""): with error
	// string("string")
	// [`int err` `string err`]
}
