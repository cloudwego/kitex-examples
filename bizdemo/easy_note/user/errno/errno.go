// Copyright 2021 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package errno

import (
	"errors"
	"fmt"
	"time"

	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/user/kitex_gen/kitex/demo/user"
)

type Errno struct {
	Code int64
	Msg  string
}

func (e Errno) Error() string {
	return fmt.Sprintf("code=%d,msg=%s", e.Code, e.Msg)
}

var (
	Success             = Errno{Code: 0, Msg: "success"}
	ServiceErr          = Errno{Code: 10001, Msg: "服务器开小差了(Service is unable to start successfully)"}
	ParamErr            = Errno{Code: 10002, Msg: "参数错误(Wrong parameter has been given)"}
	LoginErr            = Errno{Code: 10003, Msg: "用户名或密码错误(Wrong username or password)"}
	UserNotExistErr     = Errno{Code: 10004, Msg: "用户不存在(User does not exist)"}
	UserAlreadyExistErr = Errno{Code: 10004, Msg: "用户已存在(User exists already)"}
)

// ToBaseResp  build baseResp from Errno
func (e *Errno) ToBaseResp() *user.BaseResp {
	return &user.BaseResp{StatusCode: e.Code, StatusMessage: e.Msg, ServiceTime: time.Now().Unix()}
}

// BuildBaseResp  build baseResp from error
func BuildBaseResp(err error) *user.BaseResp {
	if err == nil {
		return Success.ToBaseResp()
	}

	ErrNo := Errno{}

	if errors.As(err, &ErrNo) {
		return ErrNo.ToBaseResp()
	}
	s := ServiceErr
	s.Msg = err.Error()

	return ServiceErr.ToBaseResp()
}
