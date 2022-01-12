package types

type BasicError struct {
	message string
}

func NewBasicError(message string) *BasicError {
	return &BasicError{message: message}
}

func (be *BasicError) Error() string {
	return "[BASIC ERROR] " + be.message
}
