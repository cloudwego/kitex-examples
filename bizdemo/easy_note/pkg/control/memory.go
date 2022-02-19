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

	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/pkg/errno"

	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/pkg/constants"

	"github.com/shirou/gopsutil/mem"
)

// MemReject memory access control
func MemReject(ctx context.Context, request interface{}) error {
	m := memPercent()
	if m > constants.MemRateLimit {
		return errno.ServiceErr.WithMessage(fmt.Sprintf("mem = %.2g", m))
	}
	return nil
}

func memPercent() float64 {
	memInfo, _ := mem.VirtualMemory()
	return memInfo.UsedPercent
}