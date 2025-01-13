package main

import (
	"log"

	"time"

	"github.com/Maksim646/space_vpx_satellite/internal/database/postgresql"
	"github.com/Maksim646/space_vpx_satellite/pkg/logger"

	//"github.com/docker/docker/daemon/logger"
	"github.com/heetch/sqalx"
	"github.com/jmoiron/sqlx"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
)

var config struct {
	LogLevel      string `envconfig:"LOG_LEVEL"`
	MigrationsDir string `envconfig:"MIGRATIONS_DIR"`
	PostgresURI   string `envconfig:"POSTGRES_URI"`
}

func main() {
	envconfig.MustProcess("", &config)

	if err := logger.BuildLogger(config.LogLevel); err != nil {
		log.Fatal("cannot build logger: ", err)
	}

	time.Sleep(5)

	migrator := postgresql.NewMigrator(config.PostgresURI, config.MigrationsDir)
	if err := migrator.Apply(); err != nil {

		log.Fatal("cannot apply migrations: ", err)
	}

	sqlxConn, err := sqlx.Connect("postgres", config.PostgresURI)
	if err != nil {
		log.Fatal("cannot connect to postgres db: ", err)
	}

	sqlxConn.SetMaxOpenConns(100)
	sqlxConn.SetMaxIdleConns(100)
	sqlxConn.SetConnMaxLifetime(5 * time.Minute)

	defer sqlxConn.Close()

	sqalxConn, err := sqalx.New(sqlxConn)
	if err != nil {
		log.Fatal("cannot connect to postgres db: ", err)
	}
	defer sqalxConn.Close()

	zap.L().Info("Database manage was process successfully")

	time.Sleep(10 * time.Minute)
}
