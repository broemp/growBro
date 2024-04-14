package handler

import (
	"context"
	"net/http"
	"strings"

	"github.com/broemp/cannaBro/auth"
	"github.com/broemp/cannaBro/models"
	"github.com/broemp/cannaBro/modules/session"
)

const (
	sessionCookieKey      = "session"
	sessionAccessTokenKey = "accessToken"
)

func WithUser(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "public") {
			next.ServeHTTP(w, r)
			return
		}

		store := session.Store
		session, err := store.Get(r, sessionCookieKey)
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		accessToken := session.Values[sessionAccessTokenKey]
		if accessToken == nil {
			next.ServeHTTP(w, r)
			return
		}

		ok := auth.AuthProvider.VerifyToken(accessToken.(string))
		if !ok {
			next.ServeHTTP(w, r)
			return
		}

		user := models.AuthenticatedUser{
			LoggedIn:    true,
			AccessToken: accessToken.(string),
		}
		ctx := context.WithValue(r.Context(), models.UserContextKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}

func WithAuth(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "public") {
			next.ServeHTTP(w, r)
			return
		}
		user := getAuthenticatedUser(r)
		if !user.LoggedIn {
			path := strings.Split(r.URL.Path, "/")
			http.Redirect(w, r, "/login?to="+path[1], http.StatusSeeOther)
			return
		}

		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
