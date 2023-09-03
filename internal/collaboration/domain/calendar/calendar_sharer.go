package calendar

import "github.com/maranqz/go-IDDD_Samples/internal/collaboration/domain/collaborator"

type CalendarSharer struct {
	participant *collaborator.Participant
}

func NewCalendarSharer(aParticipant *collaborator.Participant) *CalendarSharer {
	return &CalendarSharer{
		participant: aParticipant,
	}
}

func (cs *CalendarSharer) Participant() *collaborator.Participant {
	return cs.participant
}
