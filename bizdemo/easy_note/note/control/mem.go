package control

import (
	"context"
	"fmt"

	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/note/constant"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/note/errno"
	"github.com/shirou/gopsutil/mem"
)

func MemReject(ctx context.Context, request interface{}) error {
	memPercent := memPercent()
	err := errno.Errno{Code: errno.ServiceErr.Code, Msg: fmt.Sprintf("mem = %.2g", memPercent)}
	if memPercent > constant.MemRateLimit {
		return err
	}
	return nil
}

func memPercent() float64 {
	memInfo, _ := mem.VirtualMemory()
	return memInfo.UsedPercent
}
