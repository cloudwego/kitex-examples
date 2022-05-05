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
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/business/mall_api/dal"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/business/mall_api/dal/client"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/business/mall_api/handlers"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/business/mall_api/kitex_gen/cmp/ecom/user"
	_ "github.com/cloudwego/kitex-examples/bizdemo/mall/docs"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/pkg/conf"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Init() {
	dal.Init()

	handlers.AuthMiddleware, _ = jwt.New(&jwt.GinJWTMiddleware{
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
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVar handlers.UserParam
			if err := c.ShouldBind(&loginVar); err != nil {
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
// @description This is a mall demo using Gin and KiteX.

// @contact.name Bodhisatan
// @contact.url https://github.com/bodhisatan
// @contact.email bodhisatanyao@gmail.com

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
	r := gin.New()

	userGroup := r.Group("/user")
	userGroup.POST("/register", handlers.Register)
	userGroup.POST("/login", handlers.Login)

	shopGroup := r.Group("/shop")
	shopGroup.Use(handlers.AuthMiddleware.MiddlewareFunc())
	shopGroup.POST("/settle", handlers.SettleShop)
	shopGroup.GET("/id", handlers.GetShopId)

	productGroup := r.Group("/product")
	productGroup.Use(handlers.AuthMiddleware.MiddlewareFunc())
	productGroup.POST("/brand_create", handlers.BrandCreate)
	productGroup.POST("/brand_del", handlers.BrandDel)
	productGroup.POST("/brand_update", handlers.BrandUpdate)
	productGroup.GET("/get_brands", handlers.GetBrands)

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	if err := http.ListenAndServe(":8080", r); err != nil {
		klog.Fatal(err)
	}
}
