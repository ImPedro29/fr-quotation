package services

import (
	"fmt"
	"github.com/ImPedro29/fr-quotation/utils"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
	"go.uber.org/zap"
	"log"
)

var postgresInstance *sqlx.DB
var mock sqlmock.Sqlmock

func Postgres() *sqlx.DB {
	if postgresInstance == nil {
		connectionURL := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
			utils.Env.PostgresHost,
			utils.Env.PostgresUser,
			utils.Env.PostgresPassword,
			utils.Env.PostgresDatabase,
			utils.Env.PostgresPort,
			utils.Env.PostgresSSLMode,
		)

		db, err := sqlx.Connect("postgres", connectionURL)
		if err != nil {
			log.Fatalln(err)
		}

		if err := db.Ping(); err != nil {
			zap.L().Panic("failed to ping DB", zap.Error(err))
		}

		postgresInstance = db
	}

	return postgresInstance
}

func Mock() sqlmock.Sqlmock {
	return mock
}

func ActivateMock() {
	if mock == nil {
		db, mocked, err := sqlmock.Newx()
		if err != nil {
			log.Fatalln(err)
		}

		mock = mocked
		postgresInstance = db
	}
}
