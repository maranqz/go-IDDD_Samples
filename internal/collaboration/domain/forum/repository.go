package forum

import "github.com/maranqz/go-IDDD_Samples/internal/collaboration/domain/tenant"

type DiscussionRepository interface {
	DiscussionOfId(aTenantId tenant.ID, aDiscussionId DiscussionID) *Discussion
	NextIdentity() DiscussionID
	Save(aDiscussion *Discussion)
}

type ForumRepository interface {
	ForumOfId(aTenant tenant.ID, aForumId ForumID) *Forum
	NextIdentity() ForumID
	Save(aForum *Forum)
}

type PostRepository interface {
	NextIdentity() PostID
	PostOfId(aTenant tenant.ID, aPostId PostID) *Post
	Save(aPost *Post)
}
