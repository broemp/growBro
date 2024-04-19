package db

import (
	"context"
	"embed"
	"fmt"
	"log"
	"math"
	"strings"
	"time"

	"github.com/broemp/growBro/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/pgx/v5"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

var Store DBStore

func ConnectDatabase(
	dbname,
	dbuser,
	dbpassword,
	dbhost,
	sslmode string,
	ctx context.Context,
) *pgxpool.Pool {
	hostArr := strings.Split(dbhost, ":")
	host := hostArr[0]
	port := "5432"
	if len(hostArr) > 1 {
		port = hostArr[1]
	}
	uri := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		dbuser,
		dbpassword,
		dbname,
		host,
		port,
		sslmode,
	)

	pgxConfig, err := pgxpool.ParseConfig(uri)
	if err != nil {
		panic(err)
	}

	var connPool *pgxpool.Pool
	i := 1
	for {
		connPool, err = pgxpool.NewWithConfig(ctx, pgxConfig)
		timeout := math.Pow(2, float64(i))
		msg := fmt.Sprintf("couldn't connect to databse! Retrying in %d Seconds", int(timeout))
		if err := connPool.Ping(context.Background()); err == nil {
			break
		}
		zap.L().Error(msg, zap.Error(err))
		time.Sleep(time.Duration(timeout) * time.Second)
		i++
	}
	return connPool
}

func Init(fs embed.FS) {
	var (
		host    = config.Env.DBHost
		user    = config.Env.DBUser
		pass    = config.Env.DBPassword
		dbname  = config.Env.DBName
		sslmode = "disable"
	)

	if config.Env.DBSSL {
		sslmode = "require"
	}

	err := migrateDB(fs)
	if err != nil {
		zap.L().Info("Migrate DB:", zap.Error(err))
	}
	conn := ConnectDatabase(dbname, user, pass, host, sslmode, context.Background())

	Store = NewStore(conn)
	log.Println("Connected to DB")
	zap.L().Info("connected to db")
}

func migrateDB(fs embed.FS) error {
	var (
		host    = config.Env.DBHost
		user    = config.Env.DBUser
		pass    = config.Env.DBPassword
		dbname  = config.Env.DBName
		sslmode = "disable"
	)

	if config.Env.DBSSL {
		sslmode = "require"
	}

	d, err := iofs.New(fs, "db/postgres/schema")
	if err != nil {
		zap.L().Error("could not create iofs", zap.Error(err))
	}
	m, err := migrate.NewWithSourceInstance("iofs", d, fmt.Sprintf("pgx5://%s:%s@%s/%s?sslmode=%s", user, pass, host, dbname, sslmode))
	if err != nil {
		zap.L().Error("could not create new migration", zap.Error(err))
	}
	err = m.Up()
	if err != nil {
		zap.L().Error("could not upgrade database", zap.Error(err))
	}

	return nil
}
