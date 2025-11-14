package team

import serviceerrors "github.com/mr1cloud/Avito-Reviewer/internal/error"

// TeamAlreadyExistsError indicates that a team with the given name already exists.
type TeamAlreadyExistsError struct {
	teamName string
}

// ErrorStatusCode returns the HTTP status code for the error.
func (e *TeamAlreadyExistsError) ErrorStatusCode() int { return 400 }

// Code returns the error code.
func (e *TeamAlreadyExistsError) Code() string { return "TEAM_EXISTS" }

// Error returns the error message.
func (e *TeamAlreadyExistsError) Error() string { return e.teamName + " already exists" }

func NewTeamAlreadyExistsError(teamName string) serviceerrors.ServiceError {
	return &TeamAlreadyExistsError{
		teamName: teamName,
	}
}
