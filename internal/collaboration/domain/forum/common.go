package forum

import (
	"errors"
	"fmt"

	"github.com/maranqz/go-IDDD_Samples/internal/collaboration/domain/collaborator"
)

var (
	ErrAuthorEmpty = errors.New("author: empty")
)

// TODO think about it
func assertAuthor(a *collaborator.Author, p error) error {
	if a.IsEmpty() {
		return fmt.Errorf("%w: %w", p, ErrAuthorEmpty)
	}

	return nil
}
