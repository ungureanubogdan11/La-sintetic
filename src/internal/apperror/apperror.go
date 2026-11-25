package apperror

import "errors"

type Kind string

const (
	KindInvalid      Kind = "invalid_input"  // 400
	KindUnauthorized Kind = "unauthorized"   // 401
	KindForbidden    Kind = "forbidden"      // 403
	KindNotFound     Kind = "not_found"      // 404
	KindConflict     Kind = "conflict"       // 409
	KindInternal     Kind = "internal_error" // 500
)

type Error struct {
	Kind    Kind
	Message string
	Details any
	Err     error
}

func (e *Error) Error() string {
	if e.Err != nil {
		return e.Message + ": " + e.Err.Error()
	}
	return e.Message
}

func (e *Error) Unwrap() error { return e.Err }

func New(kind Kind, msg string) *Error { return &Error{Kind: kind, Message: msg} }

func Invalid(msg string, details any) *Error {
	return &Error{Kind: KindInvalid, Message: msg, Details: details}
}
func Unauthorized(msg string) *Error { return New(KindUnauthorized, msg) }
func Forbidden(msg string) *Error    { return New(KindForbidden, msg) }
func NotFound(msg string) *Error     { return New(KindNotFound, msg) }
func Conflict(msg string) *Error     { return New(KindConflict, msg) }
func Internal(err error) *Error {
	return &Error{Kind: KindInternal, Message: "internal server error", Err: err}
}

func As(err error) *Error {
	var e *Error
	if errors.As(err, &e) {
		return e
	}
	return Internal(err)
}
