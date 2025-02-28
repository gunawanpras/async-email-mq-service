package dto

type CreateEmailTaskRequest struct {
	SenderEmail    string `json:"sender_email" validate:"required,email"`
	RecipientEmail string `json:"recipient_email" validate:"required,email"`
	Subject        string `json:"subject" validate:"required"`
	Body           string `json:"body" validate:"required"`
}
