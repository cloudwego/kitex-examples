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

func NewMGetNoteService(ctx context.Context) *MGetNoteService {
	return &MGetNoteService{ctx: ctx}
}

func (s *MGetNoteService) MGetNote(req *note.MGetNoteRequest) ([]*note.Note, error) {
	noteModels, err := db.MGetNotes(s.ctx, req.NoteIds)
	if err != nil {
		return nil, err
	}

	notes := pack.Notes(noteModels)
	return notes, nil
}
