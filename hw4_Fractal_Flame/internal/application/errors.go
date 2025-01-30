package application

import "fmt"

type AppError struct {
	Message string
	Err     error
}

func (e *AppError) Error() string {
	return fmt.Sprintf("%s: %v", e.Message, e.Err)
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func WrapError(err error, message string) error {
	return &AppError{
		Message: message,
		Err:     err,
	}
}
