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

package service

import (
	"context"

	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/note/dal/db"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/note/kitex_gen/kitex/demo/note"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/note/pack"
)

type MGetNoteService struct {
	ctx context.Context
}

// NewMGetNoteService  new MGetNoteService
func NewMGetNoteService(ctx context.Context) *MGetNoteService {
	return &MGetNoteService{ctx: ctx}
}

// MGetNote  multiple get list of note info
func (s *MGetNoteService) MGetNote(req *note.MGetNoteRequest) ([]*note.Note, error) {
	noteModels, err := db.MGetNotes(s.ctx, req.NoteIds)
	if err != nil {
		return nil, err
	}

	notes := pack.Notes(noteModels)
	return notes, nil
}
