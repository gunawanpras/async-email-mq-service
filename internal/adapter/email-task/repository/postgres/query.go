package postgres

var (
	queryCreateEmailTask = `
		INSERT INTO email_tasks (id, sender_email, recipient_email, subject, body, status, created_at, created_by)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	queryGetUserByEmail = `
		SELECT id, name, email, password, role_id, created_at, created_by
		FROM users
		WHERE email = $1
	`

	queryGetEmailTaskByParams = `
		SELECT id, sender_email, recipient_email, subject, body, status, retry_count, created_at, created_by, updated_at, updated_by
		FROM email_tasks
		WHERE sender_email = $1 AND recipient_email = $2 AND subject = $3 AND body = $4
	`
)
