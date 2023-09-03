package forum

import (
	"github.com/maranqz/go-IDDD_Samples/internal/collaboration/domain/collaborator"
	"github.com/maranqz/go-IDDD_Samples/internal/collaboration/domain/tenant"
	"github.com/maranqz/go-IDDD_Samples/internal/common/domain"
)

type ForumID struct {
	domain.UUID
}

type Forum struct {
	closed         bool
	creator        *collaborator.Creator
	description    string
	exclusiveOwner string
	forumId        ForumID
	moderator      *collaborator.Moderator
	subject        string
	tenant         tenant.ID
}

func NewForum(aTenant tenant.ID, aForumId ForumID, aCreator *Creator, aModerator *collaborator.Moderator, aSubject string, aDescription string, anExclusiveOwner string) (rcvr *Forum) {
	rcvr = NewForum()
	f.assertArgumentNotNull(aCreator, "The creator must be provided.")
	f.assertArgumentNotEmpty(aDescription, "The description must be provided.")
	f.assertArgumentNotNull(aForumId, "The forum id must be provided.")
	f.assertArgumentNotNull(aModerator, "The moderator must be provided.")
	f.assertArgumentNotEmpty(aSubject, "The subject must be provided.")
	f.assertArgumentNotNull(aTenant, "The creator must be provided.")
	f.apply(NewForumStarted(aTenant, aForumId, aCreator, aModerator, aSubject, aDescription, anExclusiveOwner))
	return
}

func (f *Forum) AssignModerator(aModerator *collaborator.Moderator) {
	f.assertStateFalse(f.isClosed(), "Forum is closed.")
	f.assertArgumentNotNull(aModerator, "The moderator must be provided.")
	f.apply(NewForumModeratorChanged(f.tenant(), f.forumId(), aModerator, f.exclusiveOwner()))
}

func (f *Forum) ChangeDescription(aDescription string) {
	f.assertStateFalse(f.isClosed(), "Forum is closed.")
	f.assertArgumentNotEmpty(aDescription, "The description must be provided.")
	f.apply(NewForumDescriptionChanged(f.tenant(), f.forumId(), aDescription, f.exclusiveOwner()))
}

func (f *Forum) ChangeSubject(aSubject string) {
	f.assertStateFalse(f.isClosed(), "Forum is closed.")
	f.assertArgumentNotEmpty(aSubject, "The subject must be provided.")
	f.apply(NewForumSubjectChanged(f.tenant(), f.forumId(), aSubject, f.exclusiveOwner()))
}

func (f *Forum) Close() {
	f.assertStateFalse(f.isClosed(), "Forum is already closed.")
	f.apply(NewForumClosed(f.tenant(), f.forumId(), f.exclusiveOwner()))
}

func (f *Forum) Creator() *Creator {
	return f.creator
}

func (f *Forum) Description() string {
	return f.description
}

func (f *Forum) ExclusiveOwner() string {
	return f.exclusiveOwner
}

func (f *Forum) ForumId() ForumID {
	return f.forumId
}

func (f *Forum) HasExclusiveOwner() bool {
	return f.ExclusiveOwner() != nil
}

func (f *Forum) IsClosed() bool {
	return f.closed
}

func (f *Forum) IsModeratedBy(aModerator *collaborator.Moderator) bool {
	return f.Moderator().equals(aModerator)
}

func (f *Forum) ModeratePost(aPost *Post, aModerator *collaborator.Moderator, aSubject string, aBodyText string) {
	f.assertStateFalse(f.isClosed(), "Forum is closed.")
	f.assertArgumentNotNull(aPost, "Post may not be null.")
	f.assertArgumentEquals(aPost.forumId(), f.forumId(), "Not a post of this forum.")
	f.assertArgumentTrue(f.isModeratedBy(aModerator), "Not the moderator of this forum.")
	aPost.alterPostContent(aSubject, aBodyText)
}

func (f *Forum) Moderator() *collaborator.Moderator {
	return f.moderator
}

func (f *Forum) Reopen() {
	f.assertStateTrue(f.isClosed(), "Forum is not closed.")
	f.apply(NewForumReopened(f.tenant(), f.forumId(), f.exclusiveOwner()))
}

func (f *Forum) setClosed(isClosed bool) {
	f.closed = isClosed
}

func (f *Forum) setCreator(aCreator *Creator) {
	f.creator = aCreator
}

func (f *Forum) setDescription(aDescription string) {
	f.description = aDescription
}

func (f *Forum) setExclusiveOwner(anExclusiveOwner string) {
	f.exclusiveOwner = anExclusiveOwner
}

func (f *Forum) setForumId(aForumId ForumID) {
	f.forumId = aForumId
}

func (f *Forum) setModerator(aModerator *collaborator.Moderator) {
	f.moderator = aModerator
}

func (f *Forum) setSubject(aSubject string) {
	f.subject = aSubject
}

func (f *Forum) setTenant(aTenant tenant.ID) {
	f.tenant = aTenant
}

func (f *Forum) StartDiscussion(aForumIdentityService ForumIDentityService, anAuthor *collaborator.Author, aSubject string) *Discussion {
	return f.startDiscussionFor(aForumIdentityService, anAuthor, aSubject, nil)
}

func (f *Forum) StartDiscussionFor(aForumIdentityService ForumIDentityService, anAuthor *collaborator.Author, aSubject string, anExclusiveOwner string) *Discussion {
	if f.isClosed() {
		throw(NewIllegalStateException("Forum is closed."))
	}
	discussion := NewDiscussion(f.tenant(), f.forumId(), aForumIdentityService.nextDiscussionId(), anAuthor, aSubject, anExclusiveOwner)
	return discussion
}

func (f *Forum) Subject() string {
	return f.subject
}

func (f *Forum) Tenant() tenant.ID {
	return f.tenant
}

func (f *Forum) when(anEvent *ForumClosed) {
	f.setClosed(true)
}

func (f *Forum) when2(anEvent *ForumDescriptionChanged) {
	f.setDescription(anEvent.description())
}

func (f *Forum) when3(anEvent *ForumModeratorChanged) {
	f.setModerator(anEvent.moderator())
}

func (f *Forum) when4(anEvent *ForumReopened) {
	f.setClosed(false)
}

func (f *Forum) when5(anEvent *ForumStarted) {
	f.setCreator(anEvent.creator())
	f.setDescription(anEvent.description())
	f.setExclusiveOwner(anEvent.exclusiveOwner())
	f.setForumId(anEvent.forumId())
	f.setModerator(anEvent.moderator())
	f.setSubject(anEvent.subject())
	f.setTenant(anEvent.tenant())
}

func (f *Forum) when6(anEvent *ForumSubjectChanged) {
	f.setSubject(anEvent.subject())
}
