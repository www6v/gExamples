package basic

import (
	"errors"
	"fmt"
	"testing"
)

func Test_Main1(t *testing.T) {
	cause := errors.New("What's the cause?!")
	err := WithMessagef(cause, "something unusual has occurred")
	fmt.Println(err)

	err2 := WithMessagef(nil, "something unusual has occurred")
	fmt.Println(err2)

}

func WithMessagef(err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &withMessage{
		cause: err,
		msg:   fmt.Sprintf(format, args...),
	}
}

type withMessage struct {
	cause error
	msg   string
}

func (w *withMessage) Error() string { return w.msg + ": " + w.cause.Error() }
