package main

import (
	"context"
	"log"
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/api/constant"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/api/handlers"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/api/kitex_gen/kitex/demo/user"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/api/rpc"
	userrpc "github.com/cloudwego/kitex-examples/bizdemo/easy_note/api/rpc/user"
	"github.com/gin-gonic/gin"
)

func Init() {
	rpc.Init()
}

func main() {
	Init()
	r := gin.New()
	authMiddleware, _ := jwt.New(&jwt.GinJWTMiddleware{
		Key:        []byte(constant.SecretKey),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(int64); ok {
				return jwt.MapClaims{
					constant.IdentityKey: v,
				}
			}
			return jwt.MapClaims{}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVar handlers.UserParam
			if err := c.ShouldBind(&loginVar); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			return userrpc.CheckUser(context.Background(), &user.CheckUserRequest{UserName: loginVar.UserName, Password: loginVar.PassWord})

		},
		TokenLookup: "header: Authorization, query: token, cookie: jwt",

		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})

	r.POST("/login", authMiddleware.LoginHandler)
	r.POST("/register", handlers.Register)

	auth := r.Group("/auth")
	auth.Use(authMiddleware.MiddlewareFunc())
	auth.GET("/note", handlers.QueryNote)
	auth.POST("/note", handlers.CreateNote)
	auth.PUT("/note/:note_id", handlers.UpdateNote)
	auth.DELETE("/note/:note_id", handlers.DeleteNote)

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
