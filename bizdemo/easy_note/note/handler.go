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

package main

import (
	"context"

	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/note/errno"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/note/kitex_gen/kitex/demo/note"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/note/service"
)

// NoteServiceImpl implements the last service interface defined in the IDL.
type NoteServiceImpl struct{}

// CreateNote implements the NoteServiceImpl interface.
func (s *NoteServiceImpl) CreateNote(ctx context.Context, req *note.CreateNoteRequest) (resp *note.CreateNoteResponse, err error) {
	resp = new(note.CreateNoteResponse)

	if req.UserId <= 0 || len(req.Title) == 0 || len(req.Content) == 0 {
		resp.BaseResp = errno.ParamErr.ToBaseResp()
		return resp, nil
	}

	err = service.NewCreateNoteService(ctx).CreateNote(req)
	if err != nil {
		resp.BaseResp = errno.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = errno.Success.ToBaseResp()
	return resp, nil
}

// MGetNote implements the NoteServiceImpl interface.
func (s *NoteServiceImpl) MGetNote(ctx context.Context, req *note.MGetNoteRequest) (resp *note.MGetNoteResponse, err error) {
	resp = new(note.MGetNoteResponse)

	if len(req.NoteIds) == 0 {
		resp.BaseResp = errno.ParamErr.ToBaseResp()
		return resp, nil
	}

	notes, err := service.NewMGetNoteService(ctx).MGetNote(req)
	if err != nil {
		resp.BaseResp = errno.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = errno.Success.ToBaseResp()
	resp.Notes = notes
	return resp, nil
}

// DelNote implements the NoteServiceImpl interface.
func (s *NoteServiceImpl) DelNote(ctx context.Context, req *note.DelNoteRequest) (resp *note.DelNoteResponse, err error) {
	resp = new(note.DelNoteResponse)

	if req.NoteId <= 0 {
		resp.BaseResp = errno.ParamErr.ToBaseResp()
		return resp, nil
	}

	err = service.NewDelNoteService(ctx).DelNote(req)
	if err != nil {
		resp.BaseResp = errno.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = errno.Success.ToBaseResp()
	return resp, nil
}

// QueryNote implements the NoteServiceImpl interface.
func (s *NoteServiceImpl) QueryNote(ctx context.Context, req *note.QueryNoteRequest) (resp *note.QueryNoteResponse, err error) {
	resp = new(note.QueryNoteResponse)

	if req.UserId <= 0 || req.Limit < 0 || req.Offset < 0 {
		resp.BaseResp = errno.ParamErr.ToBaseResp()
		return resp, nil
	}

	notes, total, err := service.NewQueryNoteService(ctx).QueryNoteService(req)
	if err != nil {
		resp.BaseResp = errno.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = errno.Success.ToBaseResp()
	resp.Notes = notes
	resp.Total = total
	return resp, nil
}

// UpdateNote implements the NoteServiceImpl interface.
func (s *NoteServiceImpl) UpdateNote(ctx context.Context, req *note.UpdateNoteRequest) (resp *note.UpdateNoteResponse, err error) {
	resp = new(note.UpdateNoteResponse)

	if req.NoteId <= 0 {
		resp.BaseResp = errno.ParamErr.ToBaseResp()
		return resp, nil
	}

	err = service.NewUpdateNoteService(ctx).UpdateNote(req)
	if err != nil {
		resp.BaseResp = errno.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = errno.Success.ToBaseResp()
	return resp, nil
}
