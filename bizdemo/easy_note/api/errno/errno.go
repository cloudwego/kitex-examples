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
	Code int64
	Msg  string
}

func (e Errno) Error() string {
	return fmt.Sprintf("code=%d, msg=%s", e.Code, e.Msg)
}

var (
	Success    = Errno{Code: 0, Msg: "Success"}
	ServiceErr = Errno{Code: 10001, Msg: "Service is unable to start successfully"}
	ParamErr   = Errno{Code: 10002, Msg: "Wrong Parameter has been given"}
)

func NewErrno(code int64, msg string) Errno {
	return Errno{code, msg}
}

// DecodeErr  translate error to Errno
func DecodeErr(err error) Errno {
	Err := Errno{}
	if errors.As(err, &Err) {
		return Err
	}

	s := ServiceErr
	s.Msg = err.Error()
	return s
}
