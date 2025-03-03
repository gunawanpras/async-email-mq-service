package constant

import (
	"net/http"
)

const (
	SUCCESS = "success"
	ERROR   = "error"
)

const (
	SYSTEM = "SYSTEM"
)

const (
	EmailTaskQueueName = "email_queue"
)

const (
	EmailTaskStatusPending = "PENDING"
	EmailTaskStatusSent    = "SENT"
	EmailTaskStatusFailed  = "FAILED"
)

const (
	BindingParameterFailed = "failed to bind parameter"
)

const (
	EmailTaskCreateSuccess = "email task created successfully"
	EmailTaskCreateFailed  = "failed to create email task"
	EmailTaskGetSuccess    = "email task fetched successfully"
	EmailTaskGetFailed     = "failed to fetch email task"
	EmailTaskNotFound      = "email task not found"
	EmailTaskAlreadyExist  = "email task already exist"
	SenderEmailNotFound    = "sender email not found"
	EmailTaskSentSuccess   = "email task sent successfully"
	EmailTaskSentFailed    = "failed to send email task"
)

const (
	DbBeginTransactionFailed    = "failed to begin transaction: %v"
	DbRollbackTransactionFailed = "failed to rollback transaction: %v"
	DbCommitTransactionFailed   = "failed to commit transaction: %v"
	DataNotFound                = "data not found"
	DbReturnedMalformedData     = "database returned malformed data"
)

var (
	GenericHttpStatusMappings = map[string]int{
		BindingParameterFailed: http.StatusBadRequest,
	}

	EmailTaskHttpStatusMappings = map[string]int{
		EmailTaskCreateSuccess:      http.StatusCreated,
		EmailTaskCreateFailed:       http.StatusInternalServerError,
		EmailTaskGetSuccess:         http.StatusOK,
		EmailTaskGetFailed:          http.StatusInternalServerError,
		EmailTaskAlreadyExist:       http.StatusConflict,
		EmailTaskNotFound:           http.StatusNotFound,
		DataNotFound:                http.StatusNotFound,
		SenderEmailNotFound:         http.StatusNotFound,
		EmailTaskSentSuccess:        http.StatusOK,
		EmailTaskSentFailed:         http.StatusInternalServerError,
		DbBeginTransactionFailed:    http.StatusInternalServerError,
		DbRollbackTransactionFailed: http.StatusInternalServerError,
		DbCommitTransactionFailed:   http.StatusInternalServerError,
		DbReturnedMalformedData:     http.StatusInternalServerError,
	}
)
