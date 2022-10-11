package json

import (
	"errors"

	"github.com/wascript3r/httputil/json/errcode"
)

var (
	NotFoundError = errcode.New(
		"not_found",
		errors.New("not found"),
	)

	ForbiddenError = errcode.New(
		"forbidden",
		errors.New("forbidden"),
	)

	UnauthorizedError = errcode.New(
		"unauthorized",
		errors.New("unauthorized"),
	)

	BadRequestError = errcode.New(
		"bad_request",
		errors.New("bad request"),
	)

	InternalServerError = errcode.New(
		"internal_server_error",
		errors.New("internal server error"),
	)
)
