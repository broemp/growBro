package handler

import (
	"net/http"

	"github.com/broemp/cannaBro/auth"
	"github.com/broemp/cannaBro/modules/session"
	"github.com/broemp/cannaBro/view/login"
	"go.uber.org/zap"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) error {
	return render(w, r, login.Index())
}

func HandleLoginForm(w http.ResponseWriter, r *http.Request) error {
	username := r.FormValue("username")
	password := r.FormValue("password")

	token, err := auth.AuthProvider.LoginUser(username, password)
	if err != nil {
		loginErr := login.LoginErrors{
			Username:          username,
			InvalidCredenials: "Wrong Credentials",
		}

		zap.L().Info("Failed Login Attempt", zap.String("username", username), zap.String("ip", r.RemoteAddr))
		return render(w, r, login.LoginForm(loginErr))
	}
	zap.L().Info("Logged in User", zap.String("username", username), zap.String("ip", r.RemoteAddr))
	setAuthCookie(w, r, token)
	hxRedirect(w, r, "/")
	return nil
}

func HandleLogout(w http.ResponseWriter, r *http.Request) error {
	store := session.Store
	session, _ := store.Get(r, sessionCookieKey)
	session.Values[sessionAccessTokenKey] = ""
	session.Save(r, w)

	http.Redirect(w, r, "/login", http.StatusSeeOther)
	return nil
}

func setAuthCookie(w http.ResponseWriter, r *http.Request, accessToken string) error {
	store := session.Store
	session, _ := store.Get(r, sessionCookieKey)
	session.Values[sessionAccessTokenKey] = accessToken
	return session.Save(r, w)
}
