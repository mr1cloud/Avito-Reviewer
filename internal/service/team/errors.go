package team

import (
	"net/http"

	serviceerrors "github.com/mr1cloud/Avito-Reviewer/internal/error"
)

// TeamAlreadyExistsError indicates that a team with the given name already exists.
type TeamAlreadyExistsError struct {
	teamName string
}

func (e *TeamAlreadyExistsError) ErrorStatusCode() int { return http.StatusConflict }

func (e *TeamAlreadyExistsError) Code() string { return "TEAM_EXISTS" }

func (e *TeamAlreadyExistsError) Error() string { return e.teamName + " already exists" }

func NewTeamAlreadyExistsError(teamName string) serviceerrors.ServiceError {
	return &TeamAlreadyExistsError{
		teamName: teamName,
	}
}
