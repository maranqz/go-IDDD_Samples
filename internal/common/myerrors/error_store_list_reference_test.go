//go:build list_reference

package myerrors

import (
	"fmt"
)

func ExampleListReference() {
	var e []error
	errs := &e // or errs := ErrStore()

	var i int
	i = Check(NewInt(true))(errs)
	fmt.Printf("int(\"%d\"): with error\n", i)

	i = Check(NewInt(false))(errs)
	fmt.Printf("int(\"%d\")\n", i)

	var s string
	s = Check(NewString(true))(errs)
	fmt.Printf("string(\"%s\"): with error\n", s)

	s = Check(NewString(false))(errs)
	fmt.Printf("string(\"%s\")\n", s)

	fmt.Println(errs)

	// Output: int("0"): with error
	// int("1")
	// string(""): with error
	// string("string")
	// &[`int err` `string err`]
}
