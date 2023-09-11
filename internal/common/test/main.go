package main

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

var ErrForumID = errors.New("forumID")

type ForumID struct {
	UUIDv4
}

func NewForumID[In Strings](in In) (res ForumID, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("%w: %w", ErrForumID, err)
		}
	}()

	u, err := NewUUIDv4(in)
	if err != nil {
		return ForumID{}, err
	}

	return ForumID{UUIDv4: u}, nil
}

var (
	ErrExtID       = errors.New("extID")
	ErrExtIDEmpty  = fmt.Errorf("%w: empty", ErrExtID)
	ErrExtIDLength = fmt.Errorf("%w: length", ErrStringID)
)

// ExtID is id from external systems as partners or other organizations.
type ExtID string

var mpErrStringIDToExtID = map[error]error{
	ErrStringID:      ErrExtID,
	ErrStringIDEmpty: ErrExtIDEmpty,
}

func NewExtID[In Strings](in In) (ExtID, error) {
	s, err := newStringID(in)
	if err != nil {
		return "", MapError(err, mpErrStringIDToExtID)
	}

	if err := assertLength(s, 0, 128); err != nil {
		return "", fmt.Errorf("%w: %v", ErrExtIDLength, err)
	}

	return ExtID(s), nil
}

var (
	ErrUUID              = errors.New("uuid")
	ErrUUIDEmpty         = fmt.Errorf("%w: empty", ErrUUID)
	ErrUUIDInvalid       = fmt.Errorf("%w: invalid", ErrUUID)
	ErrUUIDInvalidLength = fmt.Errorf("%w: length should be 36", ErrUUID)
)

var ErrUUIDv4Version = fmt.Errorf("%w: v4: version", ErrUUID)

type UUIDv4 struct{ UUID }

func NewUUIDv4[In Strings](in In) (UUIDv4, error) {
	u, err := NewUUID(in)
	if err != nil {
		return UUIDv4{}, err
	}

	if u.Version() != 4 {
		return UUIDv4{}, ErrUUIDv4Version
	}

	return UUIDv4{u}, nil
}

var emptyUUID = [16]byte{}

type UUID struct{ uuid.UUID }

func (u *UUID) IsEmpty() bool {
	return u == nil || u.UUID == emptyUUID
}

var mpErrStringIDToUUID = map[error]error{
	ErrStringID:      ErrUUID,
	ErrStringIDEmpty: ErrUUIDEmpty,
}

func NewUUID[In Strings](in In) (UUID, error) {
	s, err := newStringID(in)
	if err != nil {
		return UUID{}, MapError(err, mpErrStringIDToUUID)
	}

	if err := assertLength(s, 36, 36); err != nil {
		return UUID{}, ErrUUIDInvalidLength
	}

	u, err := uuid.Parse(s)
	if err != nil {
		return UUID{}, ErrUUIDInvalid
	}

	res := UUID{UUID: u}
	if res.IsEmpty() {
		return UUID{}, ErrUUIDEmpty
	}

	return res, nil
}

var (
	ErrStringID      = errors.New("stringID")
	ErrStringIDEmpty = fmt.Errorf("%w: empty", ErrStringID)
)

type Strings interface {
	~[]byte | ~string
}

func newStringID[In Strings](in In) (string, error) {
	res := string(in)

	if res == "" {
		return res, ErrStringIDEmpty
	}

	return res, nil
}

var (
	ErrLength    = errors.New("length")
	ErrLengthMin = fmt.Errorf("%w: min", ErrLength)
	ErrLengthMax = fmt.Errorf("%w: max", ErrLength)
)

type Measurable interface {
	~string | ~[]any | map[any]any | ~chan<- any
}

func assertLength[In Measurable](in In, min, max int) error {
	ln := len(in)
	if ln < min {
		return fmt.Errorf("%w: %d less %d", ErrLengthMin, ln, min)
	} else if max < ln {
		return fmt.Errorf("%w: %d more %d", ErrLengthMax, ln, max)
	}

	return nil
}

func MapError(err error, mp map[error]error) error {
	if res, ok := mp[err]; ok {
		return res
	}

	return err
}
