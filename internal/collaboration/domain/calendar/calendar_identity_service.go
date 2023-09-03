package calendar

type CalendarIdentityService struct {
	calendarRepository      CalendarRepository
	calendarEntryRepository CalendarEntryRepository
}

func NewCalendarIdentityService(
	calendarRepo CalendarRepository,
	calendarEntryRepo CalendarEntryRepository,
) (*CalendarIdentityService, error) {
	return &CalendarIdentityService{
		calendarRepository:      calendarRepo,
		calendarEntryRepository: calendarEntryRepo,
	}, nil
}

func (c *CalendarIdentityService) NextCalendarEntryId() EntryID {
	return c.calendarEntryRepository.NextIdentity()
}

func (c *CalendarIdentityService) NextCalendarId() ID {
	return c.calendarRepository.NextIdentity()
}
