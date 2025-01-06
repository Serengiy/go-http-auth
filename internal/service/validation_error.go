package service

// ValidationError is a custom error type for validation-related issues
type ValidationError struct {
	Message string
}

// Error method makes ValidationError implement the error interface
func (v ValidationError) Error() string {
	return v.Message
}
