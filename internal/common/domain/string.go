package domain

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrStringID      = errors.New("stringID")
	ErrStringIDEmpty = fmt.Errorf("%w: empty", ErrStringID)
)

type Strings interface {
	~[]byte | ~string
}

func NewStringID[In Strings](in In) (string, error) {
	res := string(in)

	if res == "" {
		return res, ErrStringIDEmpty
	}

	return res, nil
}

var (
	ErrText      = errors.New("text")
	ErrTextEmpty = fmt.Errorf("%w: empty", ErrText)
)

type Text string

func NewText[In Strings](in In) (Text, error) {
	res := Text(strings.Trim(
		string(in),
		" ",
	))

	if res == "" {
		return res, ErrTextEmpty
	}

	return res, nil
}
