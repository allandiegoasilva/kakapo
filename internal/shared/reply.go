package shared

type Empty struct{}

type Reply[T any] struct {
	Success bool          `json:"success"`
	Message *string       `json:"message,omitempty"`
	Data    T             `json:"data,omitempty"`
	Errors  []*FieldError `json:"errors,omitempty"`
}

func SuccessReply[T any](data T) Reply[T] {
	return Reply[T]{
		Success: true,
		Data:    data,
	}
}

func ErrorReply(message string, fieldErrors []*FieldError) Reply[any] {
	return Reply[any]{
		Success: false,
		Message: &message,
		Errors:  fieldErrors,
	}
}

func FieldErrorsFromAppError(fields []FieldError) []*FieldError {
	if len(fields) == 0 {
		return nil
	}

	result := make([]*FieldError, len(fields))
	for i := range fields {
		result[i] = &fields[i]
	}

	return result
}
