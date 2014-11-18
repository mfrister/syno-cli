package synoapi

import (
	"fmt"
)

type ClientError interface {
	error
	UnderlyingError() error
}

func (e clientError) Error() string {
	underlyingErrorDesc := ""
	if e.underlyingError != nil {
		underlyingErrorDesc = fmt.Sprintf(": %s", e.underlyingError.Error())
	}
	return fmt.Sprintf("synoapi.Client: %s%s", e.Context(), underlyingErrorDesc)
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

var SYNO_ERR_CODES = map[int]string{
	402:  "This shared folder doesn't exist.",
	3308: "The password is invalid.",
}

type SynoError interface {
	ClientError
	Code() int
	Description() string
}

type synoError struct {
	code int
}

func NewSynoError(code int) SynoError {
	return synoError{code}
}

func (e synoError) Code() int {
	return e.code
}

func (e synoError) Description() string {
	if desc, ok := SYNO_ERR_CODES[e.code]; ok {
		return desc
	} else {
		return "Unknown error code"
	}
}

func (e synoError) Error() string {
	return fmt.Sprintf("synoapi.Client: API returned error code %d: %s", e.code, e.Description())
}

func (e synoError) UnderlyingError() error {
	return nil
}
