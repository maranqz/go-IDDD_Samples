package port_adapter

import (
	"github.com/maranqz/go-IDDD_Samples/internal/collaboration/domain/collaborator"
	"github.com/maranqz/go-IDDD_Samples/internal/collaboration/domain/tenant"
	"github.com/maranqz/go-IDDD_Samples/internal/common/domain"
)

type role = string

const (
	authorRole      role = "Author"
	creatorRole     role = "Creator"
	moderatorRole   role = "Moderator"
	ownerRole       role = "Owner"
	participantRole role = "Participant"
)

type TranslatingCollaboratorSrvc struct {
	userInRoleAdapter UserInRoleAdapter
}

func NewTranslatingCollaboratorService(aUserInRoleAdapter UserInRoleAdapter) *TranslatingCollaboratorSrvc {
	return &TranslatingCollaboratorSrvc{userInRoleAdapter: aUserInRoleAdapter}
}

func (t *TranslatingCollaboratorSrvc) AuthorFrom(aTenant tenant.ID, anIdentity domain.UUID) (*collaborator.Author, error) {
	author, err := t.userInRoleAdapter.
		ToCollaborator(aTenant, anIdentity, authorRole)
	if err != nil {
		return nil, err
	}

	return collaborator.AuthorFromCollaborator(author)
}

func (t *TranslatingCollaboratorSrvc) CreatorFrom(aTenant tenant.ID, anIdentity domain.UUID) (*collaborator.Creator, error) {
	creator, err := t.userInRoleAdapter.
		ToCollaborator(aTenant, anIdentity, creatorRole)
	if err != nil {
		return nil, err
	}

	return collaborator.CreatorFromCollaborator(creator)
}

func (t *TranslatingCollaboratorSrvc) ModeratorFrom(aTenant tenant.ID, anIdentity domain.UUID) (*collaborator.Moderator, error) {
	moderator, err := t.userInRoleAdapter.
		ToCollaborator(aTenant, anIdentity, moderatorRole)
	if err != nil {
		return nil, err
	}

	return collaborator.ModeratorFromCollaborator(moderator)
}

func (t *TranslatingCollaboratorSrvc) OwnerFrom(aTenant tenant.ID, anIdentity domain.UUID) (*collaborator.Owner, error) {
	owner, err := t.userInRoleAdapter.
		ToCollaborator(aTenant, anIdentity, ownerRole)
	if err != nil {
		return nil, err
	}

	return collaborator.OwnerFromCollaborator(owner)
}

func (t *TranslatingCollaboratorSrvc) ParticipantFrom(aTenant tenant.ID, anIdentity domain.UUID) (*collaborator.Participant, error) {
	participant, err := t.userInRoleAdapter.
		ToCollaborator(aTenant, anIdentity, participantRole)
	if err != nil {
		return nil, err
	}

	return collaborator.ParticipantFromCollaborator(participant)
}
