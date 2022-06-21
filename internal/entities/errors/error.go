package errors

import "errors"

var (
	InvalidDomain = errors.New("domain name is invalid")
	BadTLD        = errors.New("domain name is not .be")
)
