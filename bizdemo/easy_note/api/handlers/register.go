package handlers

import (
	"context"

	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/api/errno"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/api/kitex_gen/kitex/demo/user"
	userprc "github.com/cloudwego/kitex-examples/bizdemo/easy_note/api/rpc/user"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var registerVar UserParam

	if err := c.ShouldBind(&registerVar); err != nil {
		SendResponse(c, errno.DecodeErr(err), nil)
		return
	}

	if err := userprc.CreateUser(context.Background(), &user.CreateUserRequest{
		UserName: registerVar.UserName,
		Password: registerVar.PassWord,
	}); err != nil {
		SendResponse(c, errno.DecodeErr(err), nil)
		return
	}

	SendResponse(c, errno.Success, nil)
}
