package calendar

import (
	"errors"
	"fmt"
	"strings"

	"github.com/maranqz/go-IDDD_Samples/internal/collaboration/domain/collaborator"
	"github.com/maranqz/go-IDDD_Samples/internal/collaboration/domain/tenant"
	"github.com/maranqz/go-IDDD_Samples/internal/common/domain"
	"github.com/maranqz/go-IDDD_Samples/internal/common/myerrors"
)

var (
	ErrCalendarEntry                 = errors.New("calendarEntry")
	ErrCalendarEntryDescEmpty        = fmt.Errorf("%w: description: empty", ErrCalendarEntry)
	ErrCalendarEntryLocationEmpty    = fmt.Errorf("%w: location: empty", ErrCalendarEntry)
	ErrCalendarEntryParticipantEmpty = fmt.Errorf("%w: participant: empty", ErrCalendarEntry)
)

type EntryID struct {
	domain.UUID
}

func NewCalendarEntryID[In domain.Strings](in In) (ID, error) {
	u, err := domain.NewUUID(in)

	return ID{UUID: u}, err
}

type Invitees = map[*collaborator.Participant]struct{}

type CalendarEntry struct {
	alarm           Alarm
	calendarEntryId EntryID
	calendarID      ID
	description     string
	invitees        Invitees
	location        string
	owner           *collaborator.Owner
	repetition      Repetition
	tenant          tenant.ID
	timeSpan        TimeSpan
}

func NewCalendarEntry(
	aTenant tenant.ID,
	aCalendarId ID,
	aCalendarEntryId EntryID,
	aDescription string,
	aLocation string,
	anOwner *collaborator.Owner,
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
	if err != nil {
		return nil, err
	}

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

// AllInvitees returns copy of Invitees.
// Other solutions to get around the lack of immutable types:
// * linter to block changes, golang doesn't have immutable map
// * use lib https://github.com/benbjohnson/immutable
// proposal https://github.com/golang/go/issues/27975
func (c *CalendarEntry) AllInvitees() Invitees {
	res := make(Invitees, len(c.invitees))

	for k, v := range c.invitees {
		res[k] = v
	}

	return res
}

func (c *CalendarEntry) CalendarEntryID() EntryID {
	return c.calendarEntryId
}

func (c *CalendarEntry) CalendarID() ID {
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

func (c *CalendarEntry) assertInvitee(aParticipant *collaborator.Participant) error {
	if aParticipant.IsEmpty() {
		return ErrCalendarEntryParticipantEmpty
	}

	return nil
}

func (c *CalendarEntry) Invite(aParticipant *collaborator.Participant) error {
	if err := c.assertInvitee(aParticipant); err != nil {
		return err
	}

	if _, ok := c.invitees[aParticipant]; !ok {
		c.invitees[aParticipant] = struct{}{}
	}

	return nil
}

// TODO add error wrapping
func (c *CalendarEntry) Uninvite(aParticipant *collaborator.Participant) error {
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

func (c *CalendarEntry) Owner() *collaborator.Owner {
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
	aRepetition, errs := myerrors.Check(checkRepetitionDoesNotRepeat(aRepetition, aTimeSpan))(nil)
	errs = myerrors.Check0(assertTimeSpans(aRepetition, aTimeSpan))(errs)
	errs = myerrors.Check0(c.ChangeDescription(aDescription))(errs)
	errs = myerrors.Check0(c.Relocate(aLocation))(errs)

	if err := myerrors.First(errs); err != nil {
		return err
	}

	c.alarm = anAlarm
	c.repetition = aRepetition
	c.timeSpan = aTimeSpan

	return nil
}

func (c *CalendarEntry) Tenant() tenant.ID {
	return c.tenant
}

func (c *CalendarEntry) TimeSpan() TimeSpan {
	return c.timeSpan
}
