package teams

import (
	"errors"
	"net/http"

	"github.com/mr1cloud/Avito-Reviewer/internal/controller/rest/tools"
	"github.com/mr1cloud/Avito-Reviewer/internal/service/team"

	serviceerrors "github.com/mr1cloud/Avito-Reviewer/internal/error"

	"github.com/go-chi/chi/v5"
)

type Handlers struct {
	Router *chi.Mux
	Teams  team.Team
}

func NewTeamsHandler(teams team.Team) *Handlers {
	handlers := &Handlers{
		Router: chi.NewRouter(),
		Teams:  teams,
	}

	return handlers
}

// PostAddTeam godoc
//
//	@Summary		Add Team
//	@Description	Create a new team
//	@Tags			Teams
//
//	@Accept			json
//	@Produce		json
//	@Param			team	body		CreateUpdateTeamRequest	true	"Team to create"
//	@Success		201		{object}	model.Team
//	@Failure		400		{object}	model.ErrorResponse
//	@Failure		500		{object}	model.ErrorResponse
//
//	@Router			/team/add [post]
func (h *Handlers) PostAddTeam() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req CreateUpdateTeamRequest
		if err := tools.DecodeJSON(r, &req); err != nil {
			tools.RespondWithError(w, http.StatusBadRequest, "USER_FATAL", "Check your request body")
			return
		}

		team, err := h.Teams.CreateTeam(r.Context(), req.TeamName, req.Members)
		if err != nil {
			var srvErr serviceerrors.ServiceError
			if errors.As(err, &srvErr) {
				tools.RespondWithError(w, srvErr.ErrorStatusCode(), srvErr.Code(), srvErr.Error())
			} else {
				tools.RespondWithError(w, http.StatusInternalServerError, "FATAL", "Internal server error")
			}
			return
		}

		tools.RespondJSON(w, http.StatusCreated, team)
	}
}

// PutUpdateTeam godoc
//
//	@Summary		Update Team
//	@Description	Update an existing team
//	@Tags			Teams
//
//	@Accept			json
//	@Produce		json
//	@Param			team	body		CreateUpdateTeamRequest	true	"Team to update"
//	@Success		200		{object}	model.Team
//	@Failure		400		{object}	model.ErrorResponse
//	@Failure		500		{object}	model.ErrorResponse
//
//	@Router			/team/update [put]
func (h *Handlers) PutUpdateTeam() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req CreateUpdateTeamRequest
		if err := tools.DecodeJSON(r, &req); err != nil {
			tools.RespondWithError(w, http.StatusBadRequest, "USER_FATAL", "Check your request body")
			return
		}

		team, err := h.Teams.UpdateTeam(r.Context(), req.TeamName, req.Members)
		if err != nil {
			var srvErr serviceerrors.ServiceError
			if errors.As(err, &srvErr) {
				tools.RespondWithError(w, srvErr.ErrorStatusCode(), srvErr.Code(), srvErr.Error())
			} else {
				tools.RespondWithError(w, http.StatusInternalServerError, "FATAL", "Internal server error")
			}
			return
		}

		tools.RespondJSON(w, http.StatusOK, team)
	}
}

// GetTeam godoc
//
//	@Summary		Get Team
//	@Description	Retrieve a team by name
//	@Tags			Teams
//
//	@Produce		json
//	@Param			team_name	query		string	true	"Team Name"
//	@Success		200			{object}	model.Team
//	@Failure		400			{object}	model.ErrorResponse
//	@Failure		404			{object}	model.ErrorResponse
//	@Failure		500			{object}	model.ErrorResponse
//
//	@Router			/team/get [get]
func (h *Handlers) GetTeam() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		teamName, err := tools.GetStringQueryParam(r, "team_name", true)
		if err != nil {
			tools.RespondWithError(w, err.(serviceerrors.ServiceError).ErrorStatusCode(), "USER_FATAL", err.Error())
			return
		}

		team, err := h.Teams.GetTeam(r.Context(), teamName)
		if err != nil {
			var srvErr serviceerrors.ServiceError
			if errors.As(err, &srvErr) {
				tools.RespondWithError(w, srvErr.ErrorStatusCode(), srvErr.Code(), srvErr.Error())
			} else {
				tools.RespondWithError(w, http.StatusInternalServerError, "FATAL", "Internal server error")
			}
			return
		}

		tools.RespondJSON(w, http.StatusOK, team)
	}
}
