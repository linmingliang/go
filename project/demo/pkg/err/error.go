package err

import (
	"fmt"
)

var (
	Success     = NewError(10000000, "成功")
	ServerError = NewError(10000001, "系统错误")
)

type Error struct {
	code int
	msg  string
}

func (err *Error) Error() string {
	return fmt.Sprintf("code:%d,msg:%s", err.code, err.msg)
}

func (err *Error) Code() int {
	return err.code
}

func (err *Error) Msg() string {
	return err.msg
}
func NewError(code int, msg string) *Error {
	return &Error{
		code: code,
		msg:  msg,
	}
}
