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
