package handlers

import (
	"net/http"

	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/api/errno"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendResponse(c *gin.Context, err error, data interface{}) {
	Err := errno.DecodeErr(err)
	c.JSON(http.StatusOK, Response{
		Code:    Err.Code,
		Message: Err.Msg,
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
