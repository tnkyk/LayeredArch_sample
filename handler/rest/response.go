package rest

import (
	"fmt"
)

type HttpError struct {
	status int
	method string
	msg    string
}

func (he *HttpError) Error() string {
	return fmt.Sprintf("err: %s [status=%d method=%s]", he.msg, he.status, he.method)
}
