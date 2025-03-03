package dto

import (
	"github.com/google/uuid"
)

type CreateEmailTaskResponse struct {
	ID uuid.UUID `json:"id"`
}
