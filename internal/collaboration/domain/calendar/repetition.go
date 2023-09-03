package calendar

import (
	"errors"
	"fmt"
	"time"
)

var (
	ErrRepeatType        = errors.New("repeatType")
	ErrRepeatTypeInvalid = fmt.Errorf("%w: invalid", ErrRepeatType)
)

type RepeatType int

// +1 to move default value and for var RepeatType be incorrect value.
const (
	DoesNotRepeat RepeatType = iota + 1
	Daily
	Weekly
	Monthy
	Yearly
)

var mpRepeatType = map[RepeatType]struct{}{
	DoesNotRepeat: {},
	Daily:         {},
	Weekly:        {},
	Monthy:        {},
	Yearly:        {},
}

func NewRepeatType[In ~int](in In) (RepeatType, error) {
	r := RepeatType(in)

	if _, ok := mpRepeatType[r]; !ok {
		return 0, ErrRepeatTypeInvalid
	}

	return r, nil
}

type Repetition struct {
	ends    time.Time
	repeats RepeatType
}

func NewRepetition(
	repeats RepeatType,
	ends time.Time,
) (Repetition, error) {
	return Repetition{ends: ends, repeats: repeats}, nil
}

func DoesNotRepetition(anEnds time.Time) (Repetition, error) {
	return NewRepetition(DoesNotRepeat, anEnds)
}

func (r Repetition) Ends() time.Time {
	return r.ends
}

func (r Repetition) Repeats() RepeatType {
	return r.repeats
}
