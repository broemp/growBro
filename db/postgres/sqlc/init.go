package db

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"log"
	"math"
	"strings"
	"time"

	"github.com/broemp/growBro/config"
	_ "github.com/golang-migrate/migrate/v4/database/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
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

	fmt.Println(uri)

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

	conn := ConnectDatabase(dbname, user, pass, host, sslmode, context.Background())
	stdDB := stdlib.OpenDBFromPool(conn)
	err := migrateDB(fs, stdDB)
	if err != nil {
		zap.L().Info("Migrate DB:", zap.Error(err))
	}

	Store = NewStore(conn)
	log.Println("Connected to DB")
	zap.L().Info("connected to db")
}

func migrateDB(fs embed.FS, db *sql.DB) error {
	goose.SetBaseFS(fs)

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	if err := goose.Up(db, "db/postgres/schema"); err != nil {
		panic(err)
	}

	return nil
}
