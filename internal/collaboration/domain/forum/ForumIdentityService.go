package forum

type ForumIdentityService struct {
	discussionRepository *DiscussionRepository
	forumRepository      *ForumRepository
	postRepository       *PostRepository
}

func NewForumIdentityService(aForumRepository *ForumRepository, aDiscussionRepository *DiscussionRepository, aPostRepository *PostRepository) (rcvr ForumIdentityService) {
	rcvr = &ForumIdentityService{}
	rcvr.discussionRepository = aDiscussionRepository
	rcvr.forumRepository = aForumRepository
	rcvr.postRepository = aPostRepository
	return
}

func (rcvr ForumIdentityService) discussionRepository() *DiscussionRepository {
	return rcvr.discussionRepository
}

func (rcvr ForumIdentityService) forumRepository() *ForumRepository {
	return rcvr.forumRepository
}

func (rcvr ForumIdentityService) NextDiscussionId() DiscussionID {
	return this.discussionRepository().nextIdentity()
}

func (rcvr ForumIdentityService) NextForumID() ForumID {
	return this.forumRepository().nextIdentity()
}

func (rcvr ForumIdentityService) NextPostId() PostID {
	return this.postRepository().nextIdentity()
}

func (rcvr ForumIdentityService) postRepository() *PostRepository {
	return rcvr.postRepository
}
