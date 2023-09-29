package service

import (
	"errors"
	"fmt"
	errSvc "github.com/ms-clean-code/internal/domain/error"
	"github.com/ms-clean-code/shared"
	"net/http"
)

type UserError struct {
	err  error
	args []any
}

func (us UserService) Error(err error) *UserError {
	msg, arguments := shared.GenerateError(err)
	return &UserError{
		err:  errors.New(msg),
		args: arguments,
	}
}

func (p *UserError) GetError() error {
	return p.err
}

func (p *UserError) GetHTTPCode() int {
	val, ok := errSvc.ErrorMapHttpCode[p.err.Error()]
	if !ok {
		return http.StatusInternalServerError
	}

	return val
}

func (p *UserError) GetMessage() string {
	val, ok := errSvc.ErrorMapMessage[p.err.Error()]
	if !ok {
		return errSvc.ErrorMapMessage[errSvc.InternalServerError]
	}

	return fmt.Sprintf(val, p.args...)
}

func (p *UserError) GetCaseCode() string {
	val, ok := errSvc.ErrorMapCaseCode[p.err.Error()]
	if !ok {
		return errSvc.DefaultErrorCaseCode
	}

	return val
}
