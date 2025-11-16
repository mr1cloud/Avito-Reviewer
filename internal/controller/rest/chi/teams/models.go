package teams

import "github.com/mr1cloud/Avito-Reviewer/internal/model"

type CreateUpdateTeamRequest struct {
	TeamName string            `json:"team_name" validate:"required,max=64"`
	Members  model.TeamMembers `json:"members" validate:"required,dive"`
}

type TeamResponse struct {
	Team model.Team `json:"team"`
}
