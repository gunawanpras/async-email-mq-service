package postgres

import (
	"time"

	"github.com/google/uuid"
	"github.com/gunawanpras/async-email-mq-service/internal/core/email-task/domain"
)

type EmailTask struct {
	ID             uuid.UUID  `db:"id"`
	SenderEmail    string     `db:"sender_email"`
	RecipientEmail string     `db:"recipient_email"`
	Subject        string     `db:"subject"`
	Body           string     `db:"body"`
	Status         string     `db:"status"`
	RetryCount     int        `db:"retry_count"`
	CreatedAt      time.Time  `db:"created_at"`
	CreatedBy      string     `db:"created_by"`
	UpdatedAt      *time.Time `db:"updated_at"`
	UpdatedBy      *string    `db:"updated_by"`
}

func (u EmailTask) Validate() bool {
	if u.ID == uuid.Nil {
		return false
	}

	if u.SenderEmail == "" {
		return false
	}

	if u.RecipientEmail == "" {
		return false
	}

	if u.Subject == "" {
		return false
	}

	if u.Body == "" {
		return false
	}

	if u.Status == "" {
		return false
	}

	if u.RetryCount < 0 {
		return false
	}

	if u.CreatedAt.IsZero() {
		return false
	}

	if u.CreatedBy == "" {
		return false
	}

	if u.UpdatedAt != nil && u.UpdatedAt.IsZero() {
		return false
	}

	if u.UpdatedBy != nil && *u.UpdatedBy == "" {
		return false
	}

	return true
}

func (u EmailTask) ToModel() domain.EmailTask {
	return domain.EmailTask{
		ID:             u.ID,
		SenderEmail:    u.SenderEmail,
		RecipientEmail: u.RecipientEmail,
		Subject:        u.Subject,
		Body:           u.Body,
		Status:         u.Status,
		RetryCount:     u.RetryCount,
		CreatedAt:      u.CreatedAt,
		CreatedBy:      u.CreatedBy,
		UpdatedAt:      u.UpdatedAt,
		UpdatedBy:      u.UpdatedBy,
	}
}

type User struct {
	ID        uuid.UUID  `db:"id"`
	Name      string     `db:"name"`
	Email     string     `db:"email"`
	Password  string     `db:"password"`
	RoleID    *uuid.UUID `db:"role_id"`
	CreatedAt time.Time  `db:"created_at"`
	CreatedBy string     `db:"created_by"`
	UpdatedAt *time.Time `db:"updated_at"`
	UpdatedBy *string    `db:"updated_by"`
}

func (u User) Validate() bool {
	if u.ID == uuid.Nil {
		return false
	}

	if u.Name == "" {
		return false
	}

	if u.Email == "" {
		return false
	}

	if u.Password == "" {
		return false
	}

	if u.RoleID != nil && *u.RoleID == uuid.Nil {
		return false
	}

	if u.CreatedAt.IsZero() {
		return false
	}

	if u.CreatedBy == "" {
		return false
	}

	if u.UpdatedAt != nil && u.UpdatedAt.IsZero() {
		return false
	}

	if u.UpdatedBy != nil && *u.UpdatedBy == "" {
		return false
	}

	return true
}

func (u User) ToModel() domain.User {
	return domain.User{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		Password:  u.Password,
		RoleID:    u.RoleID,
		CreatedAt: u.CreatedAt,
		CreatedBy: u.CreatedBy,
	}
}
