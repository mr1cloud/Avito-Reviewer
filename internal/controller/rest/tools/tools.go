package tools

import (
	"encoding/json"
	"net/http"

	"github.com/mr1cloud/Avito-Reviewer/internal/model"
)

func DecodeJSON(r *http.Request, v interface{}) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	return decoder.Decode(v)
}

func RespondJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func RespondWithError(w http.ResponseWriter, code int, errorCode, message string) {
	RespondJSON(w, code, model.ErrorResponse{
		Error: model.Error{
			Code:    errorCode,
			Message: message,
		},
	})
}
