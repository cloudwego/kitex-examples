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

func DeleteNote(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	userID := int64(claims[constant.IdentityKey].(float64))
	noteIDStr := c.Param(constant.NoteID)
	noteID, err := strconv.ParseInt(noteIDStr, 10, 64)
	if err != nil {
		SendResponse(c, errno.DecodeErr(err), nil)
		return
	}

	if err = noterpc.DelNote(context.Background(), &note.DelNoteRequest{
		NoteId: noteID, UserId: userID,
	}); err != nil {
		SendResponse(c, errno.DecodeErr(err), nil)
		return
	}

	SendResponse(c, errno.Success, nil)
}
