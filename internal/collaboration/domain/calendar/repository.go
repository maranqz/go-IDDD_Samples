package calendar

import "github.com/maranqz/go-IDDD_Samples/internal/collaboration/domain/tenant"

type CalendarRepository interface {
	CalendarOfId(tenant tenant.ID, calendarId ID) (*Calendar, error)
	NextIdentity() ID
	Save(calendar *Calendar) error
}

type CalendarEntryRepository interface {
	CalendarEntryOfId(tenant tenant.ID, calendarEntryID EntryID) (*CalendarEntry, error)
	NextIdentity() EntryID
	Save(calendarEntry *CalendarEntry) error
}
