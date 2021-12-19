package pack

import (
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/note/kitex_gen/kitex/demo/note"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/note/model"
)

func Note(m *model.Note) *note.Note {
	if m == nil {
		return nil
	}

	return &note.Note{
		NoteId:     int64(m.ID),
		UserId:     m.UserID,
		Title:      m.Title,
		Content:    m.Content,
		CreateTime: m.CreatedAt.Unix(),
	}

}

func Notes(ms []*model.Note) []*note.Note {
	notes := make([]*note.Note, 0)
	for _, m := range ms {
		if n := Note(m); n != nil {
			notes = append(notes, n)
		}
	}

	return notes
}
