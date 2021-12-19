package service

import (
	"context"

	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/note/dal/db"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/note/kitex_gen/kitex/demo/note"
)

type DelNoteService struct {
	ctx context.Context
}

func NewDelNoteService(ctx context.Context) *DelNoteService {
	return &DelNoteService{
		ctx: ctx,
	}
}

func (s *DelNoteService) DelNote(req *note.DelNoteRequest) error {

	if err := db.DelNote(s.ctx, req.NoteId, req.UserId); err != nil {
		return err
	}

	return nil
}
