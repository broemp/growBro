package main

import (
	"embed"

	"github.com/broemp/cannaBro/auth"
	"github.com/broemp/cannaBro/config"
	db "github.com/broemp/cannaBro/db/postgres/sqlc"
	"github.com/broemp/cannaBro/logger"
	"github.com/broemp/cannaBro/modules/session"
	"github.com/broemp/cannaBro/server"
)

//go:embed public
var publicFS embed.FS

//go:embed db/postgres/schema
var schemaFS embed.FS

func main() {
	config.Init()
	logger.Init()
	session.Init()
	db.Init(schemaFS)
	auth.Init()
	server.Start(publicFS)
}
