package service

import (
	"context"

	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/note/dal/db"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/note/kitex_gen/kitex/demo/note"
)

type UpdateNoteService struct {
	ctx context.Context
}

func NewUpdateNoteService(ctx context.Context) *UpdateNoteService {
	return &UpdateNoteService{ctx: ctx}
}

func (s *UpdateNoteService) UpdateNote(req *note.UpdateNoteRequest) error {
	if err := db.UpdateNote(s.ctx, req.NoteId, req.UserId, req.Title, req.Content); err != nil {
		return err
	}

	return nil
}
