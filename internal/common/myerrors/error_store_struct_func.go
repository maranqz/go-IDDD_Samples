//go:build not_work_struct

package myerrors

type errs []error

// ErrStore doesn't work because type of A1 should be set when ErrStore is created, and we cannot use any or other types.
type ErrStore[A1 any] struct {
	errs
}

func (e ErrStore[A1]) Check(err error) error {
	if err != nil {
		e.errs = append(e.errs, err)
	}

	return err
}

func (e ErrStore[A1]) Check1(a1 A1, err error) (A1, error) {
	if err != nil {
		e.errs = append(e.errs, err)
	}

	return a1, err
}
