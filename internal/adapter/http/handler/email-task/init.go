package handler

import (
	"errors"
)

func New(attr InitAttribute) Handler {
	if err := attr.validate(); err != nil {
		panic(err)
	}

	return &EmailTaskHandler{
		service: attr.Service,
	}
}

func (attr InitAttribute) validate() error {
	if attr.Service.EmailTaskService == nil {
		return errors.New("missing email-task service")
	}

	return nil
}
