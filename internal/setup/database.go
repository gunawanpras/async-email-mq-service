package setup

import (
	"log"
	"time"

	"github.com/gunawanpras/async-email-mq-service/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func InitPostgres(conf *config.Config) *sqlx.DB {
	pg, err := sqlx.Connect("postgres", conf.Postgres.Primary.ConnString)
	if err != nil {
		log.Panic("failed to open postgre client for project-mq service:", err)
	}

	if err := pg.Ping(); err != nil {
		log.Panic("failed to ping PostgreSQL: %w", err)
	}

	pg.SetMaxOpenConns(conf.Postgres.Primary.MaxOpenConn)
	pg.SetMaxIdleConns(conf.Postgres.Primary.MaxIdleConn)
	pg.SetConnMaxLifetime(time.Second * time.Duration(conf.Postgres.Primary.MaxConnLifeTimeInSecond))

	return pg
}
