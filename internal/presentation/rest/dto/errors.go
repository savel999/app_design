package dto

type ValidationErrors struct {
	Message string
	Errors  []error
}

func (e *ValidationErrors) Error() string {
	return e.Message
}

func NewValidationErrors(message string, errors []error) *ValidationErrors {
	return &ValidationErrors{Message: message, Errors: errors}
}
