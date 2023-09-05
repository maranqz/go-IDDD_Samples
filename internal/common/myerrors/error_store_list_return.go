package myerrors

func Errs() []error {
	return nil
}

func First(errs []error) error {
	if len(errs) == 0 {
		return nil
	}

	return errs[0]
}

func Append(errs []error, err error) []error {
	if err != nil {
		errs = append(errs, err)
	}

	return errs
}

func Check0(err error) func(errs []error) []error {
	return func(errs []error) []error {
		errs = Append(errs, err)

		return errs
	}
}

func Check[A any](a A, err error) func(errs []error) (A, []error) {
	return func(errs []error) (A, []error) {
		errs = Append(errs, err)

		return a, errs
	}
}

func Check2[A1, A2 any](a1 A1, a2 A2, err error) func(errs []error) (A1, A2, []error) {
	return func(errs []error) (A1, A2, []error) {
		errs = Append(errs, err)

		return a1, a2, errs
	}
}
