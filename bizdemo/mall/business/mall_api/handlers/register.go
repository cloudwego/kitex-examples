package handlers

import (
	"github.com/cloudwego/kitex-examples/bizdemo/mall/business/mall_api/dal/client"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/business/mall_api/kitex_gen/cmp/ecom/user"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/pkg/errno"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var registerParam UserParam
	if err := c.ShouldBind(&registerParam); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	if len(registerParam.UserName) == 0 || len(registerParam.PassWord) == 0 {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	err := client.CreateUser(c, &user.CreateUserReq{
		UserName: registerParam.UserName,
		Password: registerParam.PassWord,
	})

	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, nil)
}
