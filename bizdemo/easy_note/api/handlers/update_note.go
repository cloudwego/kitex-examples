package handlers

import (
	"context"
	"strconv"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/api/constant"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/api/errno"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/api/kitex_gen/kitex/demo/note"
	noterpc "github.com/cloudwego/kitex-examples/bizdemo/easy_note/api/rpc/note"
	"github.com/gin-gonic/gin"
)

func UpdateNote(c *gin.Context) {
	var noteVar NoteParam

	if err := c.ShouldBind(&noteVar); err != nil {
		SendResponse(c, errno.DecodeErr(err), nil)
		return
	}

	claims := jwt.ExtractClaims(c)
	userID := int64(claims[constant.IdentityKey].(float64))

	noteIDStr := c.Param(constant.NoteID)
	noteID, err := strconv.ParseInt(noteIDStr, 10, 64)
	if err != nil {
		SendResponse(c, errno.DecodeErr(err), nil)
		return
	}

	req := &note.UpdateNoteRequest{NoteId: noteID, UserId: userID}
	if len(noteVar.Title) != 0 {
		req.Title = &noteVar.Title
	}

	if len(noteVar.Content) != 0 {
		req.Content = &noteVar.Content
	}

	if err := noterpc.UpdateNote(context.Background(), req); err != nil {
		SendResponse(c, errno.DecodeErr(err), nil)
		return
	}

	SendResponse(c, errno.Success, nil)
}
