package domain

import (
	"errors"
	"fmt"

	"github.com/maranqz/go-IDDD_Samples/internal/common/domain"
)

var (
	ErrCalendar           = errors.New("calendar")
	ErrCalendarName       = fmt.Errorf("%w: name", ErrCalendar)
	ErrCalendarNameEmpty  = fmt.Errorf("%w: empty", ErrCalendarName)
	ErrCalendarNameMaxLen = fmt.Errorf("%w: max length 255", ErrCalendarName)
	ErrCalendarDescEmpty  = fmt.Errorf("%w: description: empty", ErrCalendar)
)

type invitees map[*Participant]struct{}

// Bad idea, CalendarID loses methods, see https://go.dev/ref/spec#Type_definitions
// type CalendarID domain.UUID

type CalendarID struct {
	domain.UUID
}

func NewCalendarID[In domain.Strings](in In) (CalendarID, error) {
	u, err := domain.NewUUID(in)

	return CalendarID{UUID: u}, err
}

type Name string

var mpErrTextToName = map[error]error{
	domain.ErrTextEmpty: ErrCalendarNameEmpty,
}

func NewName[In string](in In) (Name, error) {
	t, err := domain.NewText(in)
	if err != nil {
		return "", domain.MapError(err, mpErrTextToName)
	}

	if len(t) > 255 {
		return "", ErrCalendarNameMaxLen
	}

	return Name(in), nil
}

type Calendar struct {
	calendarId  CalendarID
	description string
	name        Name
	owner       *Owner
	sharedWith  map[CalendarSharer]struct{}
	tenant      TenantID
}

func NewCalendar(
	aTenant TenantID,
	aCalendarId CalendarID,
	aName Name,
	aDescription string,
) (*Calendar, error) {

	if aDescription == "" {
		return nil, ErrCalendarDescEmpty
	}

	return &Calendar{
		calendarId:  aCalendarId,
		description: aDescription,
		name:        aName,
		sharedWith:  make(map[CalendarSharer]struct{}),
		tenant:      aTenant,
	}, nil
}

func (c *Calendar) AllSharedWith() map[CalendarSharer]struct{} {
	return c.sharedWith
}

func (c *Calendar) CalendarId() CalendarID {
	return c.calendarId
}

// ChangeDescription has duplication in NewCalendar because the description is not distinguished type.
func (c *Calendar) ChangeDescription(d string) error {
	if d == "" {
		return ErrCalendarDescEmpty
	}

	return nil
}

func (c *Calendar) Description() string {
	return c.description
}

// Rename can be deleted because Name is already valid,
// and we don't have another check inside the method.
func (c *Calendar) Rename(n Name) {
	c.name = n
}

func (c *Calendar) Name() Name {
	return c.name
}

func (c *Calendar) Owner() *Owner {
	return c.owner
}

func (c *Calendar) AddSharedWith(sharer CalendarSharer) {
	c.sharedWith[sharer] = struct{}{}
}

func (c *Calendar) RemoveSharedWith(sharer CalendarSharer) {
	delete(c.sharedWith, sharer)
}

func (c *Calendar) SharedWith() []CalendarSharer {
	sharedWith := make([]CalendarSharer, 0, len(c.sharedWith))
	for sharer := range c.sharedWith {
		sharedWith = append(sharedWith, sharer)
	}

	return sharedWith
}

func (c *Calendar) TenantID() TenantID {
	return c.tenant
}

func (c *Calendar) ScheduleCalendarEntry(
	aCalendarIdentityService CalendarIdentityService,
	aDescription string,
	aLocation string,
	anOwner *Owner,
	aTimeSpan TimeSpan,
	aRepetition Repetition,
	anAlarm Alarm,
	anInvitees invitees) (*CalendarEntry, error) {
	return NewCalendarEntry(
		c.tenant,
		c.calendarId,
		aCalendarIdentityService.NextCalendarEntryId(),
		aDescription,
		aLocation,
		anOwner,
		aTimeSpan,
		aRepetition,
		anAlarm,
		anInvitees,
	)
}
