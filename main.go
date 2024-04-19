package main

import (
	"embed"

	"github.com/broemp/growBro/auth"
	"github.com/broemp/growBro/config"
	db "github.com/broemp/growBro/db/postgres/sqlc"
	"github.com/broemp/growBro/logger"
	"github.com/broemp/growBro/server"
	"github.com/broemp/growBro/session"
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
	println("Running")
	server.Start(publicFS)
}
