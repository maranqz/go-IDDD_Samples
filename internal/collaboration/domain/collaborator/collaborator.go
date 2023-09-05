package collaborator

import (
	"strings"

	"github.com/maranqz/go-IDDD_Samples/internal/common/domain"
)

type Collaborator interface {
	EmailAddress() string
	Identity() domain.UUID
	Name() string
}

type collaborator struct {
	emailAddress string
	identity     domain.UUID
	name         string
}

func newCollaborator(identity domain.UUID, name string, emailAddress string) (collaborator, error) {
	return collaborator{
		emailAddress: emailAddress,
		identity:     identity,
		name:         name,
	}, nil
}

func (c *collaborator) IsEmpty() bool {
	// Use due to cannot call a pointer method on 'c.Identity()'
	identity := c.Identity()

	return c == nil || identity.IsEmpty()
}

func (c *collaborator) EmailAddress() string {
	return c.emailAddress
}

func (c *collaborator) Identity() domain.UUID {
	return c.identity
}

func (c *collaborator) Name() string {
	return c.name
}

func (c *collaborator) CompareTo(aCollaborator *collaborator) int {
	diff := strings.Compare(c.Identity().String(), aCollaborator.Identity().String())
	if diff == 0 {
		diff = strings.Compare(c.EmailAddress(), aCollaborator.EmailAddress())
		if diff == 0 {
			diff = strings.Compare(c.Name(), aCollaborator.Name())
		}
	}

	return diff
}
