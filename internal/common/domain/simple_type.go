package domain

import (
	"errors"
	"fmt"
)

var (
	errIntID        = errors.New("intID")
	errIntIDEmpty   = fmt.Errorf("%w: empty", errIntID)
	errIntIDInvalid = fmt.Errorf("%w: invalid", errIntID)
)

type ints interface {
	~int | ~int64
}

func newIntID[In ints](in In) (int64, error) {
	res := int64(in)

	if res == 0 {
		return 0, errIntIDEmpty
	}

	if res < 0 {
		return 0, errIntIDInvalid
	}

	return res, nil
}

type UserID int64

func (u UserID) IsValid() bool {
	return u.Validate() == nil
}

func (u UserID) Validate() error {
	if u == 0 {
		return ErrUserIDEmpty
	} else if u < 0 {
		return ErrUserIDInvalid
	}

	return nil
}

var (
	ErrUserID        = errors.New("userID")
	ErrUserIDEmpty   = fmt.Errorf("%w: empty", ErrUserID)
	ErrUserIDInvalid = fmt.Errorf("%w: invalid", ErrUserID)
)

func NewUserID[In ints](in In) (UserID, error) {
	res := UserID(in)

	if res == 0 {
		return 0, ErrUserIDEmpty
	}

	if res < 0 {
		return 0, ErrUserIDInvalid
	}

	return res, nil
}

func NewUserIDPtr[In ints](in *In) (UserID, error) {
	if in == nil {
		return 0, ErrUserIDEmpty
	}

	return NewUserID(*in)
}

var mpErrCommonIntToUserID = map[error]error{
	errIntID:        ErrUserID,
	errIntIDEmpty:   ErrUserIDEmpty,
	errIntIDInvalid: ErrUserIDInvalid,
}

func MapError(err error, mp map[error]error) error {
	if res, ok := mp[err]; ok {
		return res
	}

	return err
}

func NewUserIDFromCommonInt[In ~int | ~int64](in In) (UserID, error) {
	out, err := newIntID(in)
	if err != nil {
		return 0, MapError(err, mpErrCommonIntToUserID)
	}

	return UserID(out), nil
}
