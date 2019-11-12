package persistence

import (
	"fmt"
)

type SQLError struct {
	Msg string
	Err error
}

func (se *SQLError) Error() string {
	return fmt.Sprintf("%s:%s", se.Msg, se.Err)
}

func NewSQLError(err error, msg string) *SQLError {
	return &SQLError{
		Msg: msg,
		Err: err,
	}
}
