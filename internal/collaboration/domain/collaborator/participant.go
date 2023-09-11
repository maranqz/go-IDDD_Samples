package collaborator

import (
	"errors"
	"fmt"

	"github.com/maranqz/go-IDDD_Samples/internal/common/domain"
)

var ErrFromCollaborator = errors.New("create from collaborator")
var (
	ErrParticipant                 = errors.New("participant")
	ErrParticipantFromCollaborator = fmt.Errorf("%w: %w", ErrParticipant, ErrFromCollaborator)
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

func ParticipantFromCollaborator(collaborator Collaborator) (*Participant, error) {
	c, ok := collaborator.(*Participant)
	if !ok {
		return nil, ErrParticipantFromCollaborator
	}

	return c, nil
}

var (
	ErrOwner                 = errors.New("owner")
	ErrOwnerFromCollaborator = fmt.Errorf("%w: %w", ErrOwner, ErrFromCollaborator)
)

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

func OwnerFromCollaborator(collaborator Collaborator) (*Owner, error) {
	c, ok := collaborator.(*Owner)
	if !ok {
		return nil, ErrOwnerFromCollaborator
	}

	return c, nil
}

var (
	ErrAuthor                 = errors.New("author")
	ErrAuthorFromCollaborator = fmt.Errorf("%w: %w", ErrAuthor, ErrFromCollaborator)
)

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

func AuthorFromCollaborator(collaborator Collaborator) (*Author, error) {
	c, ok := collaborator.(*Author)
	if !ok {
		return nil, ErrAuthorFromCollaborator
	}

	return c, nil
}

var (
	ErrCreator                 = errors.New("creator")
	ErrCreatorFromCollaborator = fmt.Errorf("%w: %w", ErrCreator, ErrFromCollaborator)
)

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

func CreatorFromCollaborator(collaborator Collaborator) (*Creator, error) {
	c, ok := collaborator.(*Creator)
	if !ok {
		return nil, ErrCreatorFromCollaborator
	}

	return c, nil
}

var (
	ErrModerator                 = errors.New("moderator")
	ErrModeratorFromCollaborator = fmt.Errorf("%w: %w", ErrModerator, ErrFromCollaborator)
)

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

func ModeratorFromCollaborator(collaborator Collaborator) (*Moderator, error) {
	c, ok := collaborator.(*Moderator)
	if !ok {
		return nil, ErrModeratorFromCollaborator
	}

	return c, nil
}
