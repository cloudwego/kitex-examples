package control

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/user/constant"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/user/errno"
	"github.com/shirou/gopsutil/cpu"
)

func CpuReject(ctx context.Context, request interface{}) error {
	cpuPercent := cpuPercent()
	err := errno.Errno{Code: errno.ServiceErr.Code, Msg: fmt.Sprintf("cpu = %.2g", cpuPercent)}
	if cpuPercent > constant.CpuRateLimit {
		return err
	}
	return nil
}

func cpuPercent() float64 {
	percent, _ := cpu.Percent(time.Second, false)
	return percent[0]
}
