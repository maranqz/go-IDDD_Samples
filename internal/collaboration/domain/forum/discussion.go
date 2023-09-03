package forum

import (
	"errors"
	"fmt"

	"github.com/maranqz/go-IDDD_Samples/internal/collaboration/domain/collaborator"
	"github.com/maranqz/go-IDDD_Samples/internal/collaboration/domain/tenant"
	"github.com/maranqz/go-IDDD_Samples/internal/common/domain"
)

var (
	ErrDiscussion             = errors.New("discussion")
	ErrDiscussionSubjectEmpty = fmt.Errorf("%w: subject: empty", ErrDiscussion)
	ErrDiscussionClosed       = fmt.Errorf("%w: already closed", ErrDiscussion)
	ErrDiscussionOpened       = fmt.Errorf("%w: already opened", ErrDiscussion)
)

type DiscussionID struct {
	domain.UUID
}

type Discussion struct {
	author         *collaborator.Author
	closed         bool
	discussionID   DiscussionID
	exclusiveOwner string
	forumID        ForumID
	subject        string
	tenant         tenant.ID
}

func NewDiscussion(
	aTenantID tenant.ID,
	aForumID ForumID,
	aDiscussionID DiscussionID,
	anAuthor *collaborator.Author,
	aSubject string,
	anExclusiveOwner string,
) (*Discussion, error) {
	if err := assertAuthor(anAuthor, ErrDiscussion); err != nil {
		return nil, err
	}

	if aSubject == "" {
		return nil, ErrDiscussionSubjectEmpty
	}

	d := &Discussion{}

	d.setAuthor(anAuthor)
	d.setDiscussionID(aDiscussionID)
	d.setExclusiveOwner(anExclusiveOwner)
	d.setForumID(aForumID)
	d.setSubject(aSubject)
	d.setTenant(aTenantID)

	return d, nil
}

func (d *Discussion) Author() *collaborator.Author {
	return d.author
}

// TODO Could we do it idenpotent???
func (d *Discussion) Close() error {
	if d.IsClosed() {
		return ErrDiscussionClosed
	}

	d.setClosed(true)

	return nil
}

func (d *Discussion) DiscussionID() DiscussionID {
	return d.discussionID
}

func (d *Discussion) ExclusiveOwner() string {
	return d.exclusiveOwner
}

func (d *Discussion) ForumID() ForumID {
	return d.forumID
}

func (d *Discussion) IsClosed() bool {
	return d.closed
}

// Reply was a Post in Java
func (d *Discussion) Reply(
	aForumIdentityService ForumIdentityService,
	aReplyToPost *PostID,
	anAuthor *collaborator.Author,
	aSubject string,
	aBodyText string,
) (*Post, error) {
	return NewPostReply(
		d.Tenant(),
		d.ForumID(),
		d.DiscussionID(),
		aReplyToPost,
		aForumIdentityService.NextPostId(),
		anAuthor,
		aSubject,
		aBodyText,
	)
}

func (d *Discussion) Post(
	aForumIdentityService ForumIdentityService,
	anAuthor *collaborator.Author,
	aSubject string,
	aBodyText string,
) (*Post, error) {
	return d.Reply(aForumIdentityService, nil, anAuthor, aSubject, aBodyText)
}

func (d *Discussion) Reopen() error {
	if !d.IsClosed() {
		return ErrDiscussionOpened
	}

	d.setClosed(false)

	return nil
}

// TODO think to remove or use everywhere
func (d *Discussion) setAuthor(author *collaborator.Author) {
	d.author = author
}

func (d *Discussion) setClosed(isClosed bool) {
	d.closed = isClosed
}

func (d *Discussion) setDiscussionID(aDiscussionId DiscussionID) {
	d.discussionID = aDiscussionId
}

func (d *Discussion) setExclusiveOwner(anExclusiveOwner string) {
	d.exclusiveOwner = anExclusiveOwner
}

func (d *Discussion) setForumID(aForumId ForumID) {
	d.forumID = aForumId
}

func (d *Discussion) setSubject(aSubject string) {
	d.subject = aSubject
}

func (d *Discussion) setTenant(aTenant tenant.ID) {
	d.tenant = aTenant
}

func (d *Discussion) Subject() string {
	return d.subject
}

func (d *Discussion) Tenant() tenant.ID {
	return d.tenant
}
