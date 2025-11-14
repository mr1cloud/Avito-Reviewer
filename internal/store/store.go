package store

import (
	"github.com/mr1cloud/Avito-Reviewer/internal/repository/teams"
	"github.com/mr1cloud/Avito-Reviewer/internal/repository/users"
)

// Store is interface for accessing main repositories
type Store interface {
	// UsersRepository returns users repository
	UsersRepository() users.UsersRepository

	// TeamsRepository returns teams repository
	TeamsRepository() teams.TeamsRepository

	// Close closes the store
	Close() error
}
