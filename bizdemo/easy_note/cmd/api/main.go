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

package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/cmd/api/handlers"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/cmd/api/rpc"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/kitex_gen/userdemo"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/pkg/constants"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/pkg/errno"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/pkg/tracer"
	"github.com/hertz-contrib/jwt"
)

func Init() {
	tracer.InitJaeger(constants.ApiServiceName)
	rpc.InitRPC()
}

func main() {
	Init()
	r := server.New(
		server.WithHostPorts("127.0.0.1:8080"),
		server.WithHandleMethodNotAllowed(true),
	)
	authMiddleware, _ := jwt.New(&jwt.HertzJWTMiddleware{
		Key:        []byte(constants.SecretKey),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(int64); ok {
				return jwt.MapClaims{
					constants.IdentityKey: v,
				}
			}
			return jwt.MapClaims{}
		},
		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			switch e.(type) {
			case errno.ErrNo:
				return e.(errno.ErrNo).ErrMsg
			default:
				return e.Error()
			}
		},
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			c.JSON(consts.StatusOK, map[string]interface{}{
				"code":   errno.SuccessCode,
				"token":  token,
				"expire": expire.Format(time.RFC3339),
			})
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(code, map[string]interface{}{
				"code":    errno.AuthorizationFailedErrCode,
				"message": message,
			})
		},
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var loginVar handlers.UserParam
			if err := c.Bind(&loginVar); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			if len(loginVar.UserName) == 0 || len(loginVar.PassWord) == 0 {
				return "", jwt.ErrMissingLoginValues
			}

			return rpc.CheckUser(context.Background(), &userdemo.CheckUserRequest{UserName: loginVar.UserName, Password: loginVar.PassWord})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})

	r.Use(recovery.Recovery(recovery.WithRecoveryHandler(
		func(ctx context.Context, c *app.RequestContext, err interface{}, stack []byte) {
			hlog.SystemLogger().CtxErrorf(ctx, "[Recovery] err=%v\nstack=%s", err, stack)
			c.JSON(consts.StatusInternalServerError, map[string]interface{}{
				"code":    errno.ServiceErrCode,
				"message": fmt.Sprintf("[Recovery] err=%v\nstack=%s", err, stack),
			})
		})))

	v1 := r.Group("/v1")
	user1 := v1.Group("/user")
	user1.POST("/login", authMiddleware.LoginHandler)
	user1.POST("/register", handlers.Register)

	note1 := v1.Group("/note")
	note1.Use(authMiddleware.MiddlewareFunc())
	note1.GET("/query", handlers.QueryNote)
	note1.POST("", handlers.CreateNote)
	note1.PUT("/:note_id", handlers.UpdateNote)
	note1.DELETE("/:note_id", handlers.DeleteNote)

	r.NoRoute(func(ctx context.Context, c *app.RequestContext) {
		c.String(consts.StatusOK, "no route")
	})
	r.NoMethod(func(ctx context.Context, c *app.RequestContext) {
		c.String(consts.StatusOK, "no method")
	})
	r.Spin()
}
