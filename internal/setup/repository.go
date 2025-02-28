package setup

import (
	emailTaskRepoPg "github.com/gunawanpras/async-email-mq-service/internal/adapter/email-task/repository/postgres"
	emailTaskRepo "github.com/gunawanpras/async-email-mq-service/internal/core/email-task/port"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	EmailTaskRepo emailTaskRepo.Repository
}

func NewRepository(db *sqlx.DB) Repository {
	emailTaskRepo := emailTaskRepoPg.New(emailTaskRepoPg.InitAttribute{
		DB: emailTaskRepoPg.DB{
			Db: db,
		},
	})

	return Repository{
		EmailTaskRepo: emailTaskRepo,
	}
}
