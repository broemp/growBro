package models

const UserContextKey = "user"

type AuthenticatedUser struct {
	Username    string
	LoggedIn    bool
	AccessToken string
}
