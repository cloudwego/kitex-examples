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

	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/note/kitex_gen/kitex/demo/note"
)

type Errno struct {
	ErrCode int64
	ErrMsg  string
}

func (e Errno) Error() string {
	return fmt.Sprintf("code=%d, msg=%s", e.ErrCode, e.ErrMsg)
}

var (
	Success    = Errno{ErrCode: 0, ErrMsg: "Success"}
	ServiceErr = Errno{ErrCode: 10001, ErrMsg: "Service is unable to start successfully"}
	ParamErr   = Errno{ErrCode: 10002, ErrMsg: "Wrong Parameter has been given"}
)

// ToBaseResp build baseResp from Errno
func (e *Errno) ToBaseResp() *note.BaseResp {
	return &note.BaseResp{StatusCode: e.ErrCode, StatusMessage: e.ErrMsg, ServiceTime: time.Now().Unix()}
}

// BuildBaseResp build baseResp from error
func BuildBaseResp(err error) *note.BaseResp {
	if err == nil {
		return Success.ToBaseResp()
	}

	ErrNo := Errno{}
	if errors.As(err, &ErrNo) {
		return ErrNo.ToBaseResp()
	}

	s := ServiceErr
	s.ErrMsg = err.Error()
	return s.ToBaseResp()
}
