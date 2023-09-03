package forum

import (
	"time"

	"github.com/maranqz/go-IDDD_Samples/internal/collaboration/domain/tenant"
)

type ForumStarted struct {
	creator        *Creator
	description    string
	eventVersion   int
	exclusiveOwner string
	forumId        ForumID
	moderator      *Moderator
	occurredOn     time.Time
	subject        string
	tenant         tenant.ID
}

func NewForumStarted(aTenant tenant.ID, aForumId ForumID, aCreator *Creator, aModerator *Moderator, aSubject string, aDescription string, anExclusiveOwner string) (rcvr *ForumStarted) {
	rcvr = &ForumStarted{}
	rcvr.creator = aCreator
	rcvr.description = aDescription
	rcvr.eventVersion = 1
	rcvr.exclusiveOwner = anExclusiveOwner
	rcvr.forumId = aForumId
	rcvr.moderator = aModerator
	rcvr.occurredOn = NewDate()
	rcvr.subject = aSubject
	rcvr.tenant = aTenant
	return
}

func (rcvr *ForumStarted) Creator() *Creator {
	return rcvr.creator
}

func (rcvr *ForumStarted) Description() string {
	return rcvr.description
}

func (rcvr *ForumStarted) EventVersion() int {
	return rcvr.eventVersion
}

func (rcvr *ForumStarted) ExclusiveOwner() string {
	return rcvr.exclusiveOwner
}

func (rcvr *ForumStarted) ForumId() ForumID {
	return rcvr.forumId
}

func (rcvr *ForumStarted) Moderator() *Moderator {
	return rcvr.moderator
}

func (rcvr *ForumStarted) OccurredOn() time.Time {
	return rcvr.occurredOn
}

func (rcvr *ForumStarted) Subject() string {
	return rcvr.subject
}

func (rcvr *ForumStarted) Tenant() tenant.ID {
	return rcvr.tenant
}
