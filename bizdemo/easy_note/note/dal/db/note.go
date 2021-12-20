/*
 * Copyright 2021 CloudWeGo
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package db

import (
	"context"

	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/note/model"
)

// CreateNote  create notes
func CreateNote(ctx context.Context, notes []*model.Note) error {
	conn := GetDBWriter(ctx)
	if err := conn.Create(notes).Error; err != nil {
		return err
	}
	return nil
}

func MGetNotes(ctx context.Context, noteIDs []int64) ([]*model.Note, error) {
	var res []*model.Note
	if len(noteIDs) == 0 {
		return res, nil
	}

	conn := GetDBReader(ctx)
	if err := conn.Where("id in ?", noteIDs).Find(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}

func UpdateNote(ctx context.Context, noteID, userID int64, title, content *string) error {
	conn := GetDBWriter(ctx)
	params := map[string]interface{}{}
	if title != nil {
		params["title"] = *title
	}

	if content != nil {
		params["content"] = *content
	}
	if err := conn.Model(&model.Note{}).Where("id = ? and user_id = ?", noteID, userID).
		Updates(params).Error; err != nil {
		return err
	}
	return nil
}

func DelNote(ctx context.Context, noteID, userID int64) error {
	conn := GetDBWriter(ctx)
	if err := conn.Where("id = ? and user_id = ? ", noteID, userID).Delete(&model.Note{}).Error; err != nil {
		return err
	}
	return nil
}

func QueryNote(ctx context.Context, userID int64, searchKey *string, limit, offset int) ([]*model.Note, int64, error) {
	conn := GetDBWriter(ctx)
	var total int64
	var res []*model.Note
	conn = conn.Model(&model.Note{}).Where("user_id = ?", userID)

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
