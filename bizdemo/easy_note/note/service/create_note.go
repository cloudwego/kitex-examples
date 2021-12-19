package service

import (
	"context"

	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/note/dal/db"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/note/kitex_gen/kitex/demo/note"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/note/model"
)

type CreateNoteService struct {
	ctx context.Context
}

func NewCreateNoteService(ctx context.Context) *CreateNoteService {
	return &CreateNoteService{ctx: ctx}
}

func (s *CreateNoteService) CreateNote(req *note.CreateNoteRequest) error {
	noteModel := &model.Note{
		UserID:  req.UserId,
		Title:   req.Title,
		Content: req.Content,
	}
	if err := db.CreateNote(s.ctx, []*model.Note{noteModel}); err != nil {
		return err
	}

	return nil
}
