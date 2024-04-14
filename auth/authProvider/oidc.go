package authProvider

import (
	"context"
	"errors"
	"fmt"

	"github.com/broemp/cannaBro/config"
	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

type OIDC struct {
	URL string
}

func (*OIDC) LoginUser(username, password string) string {
	return ""
}

func (*OIDC) VerifyToken(token string) bool {
	return true
}

// TODO: Implement OIDC Login
func InitOIDCAuth() (AuthProvider, error) {
	oidcURL := config.Env.AuthOidcUrl
	clientID := config.Env.AuthOidcClientId
	clientSecret := config.Env.AuthOidcClientSecret
	redirectURL := config.Env.URL

	// check if required Variables are set
	if len(clientID) == 0 || len(clientSecret) == 0 {
		return nil, errors.New("OICD Client ID or Secret is not set")
	}

	provider, err := oidc.NewProvider(context.Background(), oidcURL)
	if err != nil {
		return nil, err
	}

	// Configure an OpenID Connect aware OAuth2 client.
	auth2Config := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,

		// Discovery returns the OAuth2 endpoints.
		Endpoint: provider.Endpoint(),

		// "openid" is a required scope for OpenID Connect flows.
		Scopes: []string{oidc.ScopeOpenID, "profile", "email"},
	}

	fmt.Printf("auth2Config: %v\n", auth2Config)

	var authProvider AuthProvider

	return authProvider, nil
}
