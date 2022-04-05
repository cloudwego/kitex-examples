package pack

import (
	"errors"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/business/user/rpc/kitex_gen/base"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/pkg/errno"
)

// BuildBaseResp build baseResp from error
func BuildBaseResp(err error) *base.BaseResp {
	if err == nil {
		return baseResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return baseResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return baseResp(s)
}

func baseResp(err errno.ErrNo) *base.BaseResp {
	return &base.BaseResp{StatusCode: int32(err.ErrCode), StatusMessage: err.ErrMsg}
}
