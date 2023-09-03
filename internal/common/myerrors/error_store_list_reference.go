//go:build list_reference

package myerrors

func ErrStore() *[]error {
	var e []error

	return &e
}

func Check[A any](a A, err error) func(e *[]error) A {
	return func(e *[]error) A {
		if err != nil {
			*e = append(*e, err)
		}

		return a
	}
}

func Check2[A1, A2 any](a1 A1, a2 A2, err error) func(e *[]error) (A1, A2) {
	return func(e *[]error) (A1, A2) {
		if err != nil {
			*e = append(*e, err)
		}

		return a1, a2
	}
}
