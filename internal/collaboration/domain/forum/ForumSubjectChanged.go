package forum

import (
	"time"

	"github.com/maranqz/go-IDDD_Samples/internal/collaboration/domain/tenant"
)

type ForumSubjectChanged struct {
	eventVersion   int
	exclusiveOwner string
	forumId        ForumID
	occurredOn     time.Time
	subject        string
	tenant         tenant.ID
}

func NewForumSubjectChanged(aTenant tenant.ID, aForumId ForumID, aSubject string, anExclusiveOwner string) (rcvr *ForumSubjectChanged) {
	rcvr = &ForumSubjectChanged{}
	rcvr.eventVersion = 1
	rcvr.exclusiveOwner = anExclusiveOwner
	rcvr.forumId = aForumId
	rcvr.occurredOn = NewDate()
	rcvr.subject = aSubject
	rcvr.tenant = aTenant
	return
}

func (rcvr *ForumSubjectChanged) EventVersion() int {
	return rcvr.eventVersion
}

func (rcvr *ForumSubjectChanged) ExclusiveOwner() string {
	return rcvr.exclusiveOwner
}

func (rcvr *ForumSubjectChanged) ForumId() ForumID {
	return rcvr.forumId
}

func (rcvr *ForumSubjectChanged) OccurredOn() time.Time {
	return rcvr.occurredOn
}

func (rcvr *ForumSubjectChanged) Subject() string {
	return rcvr.subject
}

func (rcvr *ForumSubjectChanged) Tenant() tenant.ID {
	return rcvr.tenant
}
