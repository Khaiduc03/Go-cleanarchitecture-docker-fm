package exception

type UnauthorizedError struct {
	Message string `json:"message"`
}

func (unauthorizedError UnauthorizedError) Error() UnauthorizedError {
	return unauthorizedError
}
