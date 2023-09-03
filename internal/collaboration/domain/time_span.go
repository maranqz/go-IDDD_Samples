package domain

import (
	"errors"
	"fmt"
	"time"
)

var (
	ErrTimeSpan               = errors.New("TimeSpan")
	ErrTimeSpanEndsMoreBegins = fmt.Errorf("%w: ends before it begins", ErrTimeSpan)
)

type TimeSpan struct {
	begins time.Time
	ends   time.Time
}

func NewTimeSpan(begin time.Time, ends time.Time) (res TimeSpan, err error) {
	if begin.After(ends) {
		return res, ErrTimeSpanEndsMoreBegins
	}

	return TimeSpan{
		begins: begin,
		ends:   ends,
	}, nil
}

func (t TimeSpan) Begins() time.Time {
	return t.begins
}

func (t TimeSpan) Ends() time.Time {
	return t.ends
}
