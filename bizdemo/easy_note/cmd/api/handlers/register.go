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

package handlers

import (
	"context"

	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/pkg/errno"

	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/kitex_gen/userdemo"

	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/cmd/api/rpc"

	"github.com/gin-gonic/gin"
)

// Register register user info
func Register(c *gin.Context) {
	var registerVar UserParam
	if err := c.ShouldBind(&registerVar); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	if len(registerVar.UserName) == 0 || len(registerVar.PassWord) == 0 {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	err := rpc.CreateUser(context.Background(), &userdemo.CreateUserRequest{
		UserName: registerVar.UserName,
		Password: registerVar.PassWord,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, nil)
}
