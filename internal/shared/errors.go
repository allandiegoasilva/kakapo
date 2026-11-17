package shared

import "errors"

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ErrorType string

var (
	NotFound           = errors.New("NOT_FOUND")
	InvalidField       = errors.New("INVALID_FIELD")
	InternalServer     = errors.New("INTERNAL_ERROR")
	Unauthorized       = errors.New("UNAUTHORIZED")
	BadRequest         = errors.New("BAD_REQUEST")
	ServiceUnavailable = errors.New("SERVICE_UNAVAILABLE")
)

type AppError struct {
	Message string
	Fields  []FieldError
	Err     error
	Cause   error
}

type ErrorInput struct {
	Message string
	Fields  []FieldError
	Cause   error
}

func (e *AppError) Error() string {
	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func ErrorCause(err error) error {
	var appErr *AppError
	if errors.As(err, &appErr) && appErr.Cause != nil {
		return appErr.Cause
	}
	return err
}

func ValidationError(input ErrorInput) error {
	return &AppError{
		Message: input.Message,
		Fields:  input.Fields,
		Err:     BadRequest,
	}
}

func UnauthorizedError(input ErrorInput) error {
	return &AppError{
		Message: input.Message,
		Fields:  input.Fields,
		Err:     Unauthorized,
	}
}

func NotFoundError(input ErrorInput) error {
	return &AppError{
		Message: input.Message,
		Fields:  input.Fields,
		Err:     NotFound,
	}
}

func InternalServerError(input ErrorInput) error {
	return &AppError{
		Message: input.Message,
		Fields:  input.Fields,
		Err:     InternalServer,
		Cause:   input.Cause,
	}
}
