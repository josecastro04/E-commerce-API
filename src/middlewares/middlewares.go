package middlewares

import (
	"api/src/authentication"
	"api/src/responses"
	"errors"
	"net/http"
)

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := authentication.ValidateToken(r); err != nil {
			responses.Erro(w, http.StatusForbidden, err)
			return
		}
		next(w, r)
	}
}

func Authorize(roletype string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		role, err := authentication.ExtractRoleFromToken(r)
		if err != nil {
			responses.Erro(w, http.StatusInternalServerError, err)
			return
		}

		if role != roletype {
			responses.Erro(w, http.StatusForbidden, errors.New("Can't acess this page"))
			return
		}
		next(w, r)
	}
}
