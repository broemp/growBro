package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type env struct {
	URL        string `mapstructure:"APP_URL"`
	Port       uint16 `mapstructure:"APP_PORT"`
	ConfigPath string `mapstructure:"CONFIG_PATH"`

	// Logger
	LogLevel string `mapstructure:"LOG_LEVEL"`

	// Session
	SessionSecret            string `mapstructure:"SESSION_SECRET"`
	SessionEncryptionKey     string `mapstructure:"SESSION_ENCRYPTION_KEY"`
	SessionMaxCookieAgeInMin int    `mapstructure:"SESSION_MAX_COOKIE_AGE_IN_MIN"`

	// Video
	VideoStreamPort int `mapstructure:"VIDEOSTREAM_PORT"`

	// DB
	// Postgres
	DBHost     string `mapstructure:"DB_HOST"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
	DBSSL      bool   `mapstructure:"DB_SSL"`

	// Auth
	// Local Auth
	AuthLocalEnable       bool   `mapstructure:"AUTH_LOCAL_ENABLE"`
	AuthLocalUser         string `mapstructure:"AUTH_LOCAL_USER"`
	AuthLocalPasswordHash string `mapstructure:"AUTH_LOCAL_PASSWORD_HASH"`
	AuthLocalPassword     string `mapstructure:"AUTH_LOCAL_PASSWORD_PLAIN"`

	// OIDC Auth
	AuthOidcEnable       bool   `mapstructure:"AUTH_OIDC_ENABLE"`
	AuthOidcUrl          string `mapstructure:"AUTH_OIDC_URL"`
	AuthOidcClientId     string `mapstructure:"AUTH_OIDC_CLIENT_ID"`
	AuthOidcClientSecret string `mapstructure:"AUTH_OIDC_CLIENT_SECERT"`
}

var Env env

func Init() {
	viper.SetDefault("CONFIG_PATH", "/config")
	viper.SetDefault("SESSION_MAX_COOKIE_AGE_IN_MIN", 60)
	viper.AutomaticEnv()
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.ReadInConfig()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("no .env Found. Using only Variables from the Environment")
	}

	viper.Unmarshal(&Env)
}
