package exceptions

type ValidationError struct {
	Message string
}

func NewValidationError(message string) *ValidationError {
	return &ValidationError { Message: message }
}

func (validationError *ValidationError) Error() string {
	return validationError.Message
}

type ServerError struct {
	Message string
}

func NewServerError(message string) *ServerError {
	return &ServerError { Message: message }
}

func (serverError *ServerError) Error() string {
	return serverError.Message
}