package tools

import (
	"errors"
	"net/http"
)

var ErrMissingQueryParameter = errors.New("missing query parameter")

func GetStringQueryParam(r *http.Request, key string, required bool) (string, error) {
	value := r.URL.Query().Get(key)
	if value == "" {
		if required {
			return "", ErrMissingQueryParameter
		}
		return "", nil
	}

	return value, nil
}
