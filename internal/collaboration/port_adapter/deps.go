package port_adapter

import (
	"github.com/maranqz/go-IDDD_Samples/internal/collaboration/domain/collaborator"
	"github.com/maranqz/go-IDDD_Samples/internal/collaboration/domain/tenant"
	"github.com/maranqz/go-IDDD_Samples/internal/common/domain"
)

type UserInRoleAdapter interface {
	ToCollaborator(aTenant tenant.ID, anIdentity domain.UUID, aRoleName string) (collaborator.Collaborator, error)
}
