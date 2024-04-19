package handler

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/broemp/growBro/models"
)

func render(w http.ResponseWriter, r *http.Request, comp templ.Component) error {
	return comp.Render(r.Context(), w)
}

func getAuthenticatedUser(r *http.Request) models.AuthenticatedUser {
	user, ok := r.Context().Value(models.UserContextKey).(models.AuthenticatedUser)
	if !ok {
		return models.AuthenticatedUser{}
	}
	return user
}

func hxRedirect(w http.ResponseWriter, r *http.Request, to string) error {
	if len(r.Header.Get("HX-Request")) > 0 {
		w.Header().Set("HX-Redirect", to)
		w.WriteHeader(http.StatusSeeOther)
		return nil
	}
	http.Redirect(w, r, to, http.StatusSeeOther)
	return nil
}
