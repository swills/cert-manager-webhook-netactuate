package netactuate

import "errors"

var (
	ErrHTTPNotOK      = errors.New("bad http status code")
	ErrDomainNotFound = errors.New("domain not found")
	ErrUnknown        = errors.New("unknown error")
)
