package collaborator

import "github.com/maranqz/go-IDDD_Samples/internal/collaboration/domain/tenant"

type CollaboratorService interface {
	AuthorFrom(aTenant tenant.ID, anIdentity string) (*Author, error)
	CreatorFrom(aTenant tenant.ID, anIdentity string) (*Creator, error)
	ModeratorFrom(aTenant tenant.ID, anIdentity string) (*Moderator, error)
	OwnerFrom(aTenant tenant.ID, anIdentity string) (*Owner, error)
	ParticipantFrom(aTenant tenant.ID, anIdentity string) (*Participant, error)
}
