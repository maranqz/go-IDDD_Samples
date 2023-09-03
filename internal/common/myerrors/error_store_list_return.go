//go:build list_return

package myerrors

func ErrStore() []error {
	return nil
}

func Filter(errs []error, err error) []error {
	if err != nil {
		errs = append(errs, err)
	}

	return errs
}

func Check[A any](a A, err error) func(errs []error) (A, []error) {
	return func(errs []error) (A, []error) {
		errs = Filter(errs, err)

		return a, errs
	}
}

func Check2[A1, A2 any](a1 A1, a2 A2, err error) func(errs []error) (A1, A2, []error) {
	return func(errs []error) (A1, A2, []error) {
		errs = Filter(errs, err)

		return a1, a2, errs
	}
}
