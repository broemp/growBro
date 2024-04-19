package authProvider

import (
	"context"
	"encoding/base64"
	"errors"
	"math/rand/v2"
	"strconv"
	"time"

	"github.com/broemp/growBro/config"
	db "github.com/broemp/growBro/db/postgres/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
	"golang.org/x/crypto/bcrypt"
)

type LocalAuth struct {
	username string
	password string
}

func InitLocalAuth() *LocalAuth {
	localAuth := LocalAuth{
		username: config.Env.AuthLocalUser,
		password: config.Env.AuthLocalPassword,
	}
	return &localAuth
}

func (l *LocalAuth) LoginUser(username, password string) (string, error) {
	if username != l.username || password != l.password {
		return "", errors.New("wrong credentials")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(strconv.Itoa(rand.IntN(64))), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	token := base64.StdEncoding.EncodeToString(hash)
	time := pgtype.Timestamp{
		Time:  time.Now().Add(time.Hour * 24),
		Valid: true,
	}

	arg := db.CreateAccessTokenParams{
		Token:     token,
		ValidTill: time,
	}

	db.Store.CreateAccessToken(context.Background(), arg)
	return token, err
}

func (l *LocalAuth) VerifyToken(token string) bool {
	_, err := db.Store.GetToken(context.Background(), token)
	return err == nil
}
