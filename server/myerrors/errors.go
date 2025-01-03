package myerrors

type ValidationError struct {
	MESSAGE string
}

func (e *ValidationError) Error() string {
	return e.MESSAGE
}
