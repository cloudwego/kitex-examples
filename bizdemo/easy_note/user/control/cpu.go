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

package control

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/user/constant"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/user/errno"
	"github.com/shirou/gopsutil/cpu"
)

// CPUReject cpu acl control
func CPUReject(ctx context.Context, request interface{}) error {
	c := cpuPercent()
	if c > constant.CPURateLimit {
		return errno.ServiceErr.WithMessage(fmt.Sprintf("cpu = %.2g", c))
	}
	return nil
}

func cpuPercent() float64 {
	percent, _ := cpu.Percent(time.Second, false)
	return percent[0]
}
