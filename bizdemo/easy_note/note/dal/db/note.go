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

package db

import (
	"context"

	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/note/model"
)

// CreateNote  create note info
func CreateNote(ctx context.Context, notes []*model.Note) error {
	if err := DB.WithContext(ctx).Create(notes).Error; err != nil {
		return err
	}
	return nil
}

// MGetNotes  bulk get list of note info
func MGetNotes(ctx context.Context, noteIDs []int64) ([]*model.Note, error) {
	var res []*model.Note
	if len(noteIDs) == 0 {
		return res, nil
	}

	if err := DB.WithContext(ctx).Where("id in ?", noteIDs).Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}

// UpdateNote  update note info
func UpdateNote(ctx context.Context, noteID, userID int64, title, content *string) error {
	params := map[string]interface{}{}
	if title != nil {
		params["title"] = *title
	}
	if content != nil {
		params["content"] = *content
	}
	return DB.WithContext(ctx).Model(&model.Note{}).Where("id = ? and user_id = ?", noteID, userID).
		Updates(params).Error
}

// DelNote  delete note info
func DelNote(ctx context.Context, noteID, userID int64) error {
	return DB.WithContext(ctx).Where("id = ? and user_id = ? ", noteID, userID).Delete(&model.Note{}).Error
}

// QueryNote  query list of note info
func QueryNote(ctx context.Context, userID int64, searchKey *string, limit, offset int) ([]*model.Note, int64, error) {
	var total int64
	var res []*model.Note
	conn := DB.WithContext(ctx).Model(&model.Note{}).Where("user_id = ?", userID)

	if searchKey != nil {
		conn = conn.Where("title like ?", "%"+*searchKey+"%")
	}

	if err := conn.Count(&total).Error; err != nil {
		return res, total, err
	}

	if err := conn.Limit(limit).Offset(offset).Find(&res).Error; err != nil {
		return res, total, err
	}

	return res, total, nil
}
