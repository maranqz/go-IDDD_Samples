package forum

type IdentityService struct {
	discussionRepository DiscussionRepository
	forumRepository      ForumRepository
	postRepository       PostRepository
}

func NewIdentityService(
	aForumRepository ForumRepository,
	aDiscussionRepository DiscussionRepository,
	aPostRepository PostRepository,
) (*IdentityService, error) {
	return &IdentityService{
		discussionRepository: aDiscussionRepository,
		forumRepository:      aForumRepository,
		postRepository:       aPostRepository,
	}, nil
}

func (f IdentityService) NextDiscussionId() DiscussionID {
	return f.discussionRepository.NextIdentity()
}

func (f IdentityService) NextForumID() ForumID {
	return f.forumRepository.NextIdentity()
}

func (f IdentityService) NextPostId() PostID {
	return f.postRepository.NextIdentity()
}
