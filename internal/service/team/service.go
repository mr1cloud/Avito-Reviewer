package team

import (
	"context"

	"github.com/mr1cloud/Avito-Reviewer/internal/model"
)

type Team interface {
	// CreateTeam creates a new team with the given name and members
	CreateTeam(ctx context.Context, teamName string, members model.TeamMembers) (*model.Team, error)

	// UpdateTeam updates the members of an existing team
	UpdateTeam(ctx context.Context, teamName string, members model.TeamMembers) (*model.Team, error)

	// GetTeam retrieves a team by its name
	GetTeam(ctx context.Context, teamName string) (*model.Team, error)
}
