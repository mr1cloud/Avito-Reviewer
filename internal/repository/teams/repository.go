package teams

import (
	"context"

	"github.com/mr1cloud/Avito-Reviewer/internal/model"
)

// TeamsRepository stores teams data
//
//goland:noinspection GoNameStartsWithPackageName
type TeamsRepository interface {
	// InsertTeam inserts a new team with the given name and members
	InsertTeam(ctx context.Context, teamName string, members model.TeamMembers) error
	// UpdateTeam updates the members of an existing team
	UpdateTeam(ctx context.Context, teamName string, oldMembers model.TeamMembers, newMembers model.TeamMembers) error
	// SelectTeam retrieves a team by its name
	SelectTeam(ctx context.Context, teamName string) (*model.Team, error)
}
