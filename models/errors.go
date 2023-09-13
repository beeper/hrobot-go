package models

import (
	"fmt"
)

// 400 - 499 Range
type ErrorClientSide struct {
	StatusCode int
	Body       []byte
}

func (e *ErrorClientSide) Error() string {
	return fmt.Sprintf("%d-%s", e.StatusCode, e.Body)
}

// 500+ Range
type ErrorServerSide struct {
	StatusCode int
	Body       []byte
}

func (e *ErrorServerSide) Error() string {
	return fmt.Sprintf("%d-%s", e.StatusCode, e.Body)
}
