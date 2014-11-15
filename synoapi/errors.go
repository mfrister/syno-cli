package synoapi

import (
	"fmt"
)

type ClientError interface {
	error
	Context() string
	UnderlyingError() error
}

func (e clientError) Error() string {
	return fmt.Sprintf("synoapi.Client: %s: %s", e.context, e.underlyingError.Error())
}

type clientError struct {
	context         string
	underlyingError error
}

func NewClientError(context string, err error) ClientError {
	return clientError{context, err}
}

func (e clientError) Context() string {
	return e.context
}

func (e clientError) UnderlyingError() error {
	return e.underlyingError
}
