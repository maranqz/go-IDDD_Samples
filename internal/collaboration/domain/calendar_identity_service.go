package domain

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

func (c *CalendarIdentityService) NextCalendarEntryId() CalendarEntryID {
	return c.calendarEntryRepository.NextIdentity()
}

func (c *CalendarIdentityService) NextCalendarId() CalendarID {
	return c.calendarRepository.NextIdentity()
}
