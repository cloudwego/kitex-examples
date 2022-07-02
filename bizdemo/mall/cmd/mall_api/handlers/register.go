// Copyright 2022 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package handlers

import (
	"github.com/cloudwego/kitex-examples/bizdemo/mall/cmd/mall_api/dal/client"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/cmd/mall_api/kitex_gen/cmp/ecom/user"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/pkg/errno"
	"github.com/gin-gonic/gin"
)

// Register godoc
// @Summary 用户注册
// @Description 用户注册
// @Tags 用户模块
// @Accept json
// @Produce json
// @Param userParam body handlers.UserParam true "注册信息"
// @Success 200 {object} handlers.Response
// @Router /user/register [post]
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
