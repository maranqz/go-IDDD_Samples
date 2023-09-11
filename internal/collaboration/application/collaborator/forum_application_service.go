package forum

import (
	dforum "github.com/maranqz/go-IDDD_Samples/internal/collaboration/domain/forum"
	"github.com/maranqz/go-IDDD_Samples/internal/collaboration/domain/tenant"
	"github.com/maranqz/go-IDDD_Samples/internal/common/domain"
)

type ForumApplicationService struct {
	collaboratorService    *CollaboratorService
	discussionQueryService *DiscussionQueryService
	discussionRepository   *DiscussionRepository
	forumIdentityService   *ForumIdentityService
	forumQueryService      *ForumQueryService
	forumRepository        *ForumRepository
}

func NewForumApplicationService(aForumQueryService *ForumQueryService, aForumRepository *ForumRepository, aForumIdentityService *ForumIdentityService, aDiscussionQueryService *DiscussionQueryService, aDiscussionRepository *DiscussionRepository, aCollaboratorService *CollaboratorService) (rcvr *ForumApplicationService) {
	rcvr = &ForumApplicationService{}
	rcvr.collaboratorService = aCollaboratorService
	rcvr.discussionQueryService = aDiscussionQueryService
	rcvr.discussionRepository = aDiscussionRepository
	rcvr.forumIdentityService = aForumIdentityService
	rcvr.forumQueryService = aForumQueryService
	rcvr.forumRepository = aForumRepository
	return
}
func (f *ForumApplicationService) AssignModeratorToForum(aTenantId string, aForumId string, aModeratorId string) {
	tenant := NewTenant(aTenantId)
	forum := this.forumRepository().forumOfId(tenant, NewForumId(aForumId))
	moderator := this.collaboratorService().moderatorFrom(tenant, aModeratorId)
	forum.assignModerator(moderator)
	this.forumRepository().save(forum)
}
func (f *ForumApplicationService) ChangeForumDescription(aTenantId string, aForumId string, aDescription string) {
	tenant := NewTenant(aTenantId)
	forum := this.forumRepository().forumOfId(tenant, NewForumId(aForumId))
	forum.changeDescription(aDescription)
	this.forumRepository().save(forum)
}
func (f *ForumApplicationService) ChangeForumSubject(aTenantId string, aForumId string, aSubject string) {
	tenant := NewTenant(aTenantId)
	forum := this.forumRepository().forumOfId(tenant, NewForumId(aForumId))
	forum.changeSubject(aSubject)
	this.forumRepository().save(forum)
}
func (f *ForumApplicationService) CloseForum(aTenantId string, aForumId string) {
	tenant := NewTenant(aTenantId)
	forum := this.forumRepository().forumOfId(tenant, NewForumId(aForumId))
	forum.close()
	this.forumRepository().save(forum)
}
func (f *ForumApplicationService) collaboratorService() *CollaboratorService {
	return f.collaboratorService
}
func (f *ForumApplicationService) discussionQueryService() *DiscussionQueryService {
	return f.discussionQueryService
}
func (f *ForumApplicationService) discussionRepository() *DiscussionRepository {
	return f.discussionRepository
}
func (f *ForumApplicationService) forumIdentityService() *ForumIdentityService {
	return f.forumIdentityService
}
func (f *ForumApplicationService) forumQueryService() *ForumQueryService {
	return f.forumQueryService
}
func (f *ForumApplicationService) forumRepository() *dforum.ForumRepository {
	return f.forumRepository
}
func (f *ForumApplicationService) ReopenForum(aTenantId string, aForumId string) {
	tenant := NewTenant(aTenantId)
	forum := this.forumRepository().forumOfId(tenant, NewForumId(aForumId))
	forum.reopen()
	this.forumRepository().save(forum)
}
func (f *ForumApplicationService) StartExclusiveForum(aTenantId string, anExclusiveOwner string, aCreatorId string, aModeratorId string, aSubject string, aDescription string, aResult *ForumCommandResult) {
	tenant := NewTenant(aTenantId)
	forumId := this.forumQueryService().forumIdOfExclusiveOwner(aTenantId, anExclusiveOwner)
	forum := nil
	if forumId != nil {
		forum = this.forumRepository().forumOfId(tenant, NewForumId(forumId))
	}
	if forum == nil {
		forum = this.startNewForum(tenant, aCreatorId, aModeratorId, aSubject, aDescription, anExclusiveOwner)
	}
	if aResult != nil {
		aResult.resultingForumId(forum.forumId().id())
	}
}
func (f *ForumApplicationService) StartExclusiveForumWithDiscussion(aTenantId string, anExclusiveOwner string, aCreatorId string, aModeratorId string, anAuthorId string, aForumSubject string, aForumDescription string, aDiscussionSubject string, aResult *ForumCommandResult) {
	tenant := NewTenant(aTenantId)
	forumId := this.forumQueryService().forumIdOfExclusiveOwner(aTenantId, anExclusiveOwner)
	forum := nil
	if forumId != nil {
		forum = this.forumRepository().forumOfId(tenant, NewForumId(forumId))
	}
	if forum == nil {
		forum = this.startNewForum(tenant, aCreatorId, aModeratorId, aForumSubject, aForumDescription, anExclusiveOwner)
	}
	discussionId := this.discussionQueryService().discussionIdOfExclusiveOwner(aTenantId, anExclusiveOwner)
	discussion := nil
	if discussionId != nil {
		discussion = this.discussionRepository().discussionOfId(tenant, NewDiscussionId(discussionId))
	}
	if discussion == nil {
		author := this.collaboratorService().authorFrom(tenant, anAuthorId)
		discussion = forum.startDiscussionFor(this.forumIdentityService(), author, aDiscussionSubject, anExclusiveOwner)
		this.discussionRepository().save(discussion)
	}
	if aResult != nil {
		aResult.resultingForumId(forum.forumId().id())
		aResult.resultingDiscussionId(discussion.discussionId().id())
	}
}

func (f *ForumApplicationService) StartForum(
	aTenantID tenant.ID,
	aCreatorID domain.UUID,
	aModeratorID domain.UUID,
	aSubject Subject,
	aDescription string,
	anExclusiveOwner string,
) (*dforum.Forum, error) {
	creator, err := f.collaboratorService().CreatorFrom(aTenantID, aCreatorID)
	if err != nil {
		return nil, err
	}

	moderator := f.collaboratorService().ModeratorFrom(aTenantID, aModeratorID)
	if err != nil {
		return nil, err
	}

	return dforum.NewForum(
		aTenantID,
		f.forumRepository().NextIdentity(),
		creator,
		moderator,
		aSubject,
		aDescription,
		anExclusiveOwner,
	)
}
