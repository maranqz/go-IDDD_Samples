package domain

type CalendarSharer struct {
	participant *Participant
}

func NewCalendarSharer(aParticipant *Participant) *CalendarSharer {
	return &CalendarSharer{
		participant: aParticipant,
	}
}

func (cs *CalendarSharer) Participant() *Participant {
	return cs.participant
}
