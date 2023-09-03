package domain

type CalendarRepository interface {
	CalendarOfId(tenant TenantID, calendarId CalendarID) (*Calendar, error)
	NextIdentity() CalendarID
	Save(calendar *Calendar) error
}

type CalendarEntryRepository interface {
	CalendarEntryOfId(tenant TenantID, calendarEntryID CalendarEntryID) (*CalendarEntry, error)
	NextIdentity() CalendarEntryID
	Save(calendarEntry *CalendarEntry) error
}
