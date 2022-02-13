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
)

const (
	SuccessCode             = 0
	ServiceErrCode          = 10001
	ParamErrCode            = 10002
	LoginErrCode            = 10003
	UserNotExistErrCode     = 10004
	UserAlreadyExistErrCode = 10005
)

type Errno struct {
	ErrCode int64
	ErrMsg  string
}

func (e Errno) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", e.ErrCode, e.ErrMsg)
}

func NewErrno(code int64, msg string) Errno {
	return Errno{code, msg}
}

func (e Errno) WithMessage(msg string) Errno {
	e.ErrMsg = msg
	return e
}

var (
	Success             = Errno{ErrCode: SuccessCode, ErrMsg: "Success"}
	ServiceErr          = Errno{ErrCode: ServiceErrCode, ErrMsg: "Service is unable to start successfully"}
	ParamErr            = Errno{ErrCode: ParamErrCode, ErrMsg: "Wrong Parameter has been given"}
	LoginErr            = NewErrno(LoginErrCode, "Wrong username or password")
	UserNotExistErr     = NewErrno(UserNotExistErrCode, "User does not exists")
	UserAlreadyExistErr = NewErrno(UserAlreadyExistErrCode, "User already exists")
)

// ConvertErr convert error to Errno
func ConvertErr(err error) Errno {
	Err := Errno{}
	if errors.As(err, &Err) {
		return Err
	}

	s := ServiceErr
	s.ErrMsg = err.Error()
	return s
}
