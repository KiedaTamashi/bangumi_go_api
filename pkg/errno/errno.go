package errno

import (
	"fmt"
	"github.com/pkg/errors"
)

const (
	ErrUnknown        = -101
	ErrBadRequest     = -102 // http 400
	ErrInternalServer = -103 // http 5xx
)

type ErrNo struct {
	code int
	err  error
}

func (e *ErrNo) Error() string {
	return e.err.Error()
}

func (e *ErrNo) StackTrace() string {
	return fmt.Sprintf("[code=%d] %+v", e.code, e.err)
}

func (e *ErrNo) Code() int {
	return e.code
}

func Errorf(code int, format string, args ...interface{}) error {
	return &ErrNo{
		code: code,
		err:  errors.Errorf(format, args...),
	}
}
