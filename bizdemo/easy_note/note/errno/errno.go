package errno

import (
	"errors"
	"fmt"
	"time"

	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/note/kitex_gen/kitex/demo/note"
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
	ParamErr   = Errno{Code: 10002, Msg: "参数错误"}
)

func (e *Errno) ToBaseResp() *note.BaseResp {
	return &note.BaseResp{StatusCode: e.Code, StatusMessage: e.Msg, ServiceTime: time.Now().Unix()}
}

func BuildBaseResp(err error) *note.BaseResp {
	if err == nil {
		return Success.ToBaseResp()
	}

	ErrNo := Errno{}

	if errors.As(err, &ErrNo) {
		return ErrNo.ToBaseResp()
	}

	s := ServiceErr
	s.Msg = err.Error()

	return s.ToBaseResp()

}
