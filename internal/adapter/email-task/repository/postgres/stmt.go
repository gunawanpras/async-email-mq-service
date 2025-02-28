package postgres

import (
	"log"

	"github.com/jmoiron/sqlx"
)

func (repo *EmailTaskRepository) prepareStatements() {
	repo.statement = StatementList{}
}

func (repo *EmailTaskRepository) prepareCreateEmailTask() {
	var (
		err  error
		stmt *sqlx.Stmt
		db   = repo.db.Db
	)

	if stmt, err = db.Preparex(queryCreateEmailTask); err != nil {
		log.Panic("[prepareCreateEmailTask] error:", err)
	}
	repo.statement.CreateEmailTask = stmt
}

func (repo *EmailTaskRepository) prepareGetEmailTaskByParams() {
	var (
		err  error
		stmt *sqlx.Stmt
		db   = repo.db.Db
	)

	if stmt, err = db.Preparex(queryGetEmailTaskByParams); err != nil {
		log.Panic("[prepareGetEmailTaskByParams] error:", err)
	}
	repo.statement.GetEmailTaskByParams = stmt
}

func (repo *EmailTaskRepository) prepareGetUserByEmail() {
	var (
		err  error
		stmt *sqlx.Stmt
		db   = repo.db.Db
	)

	if stmt, err = db.Preparex(queryGetUserByEmail); err != nil {
		log.Panic("[prepareGetUserByEmail] error:", err)
	}
	repo.statement.GetUserByEmail = stmt
}
