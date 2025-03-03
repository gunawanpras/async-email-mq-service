package postgres

import "github.com/jmoiron/sqlx"

type EmailTaskRepository struct {
	db        DB
	statement StatementList
}

type DB struct {
	Db *sqlx.DB
}

type StatementList struct {
	CreateEmailTask      *sqlx.Stmt
	GetEmailTaskByParams *sqlx.Stmt
	GetUserByEmail       *sqlx.Stmt
	GetEmailTaskByID     *sqlx.Stmt
}

type InitAttribute struct {
	DB DB
}
