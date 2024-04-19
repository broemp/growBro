package session

import (
	"github.com/broemp/growBro/config"
	"github.com/gorilla/sessions"
)

var Store *sessions.CookieStore

func Init() {
	sessionSecret := config.Env.SessionSecret
	sessionEncryption := config.Env.SessionEncryptionKey
	cookieMaxAge := config.Env.SessionMaxCookieAgeInMin

	Store = sessions.NewCookieStore([]byte(sessionSecret), []byte(sessionEncryption))
	Store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   60 * cookieMaxAge,
		HttpOnly: true,
	}
}
