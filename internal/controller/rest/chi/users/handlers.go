package users

import (
	"errors"
	"net/http"

	"github.com/mr1cloud/Avito-Reviewer/internal/controller/rest/tools"
	"github.com/mr1cloud/Avito-Reviewer/internal/service/user"

	serviceerrors "github.com/mr1cloud/Avito-Reviewer/internal/error"

	"github.com/go-chi/chi/v5"
)

type Handlers struct {
	Router *chi.Mux
	Users  user.User
}

func NewUsersHandler(users user.User) *Handlers {
	handlers := &Handlers{
		Router: chi.NewRouter(),
		Users:  users,
	}

	return handlers
}

// PostSetUserIsActive godoc
//
//	@Summary		Set User Is Active
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
