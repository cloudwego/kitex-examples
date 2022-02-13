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
	"net/http"

	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/pkg/errno"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// SendResponse pack response
func SendResponse(c *gin.Context, err error, data interface{}) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, Response{
		Code:    Err.ErrCode,
		Message: Err.ErrMsg,
		Data:    data,
	})
}

type NoteParam struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type UserParam struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}
