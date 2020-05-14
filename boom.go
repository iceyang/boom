package boom

import (
	"net/http"

	"github.com/go-errors/errors"
)

type Error struct {
	Err        *errors.Error
	msg        string
	statusCode int
}

func (e *Error) Unwarp() error {
	return e.Err
}

func (e *Error) Error() string {
	return e.msg
}

func (e *Error) Message() string {
	return e.msg
}

func (e *Error) Status() int {
	return e.statusCode
}

func (e *Error) Stack() []byte {
	return e.Err.Stack()
}

func (e *Error) ErrorStack() string {
	return e.Err.ErrorStack()
}

func Boom(code int, msgs ...string) *Error {
	msg := http.StatusText(code)
	if len(msgs) > 0 {
		msg = msgs[0]
	}
	return &Error{
		Err:        errors.Wrap(msg, 1),
		msg:        msg,
		statusCode: code,
	}
}
