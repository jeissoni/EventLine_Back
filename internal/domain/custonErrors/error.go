package domain

import (
	"errors"
	"fmt"
)

const (
	ErrCodeDuplicatedKey = "duplicated_key"
	ErrCodeNotFound      = "not_found"
)

var (
	ErrDuplicatedKey = errors.New("duplicated key")
	ErrNotFound      = errors.New("not found")
)

type DomainError struct {
	Code    string `json:"code"` // duplicated_key, not_found, etc
	Message string `json:"message"`
}

func NewDomainError(code, message string) DomainError {
	return DomainError{
		Code:    code,
		Message: message,
	}
}

func (e DomainError) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}
