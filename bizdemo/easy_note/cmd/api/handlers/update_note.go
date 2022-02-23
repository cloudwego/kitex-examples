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
	"strconv"

	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/pkg/constants"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/pkg/errno"

	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/kitex_gen/notedemo"

	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/cmd/api/rpc"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// UpdateNote update user info
func UpdateNote(c *gin.Context) {
	var noteVar NoteParam
	if err := c.ShouldBind(&noteVar); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	claims := jwt.ExtractClaims(c)
	userID := int64(claims[constants.IdentityKey].(float64))
	noteIDStr := c.Param(constants.NoteID)
	noteID, err := strconv.ParseInt(noteIDStr, 10, 64)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	if noteID <= 0 {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	req := &notedemo.UpdateNoteRequest{NoteId: noteID, UserId: userID}
	if len(noteVar.Title) != 0 {
		req.Title = &noteVar.Title
	}
	if len(noteVar.Content) != 0 {
		req.Content = &noteVar.Content
	}
	if err = rpc.UpdateNote(context.Background(), req); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, nil)
}
