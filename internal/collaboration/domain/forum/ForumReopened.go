package forum

import (
	"time"

	"github.com/maranqz/go-IDDD_Samples/internal/collaboration/domain/tenant"
)

type ForumReopened struct {
	eventVersion   int
	exclusiveOwner string
	forumId        ForumID
	occurredOn     time.Time
	tenant         tenant.ID
}

func NewForumReopened(aTenant tenant.ID, aForumId ForumID, anExclusiveOwner string) (rcvr *ForumReopened) {
	rcvr = &ForumReopened{}
	rcvr.eventVersion = 1
	rcvr.exclusiveOwner = anExclusiveOwner
	rcvr.forumId = aForumId
	rcvr.occurredOn = NewDate()
	rcvr.tenant = aTenant
	return
}

func (rcvr *ForumReopened) EventVersion() int {
	return rcvr.eventVersion
}

func (rcvr *ForumReopened) ExclusiveOwner() string {
	return rcvr.exclusiveOwner
}

func (rcvr *ForumReopened) ForumId() ForumID {
	return rcvr.forumId
}

func (rcvr *ForumReopened) OccurredOn() time.Time {
	return rcvr.occurredOn
}

func (rcvr *ForumReopened) Tenant() tenant.ID {
	return rcvr.tenant
}
