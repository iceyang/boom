package boom

import (
	"net/http"

	"github.com/go-errors/errors"
)

type boomFunc func(msgs ...string) *Error

func wrapFunc(code int) boomFunc {
	return func(msgs ...string) *Error {
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
}

func Boom(code int, msgs ...string) *Error {
	return wrapFunc(code)(msgs...)
}
