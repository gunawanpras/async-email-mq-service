package domain

import (
	"time"

	"github.com/google/uuid"
)

type EmailTask struct {
	ID             uuid.UUID
	SenderEmail    string
	RecipientEmail string
	Subject        string
	Body           string
	Status         string
	RetryCount     int
	CreatedAt      time.Time
	CreatedBy      string
	UpdatedAt      *time.Time
	UpdatedBy      *string
}

type User struct {
	ID        uuid.UUID
	Name      string
	Email     string
	Password  string
	RoleID    *uuid.UUID
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt *time.Time
	UpdatedBy *string
}
