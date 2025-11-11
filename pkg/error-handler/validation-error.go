package errorhandler

type ValidationError struct {
	Message string `json:"message"`
	Details any    `json:"details,omitempty"`
}

func (v *ValidationError) Error() string {
	return v.Message
}

func NewValidationError(message string, details any) *ValidationError {
	return &ValidationError{
		Message: message,
		Details: details,
	}
}
