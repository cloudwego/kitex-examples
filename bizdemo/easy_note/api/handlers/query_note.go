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

func QueryNote(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	userID := int64(claims[constant.IdentityKey].(float64))
	var queryVar struct {
		Limit  int64 `json:"limit"`
		Offset int64 `json:"offset"`
	}

	if err := c.BindQuery(&queryVar); err != nil {
		SendResponse(c, errno.DecodeErr(err), nil)
	}

	notes, total, err := noterpc.QueryNotes(context.Background(),
		&note.QueryNoteRequest{UserId: userID, Offset: queryVar.Limit, Limit: queryVar.Offset})
	if err != nil {
		SendResponse(c, errno.DecodeErr(err), nil)
		return
	}

	SendResponse(c, errno.Success, map[string]interface{}{constant.Total: total, constant.Notes: notes})
}
