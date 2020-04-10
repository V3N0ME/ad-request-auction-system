package errors

import "fmt"

//AppError struct holds the error code and message
type AppError struct {
	code        int
	description string
}

func (e AppError) Error() string {
	return fmt.Sprintf("%s", e.description)
}

//Code returns the error code
func (e AppError) Code() string {
	return fmt.Sprintf("%d", e.code)
}

//New initializes new error
func New(description string, code int) error {
	return &AppError{
		code:        code,
		description: description,
	}
}
