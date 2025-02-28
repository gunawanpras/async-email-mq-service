package handler

import (
	"github.com/gunawanpras/async-email-mq-service/internal/core/email-task/port"
)

type (
	EmailTaskHandler struct {
		service ServiceAttribute
	}

	ServiceAttribute struct {
		EmailTaskService port.Service
	}

	InitAttribute struct {
		Service ServiceAttribute
	}
)
