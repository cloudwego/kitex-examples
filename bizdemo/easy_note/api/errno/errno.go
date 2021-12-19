package errno

import (
	"errors"
	"fmt"
)

type Errno struct {
	Code int64
	Msg  string
}

func (e Errno) Error() string {
	return fmt.Sprintf("code=%d,msg=%s", e.Code, e.Msg)
}

var (
	Success    = Errno{Code: 0, Msg: "success"}
	ServiceErr = Errno{Code: 10001, Msg: "服务器开小差了"}
)

func NewErrno(code int64, msg string) Errno {
	return Errno{code, msg}
}

func DecodeErr(err error) Errno {
	var Err = Errno{}

	if errors.As(err, &Err) {
		return Err
	}
	s := ServiceErr
	s.Msg = err.Error()

	return s
}
