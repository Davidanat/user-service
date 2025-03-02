package error

import "errors"

var (
	ErrInternalServerError = errors.New("internal server error")
	ErrSQLError            = errors.New("database server failed to execute query")
	ErrUnauthorized        = errors.New("unauthorized")
	ErrInvalidToken        = errors.New("invalid token")
	ErrForbidden           = errors.New("forbidden")
	ErrTooManyRequest      = errors.New("too many request")
)

var GeneralErrors = []error{
	ErrInternalServerError,
	ErrSQLError,
	ErrUnauthorized,
	ErrInvalidToken,
	ErrForbidden,
	ErrTooManyRequest,
}
