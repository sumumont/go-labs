package main

import (
	"fmt"
	"github.com/go-labs/internal/logging"
	"sync"
)

type MultiError struct {
	Code   int         `json:"code"`
	Data   interface{} `json:"data"`
	Errors []error     `json:"errors"`
	sync.Mutex
}

func (e *MultiError) Error() string {
	if len(e.Errors) > 100 {
		s := e.Errors[:100]
		return fmt.Sprintf("errors count:%d,is too much,just print 100 errors: %v", len(e.Errors), s)
	}
	return fmt.Sprintf("errors count:%d errors: %v", len(e.Errors), e.Errors)
}
func (e *MultiError) Errno() int {
	return e.Code
}
func (e *MultiError) Append(errs ...error) {
	e.Errors = append(e.Errors, errs...)
}
func (e *MultiError) Println() {
	for _, err := range e.Errors {
		logging.Error(err).Send()
	}
}
func (e *MultiError) IsErr() bool {
	return len(e.Errors) != 0
}
func NewMultiError() *MultiError {
	return &MultiError{
		Code: -1,
	}
}
