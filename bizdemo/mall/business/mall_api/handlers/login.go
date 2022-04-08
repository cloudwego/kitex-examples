package handlers

import "github.com/gin-gonic/gin"

// Login godoc
// @Summary 用户登录
// @Description 用户登录
// @Tags 用户模块
// @Accept json
// @Produce json
// @Param userParam body handlers.UserParam true "账号信息"
// @Success 200 {object} handlers.LoginResponse
// @Router /user/login [post]
func Login(c *gin.Context) {
	AuthMiddleware.LoginHandler(c)
}
