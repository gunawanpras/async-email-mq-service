package handler

import (
	"github.com/gunawanpras/async-email-mq-service/internal/core/email-task/port"
)

type (
	ServiceAttribute struct {
		EmailTaskService port.Service
	}

	EmailTaskHandler struct {
		service ServiceAttribute
	}

	InitAttribute struct {
		Service ServiceAttribute
	}
)
