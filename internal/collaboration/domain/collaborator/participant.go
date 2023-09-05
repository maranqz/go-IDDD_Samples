package collaborator

import (
	"github.com/maranqz/go-IDDD_Samples/internal/common/domain"
)

type Participant struct {
	collaborator
}

func NewParticipant(identity domain.UUID, name string, emailAddress string) (*Participant, error) {
	c, err := newCollaborator(identity, name, emailAddress)
	if err != nil {
		return nil, err
	}

	return &Participant{
		collaborator: c,
	}, nil
}

type Owner struct {
	collaborator
}

func NewOwner(identity domain.UUID, name string, emailAddress string) (*Owner, error) {
	c, err := newCollaborator(identity, name, emailAddress)
	if err != nil {
		return nil, err
	}

	return &Owner{
		collaborator: c,
	}, nil
}

type Author struct {
	collaborator
}

func NewAuthor(
	identity domain.UUID,
	name string,
	emailAddress string,
) (*Author, error) {
	c, err := newCollaborator(identity, name, emailAddress)
	if err != nil {
		return nil, err
	}

	return &Author{
		collaborator: c,
	}, nil
}

type Creator struct {
	collaborator
}

func NewCreator(identity domain.UUID, name string, emailAddress string) (*Creator, error) {
	c, err := newCollaborator(identity, name, emailAddress)
	if err != nil {
		return nil, err
	}

	return &Creator{
		collaborator: c,
	}, nil
}

type Moderator struct {
	collaborator
}

func NewModerator(identity domain.UUID, name string, emailAddress string) (*Moderator, error) {
	c, err := newCollaborator(identity, name, emailAddress)
	if err != nil {
		return nil, err
	}

	return &Moderator{
		collaborator: c,
	}, nil
}
