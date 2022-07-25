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

package main

import (
	"context"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/cmd/mall_api/dal"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/cmd/mall_api/dal/client"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/cmd/mall_api/handlers"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/cmd/mall_api/kitex_gen/cmp/ecom/user"
	_ "github.com/cloudwego/kitex-examples/bizdemo/mall/docs"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/pkg/conf"
	"github.com/hertz-contrib/jwt"
	"github.com/hertz-contrib/swagger"
	swaggerFiles "github.com/swaggo/files"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func Init() {
	dal.Init()

	handlers.AuthMiddleware, _ = jwt.New(&jwt.HertzJWTMiddleware{
		Key:        []byte(conf.SecretKey),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(int64); ok {
				return jwt.MapClaims{
					conf.IdentityKey: v,
				}
			}
			return jwt.MapClaims{}
		},
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var loginVar handlers.UserParam
			if err := c.Bind(&loginVar); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			if len(loginVar.UserName) == 0 || len(loginVar.PassWord) == 0 {
				return "", jwt.ErrMissingLoginValues
			}

			return client.CheckUser(context.Background(), &user.CheckUserReq{UserName: loginVar.UserName, Password: loginVar.PassWord})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})
}

// @title Mall
// @version 1.0
// @description This is a mall demo using Hertz and KiteX.

// @contact.name CloudWeGo
// @contact.url https://github.com/cloudwego
// @contact.email conduct@cloudwego.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.apikey TokenAuth
// @in header
// @name Authorization

// @host localhost:8080
// @BasePath /
// @schemes http
func main() {
	Init()
	h := server.Default(server.WithHostPorts("localhost:8080"))

	userGroup := h.Group("/user")
	userGroup.POST("/register", handlers.Register)
	userGroup.POST("/login", handlers.Login)

	shopGroup := h.Group("/shop")
	shopGroup.Use(handlers.AuthMiddleware.MiddlewareFunc())
	shopGroup.POST("/settle", handlers.SettleShop)
	shopGroup.GET("/id", handlers.GetShopId)

	productGroup := h.Group("/product")
	productGroup.Use(handlers.AuthMiddleware.MiddlewareFunc())
	productGroup.POST("/brand_create", handlers.BrandCreate)
	productGroup.POST("/brand_del", handlers.BrandDel)
	productGroup.POST("/brand_update", handlers.BrandUpdate)
	productGroup.GET("/get_brands", handlers.GetBrands)

	url := swagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	h.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler, url))

	h.Spin()
}
