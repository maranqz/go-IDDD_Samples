package forum

import (
	"errors"
	"fmt"

	"github.com/maranqz/go-IDDD_Samples/internal/collaboration/domain/collaborator"
	"github.com/maranqz/go-IDDD_Samples/internal/collaboration/domain/tenant"
	"github.com/maranqz/go-IDDD_Samples/internal/common/domain"
)

var (
	ErrForum                   = errors.New("forum")
	ErrForumCreatorEmpty       = errors.New("creator: empty")
	ErrForumClosed             = fmt.Errorf("%w: closed", ErrForum)
	ErrForumOpened             = fmt.Errorf("%w: opened", ErrForum)
	ErrForumDescEmpty          = fmt.Errorf("%w: description: empty", ErrForum)
	ErrForumSubjectEmpty       = fmt.Errorf("%w: subject: empty", ErrForum)
	ErrForumPost               = fmt.Errorf("%w: post", ErrForum)
	ErrForumPostEmpty          = fmt.Errorf("%w: empty", ErrForumPost)
	ErrForumPostForbidden      = fmt.Errorf("%w: invalid forum", ErrForumPost)
	ErrForumModeratorEmpty     = fmt.Errorf("%w: moderatory: empty", ErrForum)
	ErrForumModeratorForbidden = fmt.Errorf("%w: moderatory: invalid forum", ErrForum)
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

func NewForum(
	aTenant tenant.ID,
	aForumId ForumID,
	// TODO Why is it not ID?
	aCreator *collaborator.Creator,
	aModerator *collaborator.Moderator,
	aSubject string,
	aDescription string,
	anExclusiveOwner string,
) (*Forum, error) {
	if aCreator.IsEmpty() {
		return nil, ErrForumCreatorEmpty
	}

	if aSubject == "" {
		return nil, ErrForumSubjectEmpty
	}

	if aDescription == "" {
		return nil, ErrForumDescEmpty
	}

	// TODO think to use AssignModerator
	if aModerator.IsEmpty() {
		return nil, ErrForumModeratorEmpty
	}

	// TODO this one or fill fields?
	f := &Forum{}
	f.setCreator(aCreator)
	f.setDescription(aDescription)
	f.setExclusiveOwner(anExclusiveOwner)
	f.setForumId(aForumId)
	f.setModerator(aModerator)
	f.setSubject(aSubject)
	f.setTenant(aTenant)

	return f, nil
}

func (f *Forum) AssignModerator(aModerator *collaborator.Moderator) error {
	if err := f.assertClosed(); err != nil {
		return err
	}

	// TODO think to move in setModerator
	if aModerator.IsEmpty() {
		return ErrForumModeratorEmpty
	}

	f.setModerator(aModerator)

	return nil
}

func (f *Forum) ChangeDescription(aDescription string) error {
	if err := f.assertClosed(); err != nil {
		return err
	}

	if aDescription == "" {
		return ErrForumDescEmpty
	}

	f.setDescription(aDescription)

	return nil
}

func (f *Forum) ChangeSubject(aSubject string) error {
	if err := f.assertClosed(); err != nil {
		return err
	}

	if aSubject == "" {
		return ErrForumSubjectEmpty
	}

	f.setSubject(aSubject)

	return nil
}

func (f *Forum) Close() error {
	if err := f.assertClosed(); err != nil {
		return err
	}

	f.setClosed(false)

	return nil
}

func (f *Forum) Creator() *collaborator.Creator {
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
	return f.ExclusiveOwner() != ""
}

func (f *Forum) IsClosed() bool {
	return f.closed
}

func (f *Forum) IsModeratedBy(aModerator *collaborator.Moderator) bool {
	return !aModerator.IsEmpty() &&
		f.Moderator().Identity() == aModerator.Identity()
}

func (f *Forum) ModeratePost(
	aPost *Post,
	aModerator *collaborator.Moderator,
	aSubject string,
	aBodyText string,
) error {
	if err := f.assertClosed(); err != nil {
		return err
	}

	if aPost == nil {
		return ErrForumPostEmpty
	} else if f.ForumId() != aPost.ForumId() {
		return ErrForumPostForbidden
	}

	if f.IsModeratedBy(aModerator) {
		return ErrForumModeratorForbidden
	}

	return aPost.alterPostContent(aSubject, aBodyText)
}

func (f *Forum) Moderator() *collaborator.Moderator {
	return f.moderator
}

func (f *Forum) Reopen() error {
	if err := f.assertOpened(); err != nil {
		return err
	}

	f.setClosed(false)

	return nil
}

func (f *Forum) assertClosed() error {
	if f.IsClosed() {
		return ErrForumClosed
	}

	return nil
}

func (f *Forum) assertOpened() error {
	if !f.IsClosed() {
		return ErrForumOpened
	}

	return nil
}

func (f *Forum) setClosed(isClosed bool) {
	f.closed = isClosed
}

func (f *Forum) setCreator(aCreator *collaborator.Creator) {
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

func (f *Forum) StartDiscussion(aForumIdentityService IdentityService, anAuthor *collaborator.Author, aSubject string) (*Discussion, error) {
	return f.StartDiscussionFor(
		aForumIdentityService,
		anAuthor,
		aSubject,
		"",
	)
}

func (f *Forum) StartDiscussionFor(
	aForumIdentityService IdentityService,
	anAuthor *collaborator.Author,
	aSubject string,
	anExclusiveOwner string,
) (*Discussion, error) {
	if err := f.assertClosed(); err != nil {
		return nil, err
	}

	return newDiscussion(
		f.Tenant(),
		f.ForumId(),
		aForumIdentityService.NextDiscussionId(),
		anAuthor,
		aSubject,
		anExclusiveOwner,
	)
}

func (f *Forum) Subject() string {
	return f.subject
}

func (f *Forum) Tenant() tenant.ID {
	return f.tenant
}
