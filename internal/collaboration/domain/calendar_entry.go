package domain

import (
	"errors"
	"fmt"
	"strings"

	"github.com/maranqz/go-IDDD_Samples/internal/common/domain"
)

var (
	ErrCalendarEntry                 = errors.New("calendarEntry")
	ErrCalendarEntryDescEmpty        = fmt.Errorf("%w: description: empty", ErrCalendarEntry)
	ErrCalendarEntryLocationEmpty    = fmt.Errorf("%w: location: empty", ErrCalendarEntry)
	ErrCalendarEntryParticipantEmpty = fmt.Errorf("%w: participant: empty", ErrCalendarEntry)
)

type CalendarEntryID struct {
	domain.UUID
}

func NewCalendarEntryID[In domain.Strings](in In) (CalendarID, error) {
	u, err := domain.NewUUID(in)

	return CalendarID{UUID: u}, err
}

type Invitees = map[*Participant]struct{}

type CalendarEntry struct {
	alarm           Alarm
	calendarEntryId CalendarEntryID
	calendarID      CalendarID
	description     string
	invitees        Invitees
	location        string
	owner           *Owner
	repetition      Repetition
	tenant          TenantID
	timeSpan        TimeSpan
}

func NewCalendarEntry(
	aTenant TenantID,
	aCalendarId CalendarID,
	aCalendarEntryId CalendarEntryID,
	aDescription string,
	aLocation string,
	anOwner *Owner,
	aTimeSpan TimeSpan,
	aRepetition Repetition,
	anAlarm Alarm,
	anInvitees invitees,
) (res *CalendarEntry, err error) {
	if aDescription == "" {
		return nil, ErrCalendarEntryDescEmpty
	}

	if aLocation == "" {
		return nil, ErrCalendarEntryLocationEmpty
	}

	aRepetition, err = checkRepetitionDoesNotRepeat(aRepetition, aTimeSpan)

	if err = assertTimeSpans(aRepetition, aTimeSpan); err != nil {
		return nil, err
	}

	if anInvitees == nil {
		anInvitees = make(Invitees)
	}

	return &CalendarEntry{
		alarm:           anAlarm,
		calendarEntryId: aCalendarEntryId,
		calendarID:      aCalendarId,
		description:     aDescription,
		invitees:        anInvitees,
		location:        aLocation,
		owner:           anOwner,
		repetition:      aRepetition,
		tenant:          aTenant,
		timeSpan:        aTimeSpan,
	}, nil
}

func checkRepetitionDoesNotRepeat(repetition Repetition, timeSpan TimeSpan) (res Repetition, err error) {
	if repetition.Repeats() == DoesNotRepeat {
		repetition, err = DoesNotRepetition(timeSpan.Ends())
		if err != nil {
			return res, err
		}
	}

	return repetition, nil
}

var (
	ErrCalendarTimeSpanDoesNotRepeat      = fmt.Errorf("%w: non-repeating entry must end with time span end", ErrCalendarEntry)
	ErrCalendarTimeSpanAfterRepetitionEnd = fmt.Errorf("%w: timeSpan must end when or before repetition ends", ErrCalendarEntry)
)

func assertTimeSpans(repetition Repetition, timeSpan TimeSpan) error {
	if repetition.Repeats() == DoesNotRepeat {
		if repetition.Ends() != timeSpan.Ends() {
			return ErrCalendarTimeSpanDoesNotRepeat
		}
	} else if timeSpan.Ends().After(repetition.Ends()) {
		return ErrCalendarTimeSpanAfterRepetitionEnd
	}

	return nil
}

func (c *CalendarEntry) Alarm() Alarm {
	return c.alarm
}

// TODO linter to block changes, golang doesn't have immutable map
// or use https://github.com/benbjohnson/immutable
// https://github.com/golang/go/issues/27975
func (c *CalendarEntry) AllInvitees() Invitees {
	return c.invitees
}

func (c *CalendarEntry) CalendarEntryID() CalendarEntryID {
	return c.calendarEntryId
}

func (c *CalendarEntry) CalendarID() CalendarID {
	return c.calendarID
}

func (c *CalendarEntry) ChangeDescription(aDescription string) error {
	if aDescription == "" {
		return nil
	}

	c.description = aDescription

	return nil
}

func (c *CalendarEntry) Description() string {
	return c.description
}

func (c *CalendarEntry) assertInvitee(aParticipant *Participant) error {
	if aParticipant.IsEmpty() {
		return ErrCalendarEntryParticipantEmpty
	}

	return nil
}

func (c *CalendarEntry) Invite(aParticipant *Participant) error {
	if err := c.assertInvitee(aParticipant); err != nil {
		return err
	}

	if _, ok := c.invitees[aParticipant]; !ok {
		c.invitees[aParticipant] = struct{}{}
	}

	return nil
}

// TODO add error wrapping
func (c *CalendarEntry) Uninvite(aParticipant *Participant) error {
	if err := c.assertInvitee(aParticipant); err != nil {
		return err
	}

	if _, ok := c.invitees[aParticipant]; !ok {
		delete(c.invitees, aParticipant)
	}

	return nil
}

func (c *CalendarEntry) Location() string {
	return c.location
}

func (c *CalendarEntry) Owner() *Owner {
	return c.owner
}

func (c *CalendarEntry) Relocate(aLocation string) error {
	aLocation = strings.Trim(aLocation, "")

	if aLocation == "" {
		return nil
	}

	c.location = aLocation

	return nil
}

func (c *CalendarEntry) Repetition() Repetition {
	return c.repetition
}

func (c *CalendarEntry) Reschedule(aDescription string, aLocation string, aTimeSpan TimeSpan, aRepetition Repetition, anAlarm Alarm) (err error) {
	aRepetition, err = checkRepetitionDoesNotRepeat(aRepetition, aTimeSpan)

	// TODO Add on application level
	id, err := domain.NewUUID(inId)
	if err != nil {
		return nil, err
	}
	id, err := domain.NewUUID(inId)
	if err != nil {
		return nil, err
	}
	id, err := domain.NewUUID(inId)
	if err != nil {
		return nil, err
	}
	id, err := domain.NewUUID(inId)
	if err != nil {
		return nil, err
	}

	if err = assertTimeSpans(aRepetition, aTimeSpan); err != nil {
		return err
	} else if err = c.ChangeDescription(aDescription); err != nil {
		return err
	} else if err = c.Relocate(aLocation); err != nil {
		return err
	}

	c.alarm = anAlarm
	c.repetition = aRepetition
	c.timeSpan = aTimeSpan

	return nil
}

func (c *CalendarEntry) Tenant() TenantID {
	return c.tenant
}

func (c *CalendarEntry) TimeSpan() TimeSpan {
	return c.timeSpan
}
