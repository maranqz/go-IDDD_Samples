package forum

import (
	"errors"
	"fmt"
	"time"

	"github.com/maranqz/go-IDDD_Samples/internal/collaboration/domain/collaborator"
	"github.com/maranqz/go-IDDD_Samples/internal/collaboration/domain/tenant"
	"github.com/maranqz/go-IDDD_Samples/internal/common/domain"
)

var (
	ErrPost = errors.New("post")

	ErrPostSubjectEmpty  = fmt.Errorf("%w: subject: empty", ErrDiscussion)
	ErrPostBodyTextEmpty = fmt.Errorf("%w: description: empty", ErrDiscussion)
)

type PostID struct {
	domain.UUID
}

type Post struct {
	author        *collaborator.Author
	bodyText      string
	changedOn     time.Time
	createdOn     time.Time
	discussionId  DiscussionID
	forumId       ForumID
	postId        PostID
	replyToPostId *PostID
	subject       string
	tenant        tenant.ID
}

func NewPost(
	aTenant tenant.ID,
	aForumId ForumID,
	aDiscussionId DiscussionID,
	aPostId PostID,
	anAuthor *collaborator.Author,
	aSubject string,
	aBodyText string,
) (*Post, error) {
	return NewPostReply(
		aTenant,
		aForumId,
		aDiscussionId,
		nil,
		aPostId,
		anAuthor,
		aSubject,
		aBodyText,
	)
}

func NewPostReply(
	aTenant tenant.ID,
	aForumID ForumID,
	aDiscussionID DiscussionID,
	aReplyToPost *PostID,
	aPostID PostID,
	anAuthor *collaborator.Author,
	aSubject string,
	aBodyText string,
) (*Post, error) {

	if err := assertAuthor(anAuthor, ErrPost); err != nil {
		return nil, err
	}

	if aSubject == "" {
		return nil, ErrPostSubjectEmpty
	}

	if aBodyText == "" {
		return nil, ErrPostBodyTextEmpty
	}

	p := &Post{}
	p.setAuthor(anAuthor)
	p.setBodyText(aBodyText)
	p.setChangedOn(time.Now())
	p.setCreatedOn(time.Now())
	p.setDiscussionID(aDiscussionID)
	p.setForumID(aForumID)
	p.setPostID(aPostID)
	p.setReplyToPostId(aReplyToPost)
	p.setSubject(aSubject)
	p.setTenant(aTenant)

	return p, nil
}

func (p *Post) alterPostContent(aSubject string, aBodyText string) error {
	// TODO how to distinguish ErrPostSubjectEmpty here and in NewPostReply
	if aSubject == "" {
		return ErrPostSubjectEmpty
	}

	if aBodyText == "" {
		return ErrPostBodyTextEmpty
	}

	p.setBodyText(aBodyText)
	p.setChangedOn(time.Now())
	p.setSubject(aSubject)

	return nil
}

func (p *Post) Author() *collaborator.Author {
	return p.author
}

func (p *Post) BodyText() string {
	return p.bodyText
}

func (p *Post) ChangedOn() time.Time {
	return p.changedOn
}

func (p *Post) CreatedOn() time.Time {
	return p.createdOn
}

func (p *Post) DiscussionID() DiscussionID {
	return p.discussionId
}

func (p *Post) ForumId() ForumID {
	return p.forumId
}

func (p *Post) PostID() PostID {
	return p.postId
}

func (p *Post) ReplyToPostId() *PostID {
	return p.replyToPostId
}

func (p *Post) setAuthor(anAuthor *collaborator.Author) {
	p.author = anAuthor
}

func (p *Post) setBodyText(aBodyText string) {
	p.bodyText = aBodyText
}

func (p *Post) setChangedOn(aChangedOnDate time.Time) {
	p.changedOn = aChangedOnDate
}

func (p *Post) setCreatedOn(aCreatedOnDate time.Time) {
	p.createdOn = aCreatedOnDate
}

func (p *Post) setDiscussionID(aDiscussionId DiscussionID) {
	p.discussionId = aDiscussionId
}

func (p *Post) setForumID(aForumId ForumID) {
	p.forumId = aForumId
}

func (p *Post) setPostID(aPostId PostID) {
	p.postId = aPostId
}

func (p *Post) setReplyToPostId(aReplyToPostId *PostID) {
	p.replyToPostId = aReplyToPostId
}

func (p *Post) setSubject(aSubject string) {
	p.subject = aSubject
}

func (p *Post) setTenant(aTenant tenant.ID) {
	p.tenant = aTenant
}

func (p *Post) Subject() string {
	return p.subject
}

func (p *Post) Tenant() tenant.ID {
	return p.tenant
}
