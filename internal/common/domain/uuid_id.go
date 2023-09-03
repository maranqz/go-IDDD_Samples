package domain

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

var (
	ErrUUID              = errors.New("uuid")
	ErrUUIDEmpty         = fmt.Errorf("%w: empty", ErrUUID)
	ErrUUIDInvalid       = fmt.Errorf("%w: invalid", ErrUUID)
	ErrUUIDInvalidLength = fmt.Errorf("%w: max length is 36", ErrUUID)
)

var emptyUUID = [16]byte{}

type UUID struct{ uuid.UUID }

func (u *UUID) IsEmpty() bool {
	return u == nil || u.UUID == emptyUUID
}

// @TODO при инкапсуляции, можно потом убрать ниже лежащий тип и не публиковать нижележащий API
var mpErrStringIDToUUID = map[error]error{
	ErrStringID:      ErrUUID,
	ErrStringIDEmpty: ErrUUIDEmpty,
	errIntIDInvalid:  ErrUserIDInvalid,
}

func NewUUID[In Strings](in In) (res UUID, err error) {
	s, err := NewStringID(in)
	if err != nil {
		return res, MapError(err, mpErrStringIDToUUID)
	}

	if len(s) > 36 {
		return res, ErrUUIDInvalidLength
	}

	u, err := uuid.Parse(s)
	if err != nil {
		return res, ErrUUIDInvalid
	}

	if u == emptyUUID {
		return res, ErrUUIDEmpty
	}

	return UUID{UUID: u}, nil
}
