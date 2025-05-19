package calculator

func newClientError(err error) *ClientError {
	return &ClientError{
		err: err,
	}
}

type ClientError struct {
	err error
}

func (e *ClientError) Error() string {
	return e.err.Error()
}
