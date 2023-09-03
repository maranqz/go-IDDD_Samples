//go:build not_work_struct

package myerrors

import (
	"fmt"
)

func ExampleStructFunc() {
	errs := ErrStore[any]{}

	var i int

	i, errInt := errs.Check1(NewInt(true))
	fmt.Println(i, errInt)

	var s string

	s, errString := errs.Check1(NewString(true))
	fmt.Println(s, errString)

	fmt.Println(errs)

	// Output:
}
