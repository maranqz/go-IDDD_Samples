//go:build not_work_variadic

package myerrors

import (
	"fmt"
)

func Check[A any](a A, err error, errs []error) (A, []error) {
	if err != nil {
		errs = append(errs, err)
	}

	return a, errs
}

func ExampleListVariadic() {
	var errs []error

	var i int
	// ERROR: multiple-value NewInt(false) (value of type (int, error)) in single-value context
	// see https://go.dev/ref/spec#Calls, part about f(g(parameters_of_g))
	i, errs = Check(NewInt(false), errs)
	fmt.Printf("int(\"%d\"): with error\n", i)

	fmt.Println(errs)

	// Output: compile error
}
