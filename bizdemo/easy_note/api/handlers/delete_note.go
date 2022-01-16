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

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/api/constant"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/api/errno"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/api/kitex_gen/kitex/demo/note"
	noteRpc "github.com/cloudwego/kitex-examples/bizdemo/easy_note/api/rpc/note"
	"github.com/gin-gonic/gin"
)

// DeleteNote  delete note info
func DeleteNote(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	userID := int64(claims[constant.IdentityKey].(float64))
	noteIDStr := c.Param(constant.NoteID)
	noteID, err := strconv.ParseInt(noteIDStr, 10, 64)
	if err != nil {
		SendResponse(c, errno.DecodeErr(err), nil)
		return
	}

	if err = noteRpc.DelNote(context.Background(), &note.DelNoteRequest{
		NoteId: noteID, UserId: userID,
	}); err != nil {
		SendResponse(c, errno.DecodeErr(err), nil)
		return
	}

	SendResponse(c, errno.Success, nil)
}
