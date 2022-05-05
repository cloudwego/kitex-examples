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
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/pkg/errno"
	"github.com/gin-gonic/gin"
)

var AuthMiddleware *jwt.GinJWTMiddleware

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

type UserParam struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

type LoginResponse struct {
	Code   int64  `json:"code"`
	Expire string `json:"expire"`
	Token  string `json:"token"`
}

type ShopSettleParam struct {
	ShopName string `json:"shop_name"`
}

type BrandAddParam struct {
	BrandName  string `json:"brand_name"`
	Logo       string `json:"logo"`
	BrandStory string `json:"brand_story"`
}

type BrandEditParam struct {
	BrandId    int64   `json:"brand_id" binding:"required"`
	BrandName  *string `json:"brand_name"`
	Logo       *string `json:"logo"`
	BrandStory *string `json:"brand_story"`
}

type BrandDelParam struct {
	BrandId int64 `json:"brand_id" binding:"required"`
}
