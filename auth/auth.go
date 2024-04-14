package auth

import (
	"log"
	"os"

	"github.com/broemp/cannaBro/auth/authProvider"
	"github.com/broemp/cannaBro/config"
	"go.uber.org/zap"
)

var AuthProvider authProvider.AuthProvider

// Varibale to check if Auth is active from outside of package
var Enabled bool = false

func Init() {
	zap.L().Info("Auth")
	var err error

	if config.Env.AuthLocalEnable && config.Env.AuthOidcEnable {
		zap.L().Error("Only enable either Local Auth or OIDC Auth")
		os.Exit(0)

	} else if config.Env.AuthLocalEnable {
		zap.L().Info("Using Local Auth")

		Enabled = true
		AuthProvider = authProvider.InitLocalAuth()
	} else if config.Env.AuthOidcEnable {
		zap.L().Info("Using OIDC Auth")

		AuthProvider, err = authProvider.InitOIDCAuth()
		if err != nil {
			log.Fatal(err)
		}

	}
	Enabled = true
}
