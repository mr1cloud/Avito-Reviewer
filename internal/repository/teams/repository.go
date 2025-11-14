package teams

import (
	"context"

	"github.com/mr1cloud/Avito-Reviewer/internal/model"
)

// TeamsRepository stores teams data
//
//goland:noinspection GoNameStartsWithPackageName
type TeamsRepository interface {
	// CreateTeam creates a new team with the given name and members
	CreateTeam(ctx context.Context, teamName string, members model.TeamMembers) error
	// UpdateTeam updates the members of an existing team
	UpdateTeam(ctx context.Context, teamName string, oldMembers model.TeamMembers, newMembers model.TeamMembers) error
	// GetTeam retrieves a team by its name
	GetTeam(ctx context.Context, teamName string) (*model.Team, error)
}
