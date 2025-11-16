package users

import (
	"errors"
	"net/http"

	"github.com/mr1cloud/Avito-Reviewer/internal/controller/rest/tools"
	"github.com/mr1cloud/Avito-Reviewer/internal/service/pull-request"
	"github.com/mr1cloud/Avito-Reviewer/internal/service/user"

	serviceerrors "github.com/mr1cloud/Avito-Reviewer/internal/error"

	"github.com/go-chi/chi/v5"
)

type Handlers struct {
	Router       *chi.Mux
	Users        user.User
	PullRequests pull_request.PullRequest
}

func NewUsersHandler(users user.User, pullRequests pull_request.PullRequest) *Handlers {
	handlers := &Handlers{
		Router:       chi.NewRouter(),
		Users:        users,
		PullRequests: pullRequests,
	}

	return handlers
}

// PostSetUserIsActive godoc
//
//	@Summary		Set user is active
//	@Description	Set the active status of a user
//	@Tags			Users
//
//	@Accept			json
//	@Produce		json
//	@Param			request	body		SetUserIsActiveRequest	true	"Set User Is Active Request"
//	@Success		200		{object}	model.User
//	@Failure		400		{object}	model.ErrorResponse
//	@Failure		404		{object}	model.ErrorResponse
//	@Failure		500		{object}	model.ErrorResponse
//
//	@Router			/users/setIsActive [post]
func (h *Handlers) PostSetUserIsActive() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req SetUserIsActiveRequest
		if err := tools.DecodeJSON(r, &req); err != nil {
			tools.RespondWithError(w, http.StatusBadRequest, "USER_FATAL", "Check your request body")
			return
		}

		user, err := h.Users.UpdateUserIsActiveById(r.Context(), req.UserID, req.IsActive)
		if err != nil {
			var srvErr serviceerrors.ServiceError
			if errors.As(err, &srvErr) {
				tools.RespondWithError(w, srvErr.ErrorStatusCode(), srvErr.Code(), srvErr.Error())
			} else {
				tools.RespondWithError(w, http.StatusInternalServerError, "FATAL", "Internal server error")
			}
			return
		}

		tools.RespondJSON(w, http.StatusOK, user)
	}
}

// GetPullRequestsAssignedToUser godoc
//
//	@Summary		Get pull requests assigned to user
//	@Description	Retrieve pull requests assigned to a specific user
//	@Tags			Users
//
//	@Produce		json
//	@Param			user_id	query		string	true	"User ID"
//	@Success		200		{object}	PullRequestsAssignedToUserResponse
//	@Failure		400		{object}	model.ErrorResponse
//	@Failure		404		{object}	model.ErrorResponse
//	@Failure		500		{object}	model.ErrorResponse
//
//	@Router			/users/getReview [get]
func (h *Handlers) GetPullRequestsAssignedToUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId, err := tools.GetStringQueryParam(r, "user_id", true)
		if err != nil {
			tools.RespondWithError(w, err.(serviceerrors.ServiceError).ErrorStatusCode(), "USER_FATAL", err.Error())
			return
		}

		prs, err := h.PullRequests.GetPullRequestsAssignedToUser(r.Context(), userId)
		if err != nil {
			var srvErr serviceerrors.ServiceError
			if errors.As(err, &srvErr) {
				tools.RespondWithError(w, srvErr.ErrorStatusCode(), srvErr.Code(), srvErr.Error())
			} else {
				tools.RespondWithError(w, http.StatusInternalServerError, "FATAL", "Internal server error")
			}
			return
		}

		tools.RespondJSON(w, http.StatusOK,
			PullRequestsAssignedToUserResponse{
				UserID:       userId,
				PullRequests: prs,
			},
		)
	}
}
