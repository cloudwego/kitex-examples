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

package handlers

import (
	"context"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/api/constant"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/api/errno"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/api/kitex_gen/kitex/demo/note"
	noteRpc "github.com/cloudwego/kitex-examples/bizdemo/easy_note/api/rpc/note"
	"github.com/gin-gonic/gin"
)

// QueryNote  query list of note info
func QueryNote(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	userID := int64(claims[constant.IdentityKey].(float64))
	var queryVar struct {
		Limit         int64  `json:"limit" form:"limit"`
		Offset        int64  `json:"offset" form:"offset"`
		SearchKeyword string `json:"search_keyword" form:"search_keyword"`
	}
	if err := c.BindQuery(&queryVar); err != nil {
		SendResponse(c, errno.DecodeErr(err), nil)
	}

	if queryVar.Limit < 0 || queryVar.Offset < 0 {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	req := &note.QueryNoteRequest{UserId: userID, Offset: queryVar.Offset, Limit: queryVar.Limit}
	if len(queryVar.SearchKeyword) != 0 {
		req.SearchKey = &queryVar.SearchKeyword
	}
	notes, total, err := noteRpc.QueryNotes(context.Background(), req)
	if err != nil {
		SendResponse(c, errno.DecodeErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, map[string]interface{}{constant.Total: total, constant.Notes: notes})
}
