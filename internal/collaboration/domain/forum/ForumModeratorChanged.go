package forum

import (
	"time"

	"github.com/maranqz/go-IDDD_Samples/internal/collaboration/domain/tenant"
)

type ForumModeratorChanged struct {
	eventVersion   int
	exclusiveOwner string
	forumId        ForumID
	moderator      *Moderator
	occurredOn     time.Time
	tenant         tenant.ID
}

func NewForumModeratorChanged(aTenant tenant.ID, aForumId ForumID, aModerator *Moderator, anExclusiveOwner string) (rcvr *ForumModeratorChanged) {
	rcvr = &ForumModeratorChanged{}
	rcvr.eventVersion = 1
	rcvr.exclusiveOwner = anExclusiveOwner
	rcvr.forumId = aForumId
	rcvr.moderator = aModerator
	rcvr.occurredOn = NewDate()
	rcvr.tenant = aTenant
	return
}

func (rcvr *ForumModeratorChanged) EventVersion() int {
	return rcvr.eventVersion
}

func (rcvr *ForumModeratorChanged) ExclusiveOwner() string {
	return rcvr.exclusiveOwner
}

func (rcvr *ForumModeratorChanged) ForumId() ForumID {
	return rcvr.forumId
}

func (rcvr *ForumModeratorChanged) Moderator() *Moderator {
	return rcvr.moderator
}

func (rcvr *ForumModeratorChanged) OccurredOn() time.Time {
	return rcvr.occurredOn
}

func (rcvr *ForumModeratorChanged) Tenant() tenant.ID {
	return rcvr.tenant
}
