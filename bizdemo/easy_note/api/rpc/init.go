package rpc

import (
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/api/rpc/note"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/api/rpc/user"
)

func Init() {
	user.Init()
	note.Init()
}
