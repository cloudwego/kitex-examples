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

type Errno struct {
	ErrCode int64
	ErrMsg  string
}

func (e Errno) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", e.ErrCode, e.ErrMsg)
}

var (
	Success    = Errno{ErrCode: 0, ErrMsg: "Success"}
	ServiceErr = Errno{ErrCode: 10001, ErrMsg: "Service is unable to start successfully"}
	ParamErr   = Errno{ErrCode: 10002, ErrMsg: "Wrong Parameter has been given"}
)

func NewErrno(code int64, msg string) Errno {
	return Errno{code, msg}
}

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
