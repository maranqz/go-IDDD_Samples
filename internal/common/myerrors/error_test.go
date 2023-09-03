package myerrors

import (
	"errors"
)

func NewString(withError bool) (string, error) {
	if withError {
		return "", errors.New("`string err`")
	}

	return "string", nil
}

func NewInt(withError bool) (int, error) {
	if withError {
		return 0, errors.New("`int err`")
	}

	return 1, nil
}
