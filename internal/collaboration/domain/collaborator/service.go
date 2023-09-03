package collaborator

type CollaboratorService interface {
	AuthorFrom(aTenant TenantID, anIdentity string) (*collaborator.Author, error)
	CreatorFrom(aTenant TenantID, anIdentity string) (*Creator, error)
	ModeratorFrom(aTenant TenantID, anIdentity string) (*Moderator, error)
	OwnerFrom(aTenant TenantID, anIdentity string) (*Owner, error)
	ParticipantFrom(aTenant TenantID, anIdentity string) (*Participant, error)
}
