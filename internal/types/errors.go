package types

import "fmt"

type GreenAPIError struct {
	message string
	code    int
}

func NewGreenAPIError(message string, code int) *GreenAPIError {
	return &GreenAPIError{message: message, code: code}
}

func (e *GreenAPIError) Error() string {
	return fmt.Sprintf("%s", e.message)
}

func (e *GreenAPIError) GetErrorCode() int {
	return e.code
}
