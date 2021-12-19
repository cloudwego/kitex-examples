package handlers

import (
	"context"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/api/constant"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/api/errno"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/api/kitex_gen/kitex/demo/note"
	noterpc "github.com/cloudwego/kitex-examples/bizdemo/easy_note/api/rpc/note"
	"github.com/gin-gonic/gin"
)

func CreateNote(c *gin.Context) {
	var noteVar NoteParam

	if err := c.ShouldBind(&noteVar); err != nil {
		SendResponse(c, errno.DecodeErr(err), nil)
		return
	}
	claims := jwt.ExtractClaims(c)
	userID := int64(claims[constant.IdentityKey].(float64))
	if err := noterpc.CreateNote(context.Background(), &note.CreateNoteRequest{UserId: userID,
		Content: noteVar.Content, Title: noteVar.Title}); err != nil {
		SendResponse(c, errno.DecodeErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, nil)
}
