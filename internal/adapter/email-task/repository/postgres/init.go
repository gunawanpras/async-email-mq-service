package postgres

import (
	"fmt"
	"log"

	"github.com/gunawanpras/async-email-mq-service/internal/core/email-task/port"
)

func New(attr InitAttribute) port.Repository {
	if err := attr.validate(); err != nil {
		log.Panic(err)
	}

	repo := &EmailTaskRepository{
		db: attr.DB,
	}

	repo.prepareStatements()

	return repo
}

func (init InitAttribute) validate() error {
	if !init.DB.validate() {
		return fmt.Errorf("missing DB driver : %+v", init.DB)
	}

	return nil
}

func (db DB) validate() bool {
	return db.Db != nil
}
